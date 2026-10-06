package ui

import "github.com/a-h/templ"

func mspClientsExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp/clients", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebarClients() +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center;margin-bottom:var(--yr-space-6)">` +
				`<h1 class="msp-page-title" style="margin:0">Clients</h1>` +
				`<yr-button variant="primary">` +
				exampleIcon("plus", "", mspIconSize) +
				` Add Client</yr-button>` +
				`</div>` +

				`<div class="msp-feed">` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="success">healthy</yr-pill>` +
				`<span class="msp-client-name">Acme Corp</span>` +
				`<span class="msp-client-meta">8 users · 3 connectors · 12 workflows</span></div>` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="success">healthy</yr-pill>` +
				`<span class="msp-client-name">Beta LLC</span>` +
				`<span class="msp-client-meta">4 users · 2 connectors · 6 workflows</span></div>` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="warning">attention</yr-pill>` +
				`<span class="msp-client-name">Gamma Inc</span>` +
				`<span class="msp-client-meta">12 users · 5 connectors · 18 workflows</span></div>` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="success">healthy</yr-pill>` +
				`<span class="msp-client-name">Delta Co</span>` +
				`<span class="msp-client-meta">3 users · 1 connector · 4 workflows</span></div>` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="danger">critical</yr-pill>` +
				`<span class="msp-client-name">Epsilon Group</span>` +
				`<span class="msp-client-meta">6 users · 4 connectors · 9 workflows</span></div>` +

				`<div class="msp-client-row">` +
				`<yr-pill tone="success">healthy</yr-pill>` +
				`<span class="msp-client-name">Zeta Partners</span>` +
				`<span class="msp-client-meta">5 users · 2 connectors · 7 workflows</span></div>` +

				`</div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
