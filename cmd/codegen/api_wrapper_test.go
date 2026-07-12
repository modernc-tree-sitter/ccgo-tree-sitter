package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEnsureCoreAPI(t *testing.T) {
	const liveBody = "package grammar\n// live fixture seed\n"
	const existingBody = "package grammar\n// existing dest keep\n"

	tests := []struct {
		name    string
		setup   func(t *testing.T) (outputDir string)
		wantErr string
		check   func(t *testing.T, outputDir string)
	}{
		{
			name: "missing dest copies from live grammar/api.go fixture",
			setup: func(t *testing.T) string {
				root := t.TempDir()
				liveDir := filepath.Join(root, "grammar")
				if err := os.MkdirAll(liveDir, 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(liveDir, "api.go"), []byte(liveBody), 0o644); err != nil {
					t.Fatal(err)
				}
				// outputDir is a sibling tree so dest != live path; walk from cwd finds live.
				out := filepath.Join(root, "out", "grammar")
				if err := os.MkdirAll(out, 0o755); err != nil {
					t.Fatal(err)
				}
				chdir(t, root)
				return out
			},
			check: func(t *testing.T, outputDir string) {
				got, err := os.ReadFile(filepath.Join(outputDir, "api.go"))
				if err != nil {
					t.Fatalf("expected seeded api.go: %v", err)
				}
				if string(got) != liveBody {
					t.Fatalf("seeded content mismatch:\ngot:  %q\nwant: %q", got, liveBody)
				}
			},
		},
		{
			name: "missing dest copies via walk up from outputDir",
			setup: func(t *testing.T) string {
				// root/grammar/api.go is live; dest is root/build/grammar/api.go.
				// Walking from outputDir finds ../../grammar/api.go relative to dest dir.
				root := t.TempDir()
				liveDir := filepath.Join(root, "grammar")
				if err := os.MkdirAll(liveDir, 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(liveDir, "api.go"), []byte(liveBody), 0o644); err != nil {
					t.Fatal(err)
				}
				out := filepath.Join(root, "build", "grammar")
				if err := os.MkdirAll(out, 0o755); err != nil {
					t.Fatal(err)
				}
				// Neutral cwd with no grammar/api.go so only outputDir walk supplies the live file.
				cwd := filepath.Join(root, "cwd")
				if err := os.MkdirAll(cwd, 0o755); err != nil {
					t.Fatal(err)
				}
				chdir(t, cwd)
				return out
			},
			check: func(t *testing.T, outputDir string) {
				got, err := os.ReadFile(filepath.Join(outputDir, "api.go"))
				if err != nil {
					t.Fatalf("expected seeded api.go: %v", err)
				}
				if string(got) != liveBody {
					t.Fatalf("seeded content mismatch:\ngot:  %q\nwant: %q", got, liveBody)
				}
			},
		},
		{
			name: "existing dest is not overwritten",
			setup: func(t *testing.T) string {
				root := t.TempDir()
				// Tempting live source that must not replace dest.
				liveDir := filepath.Join(root, "grammar")
				if err := os.MkdirAll(liveDir, 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(liveDir, "api.go"), []byte(liveBody), 0o644); err != nil {
					t.Fatal(err)
				}
				out := filepath.Join(root, "out", "grammar")
				if err := os.MkdirAll(out, 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(out, "api.go"), []byte(existingBody), 0o644); err != nil {
					t.Fatal(err)
				}
				chdir(t, root)
				return out
			},
			check: func(t *testing.T, outputDir string) {
				got, err := os.ReadFile(filepath.Join(outputDir, "api.go"))
				if err != nil {
					t.Fatal(err)
				}
				if string(got) != existingBody {
					t.Fatalf("dest was overwritten:\ngot:  %q\nwant: %q", got, existingBody)
				}
			},
		},
		{
			name: "missing live returns clear error",
			setup: func(t *testing.T) string {
				root := t.TempDir()
				out := filepath.Join(root, "grammar")
				if err := os.MkdirAll(out, 0o755); err != nil {
					t.Fatal(err)
				}
				// No grammar/api.go anywhere under root; chdir so repo cwd cannot supply live.
				chdir(t, root)
				return out
			},
			wantErr: "hand-maintained grammar/api.go not found",
			check: func(t *testing.T, outputDir string) {
				if _, err := os.Stat(filepath.Join(outputDir, "api.go")); !os.IsNotExist(err) {
					t.Fatalf("api.go should remain missing, stat err=%v", err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputDir := tt.setup(t)
			err := ensureCoreAPI(outputDir)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("ensureCoreAPI: %v", err)
				}
			} else {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %q does not contain %q", err.Error(), tt.wantErr)
				}
			}
			if tt.check != nil {
				tt.check(t, outputDir)
			}
		})
	}
}

func TestFindLiveCoreAPI_skipsDest(t *testing.T) {
	root := t.TempDir()
	// Dest path does not exist; a live file lives one level up as grammar/api.go.
	liveDir := filepath.Join(root, "grammar")
	if err := os.MkdirAll(liveDir, 0o755); err != nil {
		t.Fatal(err)
	}
	body := []byte("package grammar\n// live\n")
	if err := os.WriteFile(filepath.Join(liveDir, "api.go"), body, 0o644); err != nil {
		t.Fatal(err)
	}
	// If dest were the live file path, find must not return it when missing... but missing
	// files fail Stat. Instead place dest at a different missing path under grammar/.
	// Use build/grammar/api.go as dest; live is root/grammar/api.go.
	dest := filepath.Join(root, "build", "grammar", "api.go")
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		t.Fatal(err)
	}
	cwd := filepath.Join(root, "empty-cwd")
	if err := os.MkdirAll(cwd, 0o755); err != nil {
		t.Fatal(err)
	}
	chdir(t, cwd)

	got, err := findLiveCoreAPI(dest)
	if err != nil {
		t.Fatalf("findLiveCoreAPI: %v", err)
	}
	want, err := filepath.Abs(filepath.Join(liveDir, "api.go"))
	if err != nil {
		t.Fatal(err)
	}
	gotAbs, err := filepath.Abs(got)
	if err != nil {
		t.Fatal(err)
	}
	if gotAbs != want {
		t.Fatalf("got %q want %q", gotAbs, want)
	}
}

// chdir switches process cwd for the remainder of the test and restores it on cleanup.
// Required because findLiveCoreAPI walks from os.Getwd().
func chdir(t *testing.T, dir string) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(old); err != nil {
			t.Errorf("restore cwd: %v", err)
		}
	})
}
