package ui

import "testing"

func TestRenderRowEdit(t *testing.T) {
	t.Parallel()

	for _, open := range []bool{false, true} {
		name := "closed"
		if open {
			name = "open"
		}

		html := renderString(t, withChildren(
			RowEdit("Edit name", RowEditProps{Open: open}),
			text("form"),
		))
		writeClassCoverage(t, "row_edit", name, html)
	}
}
