package ui

import "github.com/a-h/templ"

func copyFieldExamples() []Example {
	return []Example{
		{Name: "values", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:1rem;max-width:28rem">` +
				`<yr-copy-field value="https://depot.yardrail.example"></yr-copy-field>` +
				`<yr-copy-field value="tok_live_9f2c41d07a6b4e15b3c8d2a90e7f615c3ab48d70e21f4c96"></yr-copy-field>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
