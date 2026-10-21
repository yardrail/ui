package ui

import "github.com/a-h/templ"

// CreatePanelProps configures a CreatePanel.
type CreatePanelProps struct {
	Attrs templ.Attributes
}

// CreatePanel renders a collapsible creation form panel. Children are the panel content.
func CreatePanel(p *CreatePanelProps) templ.Component {
	return createPanel(p)
}
