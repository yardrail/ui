package ui

import "testing"

func TestRenderTabs(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Tabs(&TabsProps{Label: "Details"}),
		text("<button slot=\"tab\">Tab 1</button><div slot=\"panel\">Panel 1</div>"),
	))
	writeClassCoverage(t, "tabs", "default", html)
}
