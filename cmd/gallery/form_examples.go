package main

import (
	"github.com/a-h/templ"
	ui "github.com/yardrail/ui/pkg"
)

// exampleRenameURL is the rename endpoint of the example team row.
const exampleRenameURL = "/organization/teams/7/rename"

// formExamples covers both Form layouts.
func formExamples() []Example {
	return []Example{
		{Name: "stacked", Component: withChildren(
			ui.Form(ui.FormProps{Action: "/signup", ID: "signup"}),
			ui.Field("Email", "email", ui.FieldProps{Type: ui.FieldEmail, Autocomplete: ui.AutocompleteEmail}),
			ui.Field(
				"Password",
				"password",
				ui.FieldProps{Type: ui.FieldPassword, Autocomplete: ui.AutocompleteNewPassword},
			),
			ui.Button("Create account", ui.ButtonProps{Type: ui.ButtonSubmit}),
		), DisplayURL: ""},
		{Name: "inline", Component: exampleRenameForm(""), DisplayURL: ""},
	}
}

// exampleRenameForm is the inline rename form of the example team row, showing errMsg on its
// Field when non-empty.
func exampleRenameForm(errMsg string) templ.Component {
	return withChildren(
		ui.Form(ui.FormProps{
			ID:     "rename-team-7",
			Action: exampleRenameURL,
			Layout: ui.FormInline,
			Attrs: templ.Attributes{
				"hx-post":   exampleRenameURL,
				"hx-target": "#team-7",
				"hx-swap":   "outerHTML",
			},
		}),
		ui.Field(
			"Team name",
			"name",
			ui.FieldProps{Value: "Platform", HideLabel: true, Error: errMsg, Focus: errMsg != ""},
		),
		ui.Button("Save", ui.ButtonProps{Type: ui.ButtonSubmit}),
	)
}
