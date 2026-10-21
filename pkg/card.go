package ui

import "github.com/a-h/templ"

// CardProps configures a Card.
type CardProps struct {
	Interactive bool
	Attrs       templ.Attributes
}

// Card renders a generic container card. Children are the card content.
func Card(p *CardProps) templ.Component {
	return card(p)
}
