package ui

import "testing"

func TestRenderListRow(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		ListRow(&ListRowProps{}),
		text("row content"),
	))
	writeClassCoverage(t, "list_row", "default", html)
}
