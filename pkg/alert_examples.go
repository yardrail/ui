package ui

import "github.com/a-h/templ"

// alertIconSize is the icon size that sits level with an alert's text.
const alertIconSize = 20

func alertExamples() []Example {
	return []Example{
		{Name: "all-tones", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:32rem">` +
				`<yr-alert tone="success">Your changes were saved.</yr-alert>` +
				`<yr-alert tone="warning">The connector token expires in 3 days.</yr-alert>` +
				`<yr-alert tone="danger">That username is already taken.</yr-alert>` +
				`<yr-alert tone="info">A new version of this template is available.</yr-alert>` +
				`<yr-alert>No tone falls back to info.</yr-alert>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: "with-icon", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:32rem">` +
				`<yr-alert tone="success">` + exampleIcon("check-circle", "", alertIconSize) +
				`<strong>Added!</strong> The template is in your library.</yr-alert>` +
				`<yr-alert tone="danger">` + exampleIcon("circle-x", "", alertIconSize) +
				`<strong>Run failed.</strong> The webhook returned 500.</yr-alert>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
