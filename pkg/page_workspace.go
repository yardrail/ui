package ui

import "github.com/a-h/templ"

// PageWorkspaceProps configures a PageWorkspace.
type PageWorkspaceProps struct {
	Attrs templ.Attributes
}

// PageWorkspace renders the grid-based workspace page shell.
func PageWorkspace(p *PageWorkspaceProps) templ.Component {
	return pageWorkspace(p)
}
