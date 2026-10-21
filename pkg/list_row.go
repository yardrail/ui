package ui

import "github.com/a-h/templ"

// ListRowProps configures a ListRow.
type ListRowProps struct {
	Attrs templ.Attributes
}

// ListRow renders a settings-list row with icon, main, and actions sections.
func ListRow(p *ListRowProps) templ.Component {
	return listRow(p)
}
