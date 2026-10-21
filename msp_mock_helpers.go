package ui

import "github.com/a-h/templ"

const attrActive = "active"

// mspActive returns the active boolean attribute for a sidebar link when link matches active.
func mspActive(link, active string) templ.Attributes {
	if link != active {
		return templ.Attributes{}
	}

	return templ.Attributes{attrActive: true}
}
