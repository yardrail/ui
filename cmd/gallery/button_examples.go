package main

import (
	"github.com/a-h/templ"
	ui "github.com/yardrail/ui/pkg"
)

// buttonExamples covers every Button variant, the type and state props, and Attrs pass-through.
func buttonExamples() []Example {
	return []Example{
		{Name: "primary", Component: ui.Button("Save", ui.ButtonProps{}), DisplayURL: ""},
		{
			Name:       "secondary",
			Component:  ui.Button("Cancel", ui.ButtonProps{Variant: ui.ButtonSecondary}),
			DisplayURL: "",
		},
		{Name: "link", Component: ui.Button("Show details", ui.ButtonProps{Variant: ui.ButtonLink}), DisplayURL: ""},
		{
			Name:       exampleDanger,
			Component:  ui.Button("Delete organization", ui.ButtonProps{Variant: ui.ButtonDanger}),
			DisplayURL: "",
		},
		{
			Name:       "danger-quiet",
			Component:  ui.Button("Remove", ui.ButtonProps{Variant: ui.ButtonDangerQuiet}),
			DisplayURL: "",
		},
		{Name: "disabled", Component: ui.Button("Publish", ui.ButtonProps{Disabled: true}), DisplayURL: ""},
		{
			Name: "submit",
			Component: ui.Button(
				"Create team",
				ui.ButtonProps{Type: ui.ButtonSubmit, Name: "intent", Value: "create"},
			),
			DisplayURL: "",
		},
		{Name: "confirm", Component: ui.Button("Revoke key", ui.ButtonProps{
			Variant: ui.ButtonDanger,
			Confirm: "Revoke this key? Clients using it stop working immediately.",
		}), DisplayURL: ""},
		{Name: "attrs", Component: ui.Button("Archive", ui.ButtonProps{
			Variant: ui.ButtonDangerQuiet,
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
