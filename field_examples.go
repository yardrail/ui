package ui

// fieldExamples covers the hint, error, optional and hidden-label states, and the id prefix a Form
// gives its Fields.
func fieldExamples() []Example {
	return []Example{
		{Name: "default", Component: Field("Display name", "display_name", FieldProps{})},
		{Name: "hint", Component: Field("Slug", "slug", FieldProps{
			Hint: "Lowercase letters, digits and dashes.",
		})},
		{Name: "error", Component: Field("Email", "email", FieldProps{
			Type:  FieldEmail,
			Value: "pat@",
			Error: "Enter a complete email address.",
		})},
		{Name: "hint-error", Component: Field("Webhook URL", "url", FieldProps{
			Type:  FieldURL,
			Value: "ftp://example.com",
			Hint:  "Must start with https://.",
			Error: "Use an https:// URL.",
			Focus: true,
		})},
		{Name: "optional", Component: Field("Organization", "organization", FieldProps{
			Optional:     true,
			Autocomplete: AutocompleteOrganization,
		})},
		{Name: "hidden-label", Component: Field("Search teams", "q", FieldProps{
			Type:      FieldSearch,
			HideLabel: true,
		})},
		{Name: "in-form", Component: withChildren(
			Form(FormProps{ID: "rename-team-7", Action: exampleRenameURL}),
			Field("Team name", "name", FieldProps{Value: "Platform"}),
		)},
	}
}
