package ui

import "testing"

func TestRenderSectionHeader(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SectionHeader(&SectionHeaderProps{}), text("Section")))
	writeClassCoverage(t, "section_header", "default", html)
}
