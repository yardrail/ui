package ui

import "github.com/a-h/templ"

// EmptyStateProps configures an EmptyState.
type EmptyStateProps struct {
	Attrs templ.Attributes
}

// EmptyState renders an empty state placeholder. Children are the message content.
func EmptyState(p *EmptyStateProps) templ.Component {
	return emptyState(p)
}
