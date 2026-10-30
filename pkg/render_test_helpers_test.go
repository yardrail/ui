package ui

import (
	"os"
	"path/filepath"
	"testing"
)

const classCoverageDir = "testdata/classcoverage"

func writeClassCoverage(t *testing.T, component, name, html string) {
	t.Helper()

	dir := filepath.Join(classCoverageDir, component)

	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}

	path := filepath.Join(dir, name+".html")

	//nolint:gosec // test output files don't need restricted permissions
	err = os.WriteFile(path, []byte(html), 0o644)
	if err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
