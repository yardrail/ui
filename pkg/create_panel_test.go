package ui

import "testing"

func TestRenderCreatePanel(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		CreatePanel(&CreatePanelProps{}),
		text("panel content"),
	))
	writeClassCoverage(t, "create_panel", "default", html)
}
