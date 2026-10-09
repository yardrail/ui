package ui

import "github.com/a-h/templ"

// mspCurrent returns aria-current="page" for the nav link named link when it is the active one,
// and no attributes otherwise.
func mspCurrent(link, active string) templ.Attributes {
	if link != active {
		return templ.Attributes{}
	}

	return templ.Attributes{"aria-current": "page"}
}
