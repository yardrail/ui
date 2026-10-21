package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// withChildren renders parent with the given children composed inside it.
func withChildren(parent templ.Component, children ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return parent.Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for _, child := range children {
				err := child.Render(ctx, w)
				if err != nil {
					return err
				}
			}

			return nil
		})), w)
	})
}

// text renders a plain text string as a templ component.
func text(s string) templ.Component {
	return templ.Raw(s)
}
