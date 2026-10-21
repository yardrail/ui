package ui

import "github.com/a-h/templ"

// SectionHeaderProps configures a SectionHeader.
type SectionHeaderProps struct {
	Attrs templ.Attributes
}

// SectionHeader renders a titled section divider. Children are the title text and optional action link.
func SectionHeader(p SectionHeaderProps) templ.Component {
	return sectionHeader(&p)
}
