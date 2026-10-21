package ui

import "github.com/a-h/templ"

// ActivityRowProps configures an ActivityRow.
type ActivityRowProps struct {
	Href    string
	IsError bool
	Attrs   templ.Attributes
}

// ActivityRow renders a run/activity entry. When Href is set it renders as a link;
// otherwise as a div. IsError adds the is-error modifier.
func ActivityRow(p *ActivityRowProps) templ.Component {
	return activityRow(p)
}
