package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestPreprocessorCmdDefaultClangLinux(t *testing.T) {
	t.Setenv("CC", "clang")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("linux", "amd64", "-E", "-o", "-", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"clang", "-E", "-o", "-", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestResolveCCPrefersMingwGCCOverClang(t *testing.T) {
	dir := t.TempDir()
	// LookPath on Windows only matches *.exe; isolate PATH so a host MinGW
	// install cannot shadow the stub (as on windows-latest runners).
	name := "x86_64-w64-mingw32-gcc"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	gcc := filepath.Join(dir, name)
	if err := os.WriteFile(gcc, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
	t.Setenv("CC", "clang")

	want, err := exec.LookPath("x86_64-w64-mingw32-gcc")
	if err != nil {
		t.Fatal(err)
	}
	got := resolveCC("windows", "amd64")
	if got != want {
		t.Fatalf("resolveCC=%q, want %q", got, want)
	}
}

func TestResolveCCKeepsExplicitMingwGCC(t *testing.T) {
	dir := t.TempDir()
	name := "x86_64-w64-mingw32-gcc"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	shadow := filepath.Join(dir, name)
	if err := os.WriteFile(shadow, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)

	explicit := filepath.Join(dir, "explicit-gcc")
	if runtime.GOOS == "windows" {
		explicit += ".exe"
	}
	if err := os.WriteFile(explicit, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CC", explicit)

	got := resolveCC("windows", "amd64")
	if got != explicit {
		t.Fatalf("resolveCC=%q, want explicit CC %q", got, explicit)
	}
}

func TestPreprocessorCmdWindowsMingwGccFlags(t *testing.T) {
	t.Setenv("CC", "/usr/bin/x86_64-w64-mingw32-gcc")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range cmd.Args {
		if strings.HasPrefix(a, "--target=") || strings.Contains(a, "__attribute__") || strings.Contains(a, "__declspec") {
			t.Fatalf("unexpected flag %q in %v", a, cmd.Args)
		}
	}
	if !slices.Contains(cmd.Args, "-D__extension__=") {
		t.Fatalf("missing __extension__ erase in %v", cmd.Args)
	}
}

func TestSplitCompilerEnvWindowsPathKeepsBackslashes(t *testing.T) {
	const cc = `C:\mingw64\bin\x86_64-w64-mingw32-gcc.exe`
	got, err := splitCompilerEnv(cc)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{cc}
	if !slices.Equal(got, want) {
		t.Fatalf("splitCompilerEnv(%q)=%q, want %q", cc, got, want)
	}
}

func TestSplitCompilerEnvWindowsPathWithArgs(t *testing.T) {
	const cc = `C:\mingw64\bin\gcc.exe -DFOO=1`
	got, err := splitCompilerEnv(cc)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{`C:\mingw64\bin\gcc.exe`, "-DFOO=1"}
	if !slices.Equal(got, want) {
		t.Fatalf("splitCompilerEnv(%q)=%q, want %q", cc, got, want)
	}
}

func TestPreprocessorCmdWindowsBackslashCC(t *testing.T) {
	const cc = `C:\mingw64\bin\x86_64-w64-mingw32-gcc.exe`
	t.Setenv("CC", cc)
	t.Setenv("PATH", "/nonexistent")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if cmd.Path != cc && cmd.Args[0] != cc {
		t.Fatalf("Path=%q Args[0]=%q, want %q", cmd.Path, cmd.Args[0], cc)
	}
	if strings.Contains(cmd.Args[0], `C:mingw64`) {
		t.Fatalf("backslashes were eaten: %q", cmd.Args[0])
	}
}

func TestPreprocessorCmdWindowsClangKeepsTarget(t *testing.T) {
	// Force clang by pointing CC at a clang path with no mingw gcc on PATH.
	t.Setenv("CC", "clang")
	t.Setenv("PATH", "/nonexistent")
	t.Setenv("MINGW_SYSROOT", `C:\mingw64`)
	os.Unsetenv("CFLAGS")

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=x86_64-w64-mingw32") {
		t.Fatalf("args=%v", cmd.Args)
	}
}

func TestPreprocessorCmdDarwinFlags(t *testing.T) {
	t.Setenv("CC", "clang")
	os.Unsetenv("CFLAGS")

	for _, goarch := range []string{"amd64", "arm64"} {
		cmd, err := preprocessorCmd("darwin", goarch, "-E", "x.c")
		if err != nil {
			t.Fatal(err)
		}
		for _, flag := range []string{"-fno-blocks", "-D__attribute__(...)=", "-DAPI_AVAILABLE(...)=", "-D_Nonnull="} {
			if !slices.Contains(cmd.Args, flag) {
				t.Fatalf("goarch=%s args=%v, missing %q", goarch, cmd.Args, flag)
			}
		}
	}
}

func TestMingwTriple(t *testing.T) {
	cases := map[string]string{
		"amd64": "x86_64-w64-mingw32",
		"arm64": "aarch64-w64-mingw32",
		"386":   "i686-w64-mingw32",
		"":      "x86_64-w64-mingw32",
	}
	for goarch, want := range cases {
		if got := mingwTriple(goarch); got != want {
			t.Fatalf("mingwTriple(%q)=%q, want %q", goarch, got, want)
		}
	}
}

func TestPreprocessorCmdWindowsArm64ClangTarget(t *testing.T) {
	t.Setenv("CC", "clang")
	t.Setenv("PATH", "/nonexistent")
	t.Setenv("MINGW_SYSROOT", `C:\llvm-mingw`)
	os.Unsetenv("CFLAGS")

	cmd, err := preprocessorCmd("windows", "arm64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=aarch64-w64-mingw32") {
		t.Fatalf("args=%v", cmd.Args)
	}
}

func TestCcgoExtraArgs(t *testing.T) {
	if !slices.Contains(ccgoExtraArgs("darwin"), "-ignore-unsupported-alignment") {
		t.Fatal(ccgoExtraArgs("darwin"))
	}
	if !slices.Contains(ccgoExtraArgs("windows"), "--winapi-no-errno") {
		t.Fatal(ccgoExtraArgs("windows"))
	}
	if slices.Contains(ccgoExtraArgs("windows"), "-winapi-no-errno") {
		t.Fatal("single-dash -winapi-no-errno is rejected by ccgo opt parser")
	}
}
