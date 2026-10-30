package ui

import "testing"

func TestRenderCopyField(t *testing.T) {
	t.Parallel()

	html := renderString(t, CopyField("sk-abc123", &CopyFieldProps{}))
	writeClassCoverage(t, "copy_field", "default", html)
}
