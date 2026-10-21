package ui

import "github.com/a-h/templ"

// ContentProps configures a Content region.
type ContentProps struct {
	Attrs templ.Attributes
}

// Content renders the main content area alongside a sidebar.
func Content(p *ContentProps) templ.Component {
	return content(p)
}

// ContentMainProps configures a ContentMain region.
type ContentMainProps struct {
	Attrs templ.Attributes
}

// ContentMain renders the scrollable main content area within Content.
func ContentMain(p *ContentMainProps) templ.Component {
	return contentMain(p)
}
