package ui

import "github.com/a-h/templ"

// DrawerProps configures a Drawer.
type DrawerProps struct {
	Open     bool
	Position string
	Attrs    templ.Attributes
}

// Drawer renders a slide-out drawer panel.
func Drawer(p *DrawerProps) templ.Component {
	return drawer(p)
}
