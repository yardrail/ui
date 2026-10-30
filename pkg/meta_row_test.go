package ui

import "testing"

func TestRenderMetaRow(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		MetaRow(&MetaRowProps{}),
		text("<span>v1.2</span><span>2h ago</span>"),
	))
	writeClassCoverage(t, "meta_row", "default", html)
}
