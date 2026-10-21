package ui

import "github.com/a-h/templ"

// IconBadgeProps configures an IconBadge.
type IconBadgeProps struct {
	Attrs templ.Attributes
}

// IconBadge renders a circular icon container with the given tone. Children are the icon.
func IconBadge(tone string, p IconBadgeProps) templ.Component {
	return iconBadge(tone, &p)
}
