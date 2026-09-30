package ui

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// update rewrites the golden files: go test . -update.
var update = flag.Bool("update", false, "rewrite testdata golden files from the current render")

// TestSnapshots renders every example and compares it byte for byte with its golden file.
func TestSnapshots(t *testing.T) {
	t.Parallel()

	for _, group := range ExampleGroups() {
		t.Run(group.Component, func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join("testdata", group.Component)
			names := make(map[string]bool, len(group.Examples))

			for _, ex := range group.Examples {
				names[ex.Name+".html"] = true

				checkGolden(t, filepath.Join(dir, ex.Name+".html"), renderExample(t, ex))
			}

			checkNoStaleGoldens(t, dir, names)
		})
	}
}

// renderExample renders ex, with a trailing newline so golden files end like text files.
func renderExample(t *testing.T, ex Example) []byte {
	t.Helper()

	var buf bytes.Buffer

	err := ex.Component.Render(t.Context(), &buf)
	if err != nil {
		t.Fatalf("render %s: %v", ex.Name, err)
	}

	return append(buf.Bytes(), '\n')
}

// checkGolden compares got with the golden file at path, or rewrites it under -update.
func checkGolden(t *testing.T, path string, got []byte) {
	t.Helper()

	if *update {
		err := os.MkdirAll(filepath.Dir(path), 0o750)
		if err != nil {
			t.Fatal(err)
		}

		err = os.WriteFile(path, got, 0o600)
		if err != nil {
			t.Fatal(err)
		}

		return
	}

	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden file (run go test . -update to create it): %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("%s differs from the render (run go test . -update to accept)\ngot:\n%s\nwant:\n%s",
			path, got, want)
	}
}

// checkNoStaleGoldens fails on golden files in dir that no example produces.
func checkNoStaleGoldens(t *testing.T, dir string, names map[string]bool) {
	t.Helper()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".html") && !names[e.Name()] {
			t.Errorf("%s has no example; delete it", filepath.Join(dir, e.Name()))
		}
	}
}
