package ui

import "github.com/a-h/templ"

// buttonExamples covers every Button variant, the type and state props, and Attrs pass-through.
func buttonExamples() []Example {
	return []Example{
		{Name: "primary", Component: Button("Save", ButtonProps{})},
		{Name: "secondary", Component: Button("Cancel", ButtonProps{Variant: Secondary})},
		{Name: "link", Component: Button("Show details", ButtonProps{Variant: Link})},
		{Name: "danger", Component: Button("Delete organization", ButtonProps{Variant: Danger})},
		{Name: "danger-quiet", Component: Button("Remove", ButtonProps{Variant: DangerQuiet})},
		{Name: "disabled", Component: Button("Publish", ButtonProps{Disabled: true})},
		{Name: "submit", Component: Button("Create team", ButtonProps{Type: Submit, Name: "intent", Value: "create"})},
		{Name: "confirm", Component: Button("Revoke key", ButtonProps{
			Variant: Danger,
			Confirm: "Revoke this key? Clients using it stop working immediately.",
		})},
		{Name: "attrs", Component: Button("Archive", ButtonProps{
			Variant: DangerQuiet,
			Attrs: templ.Attributes{
				"id":        "archive-team-7",
				"hx-post":   "/organization/teams/7/archive",
				"hx-target": "closest tr",
				"hx-swap":   "outerHTML",
				"data-team": "7",
			},
		})},
	}
}
