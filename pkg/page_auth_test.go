package ui

import "testing"

func TestRenderPageAuth(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(PageAuth(&PageAuthProps{}), text("content")))
	writeClassCoverage(t, "page_auth", "default", html)
}
