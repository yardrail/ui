package ui

import "testing"

func TestRenderDrawer(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Drawer(&DrawerProps{Position: "right", Open: true}),
		text("drawer content"),
	))
	writeClassCoverage(t, "drawer", "default", html)
}
