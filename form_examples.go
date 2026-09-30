package ui

import "github.com/a-h/templ"

// exampleRenameURL is the rename endpoint of the example team row.
const exampleRenameURL = "/organization/teams/7/rename"

// formExamples covers both Form layouts.
func formExamples() []Example {
	return []Example{
		{Name: "stacked", Component: withChildren(
			Form(FormProps{Action: "/signup", ID: "signup"}),
			Field("Email", "email", FieldProps{Type: FieldEmail, Autocomplete: AutocompleteEmail}),
			Field("Password", "password", FieldProps{Type: FieldPassword, Autocomplete: AutocompleteNewPassword}),
			Button("Create account", ButtonProps{Type: Submit}),
		)},
		{Name: "inline", Component: exampleRenameForm("")},
	}
}

// exampleRenameForm is the inline rename form of the example team row, showing errMsg on its
// Field when non-empty.
func exampleRenameForm(errMsg string) templ.Component {
	return withChildren(
		Form(FormProps{
			ID:     "rename-team-7",
			Action: exampleRenameURL,
			Layout: FormInline,
			Attrs: templ.Attributes{
				"hx-post":   exampleRenameURL,
				"hx-target": "#team-7",
				"hx-swap":   "outerHTML",
			},
		}),
		Field("Team name", "name", FieldProps{Value: "Platform", HideLabel: true, Error: errMsg, Focus: errMsg != ""}),
		Button("Save", ButtonProps{Type: Submit}),
	)
}
