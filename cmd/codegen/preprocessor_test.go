package main

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestPreprocessorCmdDefaultClang(t *testing.T) {
	os.Unsetenv("CC")
	os.Unsetenv("CFLAGS")

	cmd, err := preprocessorCmd("-E", "-o", "-", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"clang", "-E", "-o", "-", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestPreprocessorCmdRespectsCCAndCFLAGS(t *testing.T) {
	t.Setenv("CC", "zig cc")
	t.Setenv("CFLAGS", `--target=aarch64-unknown-linux-gnu -Dfoo='long double'`)

	cmd, err := preprocessorCmd("-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"zig", "cc", "--target=aarch64-unknown-linux-gnu", "-Dfoo=long double", "-E", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestPreprocessorCmdExpandsEnvInCFLAGS(t *testing.T) {
	t.Setenv("CC", "clang")
	t.Setenv("MY_TARGET", "x86_64-unknown-linux-gnu")
	t.Setenv("CFLAGS", "--target=$MY_TARGET")

	cmd, err := preprocessorCmd("-E", "x.c")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"clang", "--target=x86_64-unknown-linux-gnu", "-E", "x.c"}
	if !slices.Equal(cmd.Args, want) {
		t.Fatalf("args=%v, want %v", cmd.Args, want)
	}
}

func TestPreprocessorCmdEmptyCC(t *testing.T) {
	t.Setenv("CC", "   ")
	_, err := preprocessorCmd("-E")
	if err == nil {
		t.Fatal("expected error for empty CC")
	}
}

func TestPreprocessorCmdInvalidSyntax(t *testing.T) {
	t.Setenv("CC", "clang 'unclosed")
	_, err := preprocessorCmd("-E")
	if err == nil || !strings.Contains(err.Error(), "parse CC:") {
		t.Fatalf("err=%v, want parse CC error", err)
	}
	if !strings.Contains(err.Error(), "quote") {
		t.Fatalf("err=%v, want quote syntax error", err)
	}
}

func TestPreprocessorIdentity(t *testing.T) {
	t.Setenv("CC", "clang")
	t.Setenv("CFLAGS", "--target=x86_64-unknown-linux-gnu")
	if got := preprocessorIdentity(); got != "clang --target=x86_64-unknown-linux-gnu" {
		t.Fatalf("identity=%q", got)
	}

	os.Unsetenv("CFLAGS")
	if got := preprocessorIdentity(); got != "clang" {
		t.Fatalf("identity=%q", got)
	}
}

func TestSanitizePreprocessedC(t *testing.T) {
	in := "int x;\n" +
		"typedef _Complex float __cfloat128 __attribute__ ((__mode__ (__TC__)));\n" +
		"typedef __float128 _Float128;\n" +
		"typedef float _Float32;\n" +
		"typedef double _Float64;\n" +
		"typedef double _Float32x;\n" +
		"typedef long double _Float64x;\n" +
		"typedef int register_t __attribute__ ((__mode__ (__word__)));\n" +
		"void f(void);\n"
	got := sanitizePreprocessedC(in)
	for _, bad := range []string{"_Float128", "_Float32", "_Float64", "_Float32x", "_Float64x", "__cfloat128", "__float128"} {
		if strings.Contains(got, bad) {
			t.Fatalf("still contains %q:\n%s", bad, got)
		}
	}
	if !strings.Contains(got, "register_t") || !strings.Contains(got, "void f(void);") {
		t.Fatalf("removed too much:\n%s", got)
	}
}
