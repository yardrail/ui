package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

const violationsFixture = "testdata/violations.templ"

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		paths    []string
		wantOut  []string
		wantCode int
	}{
		{
			name:  "class on top-level, if, for and templ.KV",
			paths: []string{violationsFixture},
			wantOut: []string{
				violationsFixture + ":4: class= is not allowed in app templ",
				violationsFixture + ":7: class= is not allowed in app templ",
				violationsFixture + ":11: class= is not allowed in app templ",
				violationsFixture + ":13: class= is not allowed in app templ",
			},
			wantCode: exitViolations,
		},
		{
			name:     "no class",
			paths:    []string{"testdata/clean.templ"},
			wantOut:  nil,
			wantCode: exitClean,
		},
		{
			name:     "missing file",
			paths:    []string{filepath.Join(t.TempDir(), "missing.templ")},
			wantOut:  nil,
			wantCode: exitError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var stdout, stderr bytes.Buffer

			code := run(tt.paths, &stdout, &stderr)
			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d (stderr: %s)", code, tt.wantCode, stderr.String())
			}

			want := ""
			if len(tt.wantOut) > 0 {
				want = strings.Join(tt.wantOut, "\n") + "\n"
			}

			if stdout.String() != want {
				t.Errorf("output =\n%s\nwant:\n%s", stdout.String(), want)
			}
		})
	}
}
