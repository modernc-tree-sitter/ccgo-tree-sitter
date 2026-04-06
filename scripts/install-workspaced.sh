#!/usr/bin/env bash
set -ex

# Download workspaced
DIR=$(mktemp -d)
git clone https://github.com/lucasew/workspaced.git "$DIR"

# Patch shellcheck to ignore third-party and grammar directories
cat << 'PATCH_EOF' > patch_shellcheck.go
package shellcheck

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasew/workspaced/pkg/provider"
	"github.com/lucasew/workspaced/pkg/provider/lint"
	"github.com/lucasew/workspaced/pkg/tool"

	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Provider struct{}

func New() lint.Linter {
	return &Provider{}
}

func init() {
	lint.Register(New())
}

func (p *Provider) Name() string {
	return "shellcheck"
}

func (p *Provider) Detect(ctx context.Context, dir string) error {
	found := false
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "third-party" || d.Name() == "grammar" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(d.Name(), ".sh") {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil {
		return provider.ErrNotApplicable
	}
	if !found {
		return provider.ErrNotApplicable
	}
	return nil
}

type Issue struct {
	File      string      `json:"file"`
	Line      int         `json:"line"`
	EndLine   int         `json:"endLine"`
	Column    int         `json:"column"`
	EndColumn int         `json:"endColumn"`
	Level     string      `json:"level"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Fix       interface{} `json:"fix"`
}

func (p *Provider) Run(ctx context.Context, dir string) (*sarif.Run, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "third-party" || d.Name() == "grammar" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(d.Name(), ".sh") {
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	args := append([]string{"-f", "json"}, files...)

	cmd, err := tool.EnsureAndRunLazyAt(ctx, dir, "shellcheck", "shellcheck", args...)
	if err != nil {
		slog.Error("failed to setup shellcheck", "err", err)
		return nil, err
	}

	cmd.Dir = dir

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()

	output := stdout.Bytes()
	if len(output) == 0 {
		if stderr.Len() > 0 {
			slog.Error("shellcheck execution failed", "stderr", stderr.String())
			return nil, fmt.Errorf("shellcheck failed: %s", stderr.String())
		}
		return nil, nil
	}

	var issues []Issue
	if err := json.Unmarshal(output, &issues); err != nil {
		slog.Error("failed to parse shellcheck output", "err", err, "output", string(output))
		return nil, err
	}

	return p.convertToSarif(issues), nil
}

func (p *Provider) convertToSarif(issues []Issue) *sarif.Run {
	toolURI := "https://github.com/koalaman/shellcheck"

	run := &sarif.Run{
		Tool: sarif.Tool{
			Driver: &sarif.ToolComponent{
				Name:           "shellcheck",
				InformationURI: &toolURI,
			},
		},
		Results: []*sarif.Result{},
	}

	for _, issue := range issues {
		level := "warning"
		switch issue.Level {
		case "error":
			level = "error"
		case "info", "style":
			level = "note"
		}

		ruleId := fmt.Sprintf("SC%d", issue.Code)
		msg := issue.Message
		fileURI := issue.File

		line := issue.Line
		endLine := issue.EndLine
		col := issue.Column
		endCol := issue.EndColumn

		result := &sarif.Result{
			RuleID: &ruleId,
			Level:  &level,
			Message: sarif.Message{
				Text: &msg,
			},
			Locations: []*sarif.Location{
				{
					PhysicalLocation: &sarif.PhysicalLocation{
						ArtifactLocation: &sarif.ArtifactLocation{
							URI: &fileURI,
						},
						Region: &sarif.Region{
							StartLine:   &line,
							EndLine:     &endLine,
							StartColumn: &col,
							EndColumn:   &endCol,
						},
					},
				},
			},
		}

		run.Results = append(run.Results, result)
	}
	return run
}
PATCH_EOF

cp patch_shellcheck.go "$DIR/pkg/provider/lint/shellcheck/shellcheck.go"
rm patch_shellcheck.go

cd "$DIR"

# Disable govulncheck and golangci as mentioned in prompt memory
sed -i '/_ "github.com\/lucasew\/workspaced\/pkg\/provider\/lint\/govulncheck"/d' pkg/module/prelude/prelude.go || true
sed -i '/_ "workspaced\/pkg\/provider\/lint\/govulncheck"/d' pkg/module/prelude/prelude.go || true

sed -i '/_ "github.com\/lucasew\/workspaced\/pkg\/provider\/lint\/golangci"/d' pkg/module/prelude/prelude.go || true
sed -i '/_ "workspaced\/pkg\/provider\/lint\/golangci"/d' pkg/module/prelude/prelude.go || true

go build -o /tmp/workspaced ./cmd/workspaced
mkdir -p ~/go/bin
cp /tmp/workspaced ~/go/bin/workspaced
