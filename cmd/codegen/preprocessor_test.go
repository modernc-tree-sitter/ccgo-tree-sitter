package main

import (
	"os"
	"path/filepath"
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
	gcc := filepath.Join(dir, "x86_64-w64-mingw32-gcc")
	if err := os.WriteFile(gcc, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("CC", "clang")

	got := resolveCC("windows", "amd64")
	if got != gcc {
		t.Fatalf("resolveCC=%q, want %q", got, gcc)
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

func TestCcgoExtraArgs(t *testing.T) {
	if !slices.Contains(ccgoExtraArgs("darwin"), "-ignore-unsupported-alignment") {
		t.Fatal(ccgoExtraArgs("darwin"))
	}
	if !slices.Contains(ccgoExtraArgs("windows"), "-winapi-no-errno") {
		t.Fatal(ccgoExtraArgs("windows"))
	}
}
