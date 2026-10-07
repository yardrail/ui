package ui

import "github.com/a-h/templ"

func liveDurationExamples() []Example {
	return []Example{
		{Name: "states", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem">` +
				`<div style="display:flex;gap:1rem;align-items:baseline">` +
				`<yr-live-duration started-at="2026-10-01T09:00:00Z"></yr-live-duration>` +
				`<code style="font-size:0.75rem;color:#5C5A59">started-at="2026-10-01T09:00:00Z"</code>` +
				`</div>` +
				`<div style="display:flex;gap:1rem;align-items:baseline">` +
				`<yr-live-duration></yr-live-duration>` +
				`<code style="font-size:0.75rem;color:#5C5A59">no started-at</code>` +
				`</div>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
