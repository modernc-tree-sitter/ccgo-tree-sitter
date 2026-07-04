package main

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestPreprocessorCmdDefaultClangLinux(t *testing.T) {
	os.Unsetenv("CC")
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

func TestPreprocessorCmdWindowsUsesMingwTarget(t *testing.T) {
	os.Unsetenv("CC")
	os.Unsetenv("CFLAGS")
	t.Setenv("MINGW_SYSROOT", `C:\mingw64`)

	cmd, err := preprocessorCmd("windows", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if cmd.Args[0] != "clang" || cmd.Args[1] != "--target=x86_64-w64-mingw32" {
		t.Fatalf("args=%v", cmd.Args)
	}
	if !slices.Contains(cmd.Args, `--sysroot=C:\mingw64`) {
		t.Fatalf("missing sysroot in %v", cmd.Args)
	}
	for _, flag := range ccgoFriendlyMacros() {
		if !slices.Contains(cmd.Args, flag) {
			t.Fatalf("missing %q in %v", flag, cmd.Args)
		}
	}

	os.Unsetenv("MINGW_SYSROOT")
	t.Setenv("MINGW_SYSROOT", `C:\mingw64`) // x86_64 root must not apply to arm64
	cmd, err = preprocessorCmd("windows", "arm64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=aarch64-w64-mingw32") {
		t.Fatalf("args=%v", cmd.Args)
	}
	for _, a := range cmd.Args {
		if strings.HasPrefix(a, "--sysroot=") {
			t.Fatalf("arm64 should reject x86_64 sysroot, got %v", cmd.Args)
		}
	}
}

func TestPreprocessorCmdDarwinFlags(t *testing.T) {
	os.Unsetenv("CC")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("darwin", "arm64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	for _, flag := range []string{"-fno-blocks", "-D__attribute__(...)=", "-DAPI_AVAILABLE(...)=", "-D_Nonnull="} {
		if !slices.Contains(cmd.Args, flag) {
			t.Fatalf("args=%v, missing %q", cmd.Args, flag)
		}
	}
}

func TestPreprocessorCmdRespectsCCAndCFLAGS(t *testing.T) {
	t.Setenv("CC", "zig cc")
	t.Setenv("CFLAGS", `--target=aarch64-unknown-linux-gnu -Dfoo='long double'`)
	os.Unsetenv("MINGW_SYSROOT")

	cmd, err := preprocessorCmd("linux", "amd64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"zig", "cc", "--target=aarch64-unknown-linux-gnu", "-Dfoo=long double", "-E", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestPreprocessorCmdEmptyCC(t *testing.T) {
	t.Setenv("CC", "   ")
	_, err := preprocessorCmd("linux", "amd64", "-E")
	if err == nil {
		t.Fatal("expected error for empty CC")
	}
}

func TestMingwTriple(t *testing.T) {
	if mingwTriple("amd64") != "x86_64-w64-mingw32" || mingwTriple("arm64") != "aarch64-w64-mingw32" {
		t.Fatalf("%q %q", mingwTriple("amd64"), mingwTriple("arm64"))
	}
}

func TestSanitizePreprocessedC(t *testing.T) {
	in := "typedef __float128 _Float128;\ntypedef int register_t;\n"
	got := sanitizePreprocessedC(in)
	if strings.Contains(got, "_Float128") || !strings.Contains(got, "register_t") {
		t.Fatalf("%q", got)
	}
}
