package ui

import "github.com/a-h/templ"

// PageAuthProps configures a PageAuth.
type PageAuthProps struct {
	Attrs templ.Attributes
}

// PageAuth renders a centered auth page layout. Children are the page content.
func PageAuth(p *PageAuthProps) templ.Component {
	return pageAuth(p)
}
