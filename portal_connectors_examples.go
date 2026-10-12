package ui

import "github.com/a-h/templ"

func portalConnectorsExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/automations/connectors", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebarAutomations("connectors") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +

				// Title row
				`<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:var(--yr-space-6)">` +
				`<h1 class="portal-page-title" style="margin-bottom:0">Connectors</h1>` +
				`<button class="yr-button">+ Add Connector</button>` +
				`</div>` +

				// Connected section
				`<yr-section-header>Connected</yr-section-header>` +
				`<div style="display:flex;flex-direction:column;gap:var(--yr-space-3);margin-bottom:var(--yr-space-8)">` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon"><img src="` + exampleLogo + `" alt="Slack" width="40" height="40"></div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Slack</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>workspace-prod</span>` +
				` <yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Manage</button>` +
				`</div></div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">HS</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">HubSpot</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>marketing-hub</span>` +
				` <yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Manage</button>` +
				`</div></div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">GS</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Google Sheets</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>ops-shared@acme.com</span>` +
				` <yr-pill tone="success" shape="pill">Connected</yr-pill>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Manage</button>` +
				`</div></div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">ST</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Stripe</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>prod-account</span>` +
				` <yr-pill tone="warning" shape="pill">Expiring</yr-pill>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Manage</button>` +
				`</div></div>` +

				`</div>` +

				// Available section
				`<yr-section-header>Available</yr-section-header>` +
				`<div style="display:flex;flex-direction:column;gap:var(--yr-space-3)">` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">JI</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Jira</div>` +
				`<div class="yr-list-row-desc">Project management and issue tracking</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button">Connect</button>` +
				`</div></div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">SF</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Salesforce</div>` +
				`<div class="yr-list-row-desc">CRM and sales pipeline management</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button">Connect</button>` +
				`</div></div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon" style="background:var(--yr-surface-muted);` +
				`font-family:var(--yr-font-ui);font-weight:600;font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">ZD</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Zendesk</div>` +
				`<div class="yr-list-row-desc">Customer support and ticketing</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button">Connect</button>` +
				`</div></div>` +

				`</div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
