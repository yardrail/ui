package ui

import "github.com/a-h/templ"

// NavShellProps configures a NavShell.
type NavShellProps struct {
	Attrs templ.Attributes
}

// NavShell renders the top navigation shell.
func NavShell(p *NavShellProps) templ.Component {
	return navShell(p)
}
