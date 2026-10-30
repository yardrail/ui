package ui

import (
	"fmt"
	"testing"

	"github.com/a-h/templ"
)

func TestRenderButton(t *testing.T) {
	t.Parallel()

	i := 0

	for _, v := range (ButtonVariant{}).All() {
		for _, bt := range (ButtonType{}).All() {
			for _, disabled := range []bool{false, true} {
				html := renderString(t, Button("Click", ButtonProps{
					Variant:  v,
					Type:     bt,
					Disabled: disabled,
				}))
				writeClassCoverage(t, "button", fmt.Sprintf("%03d", i), html)
				i++
			}
		}
	}

	html := renderString(t, Button("Submit", ButtonProps{
		Name:    "action",
		Value:   "create",
		Confirm: "Are you sure?",
		Attrs:   templ.Attributes{"hx-post": "/api"},
	}))
	writeClassCoverage(t, "button", "with-attrs", html)
}
