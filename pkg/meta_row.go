package ui

import "github.com/a-h/templ"

// MetaRowProps configures a MetaRow.
type MetaRowProps struct {
	Attrs templ.Attributes
}

// MetaRow renders a middot-joined row of metadata fragments.
func MetaRow(p *MetaRowProps) templ.Component {
	return metaRow(p)
}
