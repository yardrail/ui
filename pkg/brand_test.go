package ui

import "testing"

func TestRenderBrand(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Brand(&BrandProps{Href: "/"}),
		text("Yardrail"),
	))
	writeClassCoverage(t, "brand", "default", html)
}
