package ui

import "github.com/a-h/templ"

func connectorBindingRowExamples() []Example {
	return []Example{
		{Name: exampleRows, Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:36rem">` +

				`<div class="yr-meta-row">` +
				`<span>email.send</span>` +
				`<span>SendGrid</span>` +
				`<yr-pill tone="success" shape="pill">Bound</yr-pill>` +
				`</div>` +

				`<div class="yr-meta-row">` +
				`<span>crm.read</span>` +
				`<span>HubSpot</span>` +
				`<yr-pill tone="success" shape="pill">Bound</yr-pill>` +
				`</div>` +

				`<div class="yr-meta-row">` +
				`<span>chat.post</span>` +
				`<yr-pill tone="danger" shape="pill">Not connected</yr-pill>` +
				`<a href="#" style="font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-accent)">Set up</a>` +
				`</div>` +

				`</div>`,
		), DisplayURL: ""},
	}
}
