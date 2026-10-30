package ui

import "testing"

func TestRenderPageWorkspace(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(PageWorkspace(&PageWorkspaceProps{}), text("content")))
	writeClassCoverage(t, "page_workspace", "default", html)
}
