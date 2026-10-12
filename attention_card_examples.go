package ui

import "github.com/a-h/templ"

const attentionIconSize = 20

func attentionCardExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: templ.Raw(
			`<div style="max-width:40rem">` +
				`<div class="yr-action-card yr-action-card--urgent">` +
				`<div class="yr-icon-badge yr-icon-badge--amber">` +
				exampleIcon("clock", "", attentionIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Approve refund for order #1042</h3>` +
				`<p class="yr-action-card-desc">Waiting since 2h ago · Refund triage</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
				`</div>` +
				`</div>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: exampleDanger, Component: templ.Raw(
			`<div style="max-width:40rem">` +
				`<div class="yr-action-card yr-action-card--urgent-danger">` +
				`<div class="yr-icon-badge yr-icon-badge--danger">` +
				exampleIcon("alert-triangle", "", attentionIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Sync contacts failed</h3>` +
				`<p class="yr-action-card-desc">Connector returned 401 · 3h ago</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				`</div>` +
				`</div>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
