package ui

import (
	"fmt"
	"testing"

	"github.com/a-h/templ"
)

func TestRenderActivityRow(t *testing.T) {
	t.Parallel()

	for i, isError := range []bool{false, true} {
		html := renderString(t, withChildren(
			ActivityRow(&ActivityRowProps{IsError: isError, Href: "#"}),
			text("activity"),
		))
		writeClassCoverage(t, "activity_row", fmt.Sprintf("link-%03d", i), html)
	}

	for i, isError := range []bool{false, true} {
		html := renderString(t, withChildren(
			ActivityRow(&ActivityRowProps{IsError: isError}),
			text("activity"),
		))
		writeClassCoverage(t, "activity_row", fmt.Sprintf("div-%03d", i), html)
	}

	html := renderString(t, withChildren(
		ActivityRow(&ActivityRowProps{Href: "#", Attrs: templ.Attributes{"data-id": "123"}}),
		text("activity"),
	))
	writeClassCoverage(t, "activity_row", "with-attrs", html)
}
