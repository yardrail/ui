package ui

import "github.com/a-h/templ"

// PageSetupProps configures a PageSetup.
type PageSetupProps struct {
	Attrs templ.Attributes
}

// PageSetup renders a setup/onboarding page layout. Children are the page content.
func PageSetup(p *PageSetupProps) templ.Component {
	return pageSetup(p)
}

// PageSetupContentProps configures a PageSetupContent.
type PageSetupContentProps struct {
	Attrs templ.Attributes
}

// PageSetupContent renders the content area of a setup page.
func PageSetupContent(p *PageSetupContentProps) templ.Component {
	return pageSetupContent(p)
}
