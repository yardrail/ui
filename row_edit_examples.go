package ui

// rowEditExamples covers the closed and open states of a RowEdit holding an inline rename Form.
func rowEditExamples() []Example {
	return []Example{
		{Name: "closed", Component: withChildren(RowEdit("Rename", RowEditProps{}), exampleRenameForm(""))},
		{Name: "open", Component: withChildren(
			RowEdit("Rename", RowEditProps{Open: true}),
			exampleRenameForm("Team name is required."),
		)},
	}
}
