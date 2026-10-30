package ui

import "testing"

func TestRenderBreadcrumb(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Breadcrumb(&BreadcrumbProps{}),
		text("<a>Home</a><span>/</span><span>Current</span>"),
	))
	writeClassCoverage(t, "breadcrumb", "default", html)
}
