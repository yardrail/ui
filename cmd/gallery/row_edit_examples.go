package main

import ui "github.com/yardrail/ui/pkg"

// rowEditExamples covers the closed and open states of a RowEdit holding an inline rename Form.
func rowEditExamples() []Example {
	return []Example{
		{
			Name:       "closed",
			Component:  withChildren(ui.RowEdit("Rename", ui.RowEditProps{}), exampleRenameForm("")),
			DisplayURL: "",
		},
		{Name: "open", Component: withChildren(
			ui.RowEdit("Rename", ui.RowEditProps{Open: true}),
			exampleRenameForm("Team name is required."),
		), DisplayURL: ""},
	}
}
