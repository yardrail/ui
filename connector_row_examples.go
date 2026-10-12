package ui

import "github.com/a-h/templ"

func connectorRowExamples() []Example {
	return []Example{
		{Name: "rows", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:44rem">` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon"><img src="` + exampleLogo + `" alt="Slack" width="40" height="40"></div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Slack</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>workspace-prod</span>` +
				` <yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button yr-button--secondary">Manage</a>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon"><img src="` + exampleLogo + `" alt="HubSpot" width="40" height="40"></div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">HubSpot</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>marketing-hub</span>` +
				` <yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button yr-button--secondary">Manage</a>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon"><img src="` + exampleLogo + `" alt="Jira" width="40" height="40"></div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Jira</div>` +
				`<div class="yr-list-row-meta">` +
				`<yr-pill tone="danger" shape="pill">Disconnected</yr-pill>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button">Connect</a>` +
				`</div>` +
				`</div>` +

				`</div>`,
		)},
	}
}
