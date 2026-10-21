package main

import ui "github.com/yardrail/ui/pkg"

// fieldExamples covers the hint, error, optional and hidden-label states, and the id prefix a Form
// gives its Fields.
func fieldExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: ui.Field("Display name", "display_name", ui.FieldProps{}), DisplayURL: ""},
		{Name: "hint", Component: ui.Field("Slug", "slug", ui.FieldProps{
			Hint: "Lowercase letters, digits and dashes.",
		}), DisplayURL: ""},
		{Name: "error", Component: ui.Field("Email", "email", ui.FieldProps{
			Type:  ui.FieldEmail,
			Value: "pat@",
			Error: "Enter a complete email address.",
		}), DisplayURL: ""},
		{Name: "hint-error", Component: ui.Field("Webhook URL", "url", ui.FieldProps{
			Type:  ui.FieldURL,
			Value: "ftp://example.com",
			Hint:  "Must start with https://.",
			Error: "Use an https:// URL.",
			Focus: true,
		}), DisplayURL: ""},
		{Name: "optional", Component: ui.Field("Organization", "organization", ui.FieldProps{
			Optional:     true,
			Autocomplete: ui.AutocompleteOrganization,
		}), DisplayURL: ""},
		{Name: "hidden-label", Component: ui.Field("Search teams", "q", ui.FieldProps{
			Type:      ui.FieldSearch,
			HideLabel: true,
		}), DisplayURL: ""},
		{Name: "in-form", Component: withChildren(
			ui.Form(ui.FormProps{ID: "rename-team-7", Action: exampleRenameURL}),
			ui.Field("Team name", "name", ui.FieldProps{Value: "Platform"}),
		), DisplayURL: ""},
	}
}
