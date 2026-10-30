package ui

import "testing"

func TestRenderNavShell(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		NavShell(&NavShellProps{Left: text("logo"), Center: text("search")}),
		text("right"),
	))
	writeClassCoverage(t, "nav_shell", "default", html)
}
