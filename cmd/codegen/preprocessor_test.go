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
	want := []string{"clang", "--target=x86_64-w64-mingw32", `--sysroot=C:\mingw64`, "-E", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}

	cmd, err = preprocessorCmd("windows", "arm64", "-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Contains(cmd.Args, "--target=aarch64-w64-mingw32") {
		t.Fatalf("args=%v, want aarch64 mingw triple", cmd.Args)
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
	for _, flag := range []string{"-fno-blocks", "-D_Nonnull=", "-D_Nullable=", "-D_Null_unspecified="} {
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

func TestPreprocessorCmdInvalidSyntax(t *testing.T) {
	t.Setenv("CC", "clang 'unclosed")
	_, err := preprocessorCmd("linux", "amd64", "-E")
	if err == nil || !strings.Contains(err.Error(), "parse CC:") {
		t.Fatalf("err=%v, want parse CC error", err)
	}
}

func TestPreprocessorIdentity(t *testing.T) {
	os.Unsetenv("CC")
	os.Unsetenv("CFLAGS")
	os.Unsetenv("MINGW_SYSROOT")
	if got := preprocessorIdentity("linux", "amd64"); got != "clang" {
		t.Fatalf("identity=%q", got)
	}
	if got := preprocessorIdentity("windows", "amd64"); !strings.Contains(got, "--target=x86_64-w64-mingw32") {
		t.Fatalf("identity=%q", got)
	}
}

func TestMingwTriple(t *testing.T) {
	if mingwTriple("amd64") != "x86_64-w64-mingw32" {
		t.Fatal(mingwTriple("amd64"))
	}
	if mingwTriple("arm64") != "aarch64-w64-mingw32" {
		t.Fatal(mingwTriple("arm64"))
	}
}

func TestSanitizePreprocessedC(t *testing.T) {
	in := "int x;\n" +
		"typedef _Complex float __cfloat128 __attribute__ ((__mode__ (__TC__)));\n" +
		"typedef __float128 _Float128;\n" +
		"typedef float _Float32;\n" +
		"typedef int register_t __attribute__ ((__mode__ (__word__)));\n" +
		"void f(void);\n"
	got := sanitizePreprocessedC(in)
	for _, bad := range []string{"_Float128", "_Float32", "__cfloat128", "__float128"} {
		if strings.Contains(got, bad) {
			t.Fatalf("still contains %q:\n%s", bad, got)
		}
	}
	if !strings.Contains(got, "register_t") || !strings.Contains(got, "void f(void);") {
		t.Fatalf("removed too much:\n%s", got)
	}
}
