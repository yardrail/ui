package ui

import "github.com/a-h/templ"

// StatProps configures a Stat.
type StatProps struct {
	Label string
	Attrs templ.Attributes
}

// Stat renders a number + label pair.
func Stat(number string, p *StatProps) templ.Component {
	return stat(number, p)
}
