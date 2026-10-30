package ui

import "testing"

func TestRenderPageSetup(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(PageSetup(&PageSetupProps{}), text("content")))
	writeClassCoverage(t, "page_setup", "default", html)
}

func TestRenderPageSetupContent(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(PageSetupContent(&PageSetupContentProps{}), text("content")))
	writeClassCoverage(t, "page_setup_content", "default", html)
}
