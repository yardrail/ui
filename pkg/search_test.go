package ui

import "testing"

func TestRenderSearch(t *testing.T) {
	t.Parallel()

	for _, v := range (SearchVariant{}).All() {
		html := renderString(t, Search(&SearchProps{
			Variant:     v,
			Placeholder: "Search…",
		}))
		writeClassCoverage(t, "search", v.v, html)
	}
}
