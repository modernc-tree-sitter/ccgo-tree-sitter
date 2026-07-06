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
	if !slices.Contains(ccgoExtraArgs("windows"), "-winapi-no-errno") {
		t.Fatal(ccgoExtraArgs("windows"))
	}
}
