package ui

import "github.com/a-h/templ"

// buttonExamples covers every Button variant, the type and state props, and Attrs pass-through.
func buttonExamples() []Example {
	return []Example{
		{Name: "primary", Component: Button("Save", ButtonProps{}), DisplayURL: ""},
		{Name: "secondary", Component: Button("Cancel", ButtonProps{Variant: Secondary}), DisplayURL: ""},
		{Name: "link", Component: Button("Show details", ButtonProps{Variant: Link}), DisplayURL: ""},
		{Name: "danger", Component: Button("Delete organization", ButtonProps{Variant: Danger}), DisplayURL: ""},
		{Name: "danger-quiet", Component: Button("Remove", ButtonProps{Variant: DangerQuiet}), DisplayURL: ""},
		{Name: "disabled", Component: Button("Publish", ButtonProps{Disabled: true}), DisplayURL: ""},
		{
			Name:       "submit",
			Component:  Button("Create team", ButtonProps{Type: Submit, Name: "intent", Value: "create"}),
			DisplayURL: "",
		},
		{Name: "confirm", Component: Button("Revoke key", ButtonProps{
			Variant: Danger,
			Confirm: "Revoke this key? Clients using it stop working immediately.",
		}), DisplayURL: ""},
		{Name: "attrs", Component: Button("Archive", ButtonProps{
			Variant: DangerQuiet,
			Attrs: templ.Attributes{
				"id":        "archive-team-7",
				"hx-post":   "/organization/teams/7/archive",
				"hx-target": "closest tr",
				"hx-swap":   "outerHTML",
				"data-team": "7",
			},
		}), DisplayURL: ""},
	}
}
