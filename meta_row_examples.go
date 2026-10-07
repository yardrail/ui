package ui

import "github.com/a-h/templ"

func metaRowExamples() []Example {
	return []Example{
		{Name: "fragments", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:1rem">` +
				`<div class="yr-meta-row">` +
				`<span>12 templates published</span>` +
				`<span>4.2k installs</span>` +
				`<span>Joined 3 Mar 2026</span>` +
				`</div>` +
				`<div class="yr-meta-row">` +
				`<yr-pill tone="blue">Analytics</yr-pill>` +
				`<yr-rating value="4.5" reviews="42"></yr-rating>` +
				`<span>OAuth 2.0</span>` +
				`</div>` +
				`<div class="yr-meta-row"><span>A single fragment has no separator</span></div>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
