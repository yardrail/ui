package ui

import "github.com/a-h/templ"

// Pill shapes.
const (
	ShapePill = "pill"
)

// PillProps configures a Pill.
type PillProps struct {
	Tone  string
	Shape string
	Attrs templ.Attributes
}

// Pill renders a colored label.
func Pill(text string, p *PillProps) templ.Component {
	return pill(text, p)
}
