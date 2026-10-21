package ui

import "github.com/a-h/templ"

// ConnectorPillProps configures a ConnectorPill.
type ConnectorPillProps struct {
	Title string
	Attrs templ.Attributes
}

// ConnectorPill renders a small connector icon badge. Children are the icon or initials.
func ConnectorPill(p *ConnectorPillProps) templ.Component {
	return connectorPill(p)
}
