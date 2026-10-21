package ui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

const (
	attrHxPost       = "hx-post"
	attrHxTarget     = "hx-target"
	attrDataHxTarget = "data-hx-target"
)

func TestAttrs(t *testing.T) {
	t.Parallel()

	t.Run("dropped attributes panic", func(t *testing.T) {
		t.Parallel()

		for _, key := range []string{
			"class", "style", "onclick", "hx-on:click", "hx-confirm", "role", "tabindex",
			"data-hx-on:click", "data-hx-on::after-request", "data-hx-confirm",
		} {
			t.Run(key, func(t *testing.T) {
				t.Parallel()

				assertPanics(t, Button("Save", ButtonProps{Attrs: templ.Attributes{key: "x"}}))
			})
		}
	})

	t.Run("allowed attributes pass through", func(t *testing.T) {
		t.Parallel()

		attrs := templ.Attributes{
			attrHxPost:       "/teams/7/rename",
			attrHxTarget:     "#team-7",
			attrDataHxTarget: "#team-7-row",
			"id":             "save-7",
			"data-x":         "y",
			"aria-label":     "Save team 7",
		}
		html := renderString(t, Button("Save", ButtonProps{Attrs: attrs}))

		for _, want := range []string{
			attrHxPost + `="/teams/7/rename"`,
			attrHxTarget + `="#team-7"`,
			attrDataHxTarget + `="#team-7-row"`,
			`id="save-7"`,
			`data-x="y"`,
			`aria-label="Save team 7"`,
		} {
			if !strings.Contains(html, want) {
				t.Errorf("render is missing %s:\n%s", want, html)
			}
		}
	})

	t.Run("empty label panics", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, Button("", ButtonProps{}))
	})
}

// renderString renders c and returns the markup.
func renderString(t *testing.T, c templ.Component) string {
	t.Helper()

	var buf bytes.Buffer

	err := c.Render(t.Context(), &buf)
	if err != nil {
		t.Fatalf("render: %v", err)
	}

	return buf.String()
}

// assertPanics fails unless rendering c panics.
func assertPanics(t *testing.T, c templ.Component) {
	t.Helper()

	defer func() {
		if recover() == nil {
			t.Error("render did not panic")
		}
	}()

	err := c.Render(t.Context(), &bytes.Buffer{})
	if err != nil {
		t.Errorf("render: %v", err)
	}
}
