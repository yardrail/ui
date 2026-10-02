package ui

import "github.com/a-h/templ"

func connectorPillExamples() []Example {
	return []Example{
		{Name: "icon", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.5rem">` +
				`<div class="yr-connector-pill--icon" title="Slack"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`<div class="yr-connector-pill--icon" title="HubSpot">HS</div>` +
				`<div class="yr-connector-pill--icon" title="Google Sheets">GS</div>` +
				`</div>`,
		)},
		{Name: "large", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center">` +
				`<a href="#" class="yr-connector-pill--large" title="Slack">` +
				`<div class="yr-connector-pill-icon"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`<span class="yr-connector-pill-name">Slack</span>` +
				`</a>` +
				`<a href="#" class="yr-connector-pill--large" title="HubSpot">` +
				`<div class="yr-connector-pill-icon"><span class="yr-connector-pill-initials">HS</span></div>` +
				`<span class="yr-connector-pill-name">HubSpot</span>` +
				`</a>` +
				`<span class="yr-connector-pill--large" title="Google Sheets">` +
				`<div class="yr-connector-pill-icon"><span class="yr-connector-pill-initials">GS</span></div>` +
				`<span class="yr-connector-pill-name">Google Sheets</span>` +
				`</span>` +
				`</div>`,
		)},
	}
}
