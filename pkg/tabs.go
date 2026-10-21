package ui

import "github.com/a-h/templ"

// TabsProps configures a Tabs widget.
type TabsProps struct {
	Label    string
	SyncHash bool
	Attrs    templ.Attributes
}

// Tabs renders an ARIA tabs widget. Children are tab buttons and panels.
func Tabs(p *TabsProps) templ.Component {
	return tabs(p)
}
