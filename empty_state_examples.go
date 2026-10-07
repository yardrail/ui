package ui

import "github.com/a-h/templ"

// emptyStateIconSize is the icon size the apps use in an empty state.
const emptyStateIconSize = 36

func emptyStateExamples() []Example {
	return []Example{
		{Name: "text", Component: templ.Raw(
			`<yr-empty-state>No teams yet. Add one above.</yr-empty-state>`,
		), DisplayURL: ""},
		{Name: "with-icon", Component: templ.Raw(
			`<yr-empty-state>` + exampleIcon("search", `slot="icon"`, emptyStateIconSize) +
				`No templates match your search.</yr-empty-state>`,
		), DisplayURL: ""},
	}
}
