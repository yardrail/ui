package ui

import "github.com/a-h/templ"

func listRowExamples() []Example {
	return []Example{
		{Name: "rows", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:44rem">` +
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon"><img src="` + exampleLogo + `" alt="Slack" width="40" height="40"></div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Slack</div>` +
				`<div class="yr-list-row-desc">ellery-logistics.slack.com</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`<a href="#" class="yr-button yr-button--secondary">Manage</a>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">HS</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">HubSpot</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button">Connect</a>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-main">` +
				`<a href="#" class="yr-list-row-name">Refund triage</a>` +
				`<p class="yr-list-row-desc">Routes refund requests to the right approver.</p>` +
				`<div class="yr-list-row-meta">` +
				`<yr-pill tone="accent">Verified</yr-pill>` +
				`<span>Purchased 12 Sep 2026</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button">Install</a>` +
				`</div>` +
				`</div>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
