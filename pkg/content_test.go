package ui

import "testing"

func TestRenderContent(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Content(&ContentProps{}),
		text("content"),
	))
	writeClassCoverage(t, "content", "default", html)
}

func TestRenderContentMain(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		ContentMain(&ContentMainProps{}),
		text("main content"),
	))
	writeClassCoverage(t, "content_main", "default", html)
}
