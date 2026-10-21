package ui

import "github.com/a-h/templ"

// WorkflowCardProps configures a WorkflowCard.
type WorkflowCardProps struct {
	Status      templ.Component
	Runs        string
	SuccessRate string
	Action      templ.Component
	Attrs       templ.Attributes
}

// WorkflowCard renders an installed workflow summary card.
func WorkflowCard(name, footerText string, p *WorkflowCardProps) templ.Component {
	return workflowCard(name, footerText, p)
}
