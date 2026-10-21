package ui

import "github.com/a-h/templ"

// ActionCardProps configures an ActionCard.
type ActionCardProps struct {
	Urgent       bool
	UrgentDanger bool
	Attrs        templ.Attributes
}

// ActionCard renders an action prompt card. Children include icon-badge, body, and action sections.
func ActionCard(p *ActionCardProps) templ.Component {
	return actionCard(p)
}
