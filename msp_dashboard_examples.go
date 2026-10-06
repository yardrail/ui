package ui

import "github.com/a-h/templ"

func mspDashboardExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebar() +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="msp-page-title">Dashboard</h1>` +

				// Client health summary
				`<div class="msp-section">` +
				`<div class="msp-grid msp-grid-4">` +
				`<div class="msp-stat-card">` +
				`<p class="msp-stat-value">12</p>` +
				`<p class="msp-stat-label">Total Clients</p></div>` +
				`<div class="msp-stat-card" style="border-left:3px solid var(--yr-success)">` +
				`<p class="msp-stat-value">9</p>` +
				`<p class="msp-stat-label">Healthy</p></div>` +
				`<div class="msp-stat-card" style="border-left:3px solid var(--yr-warning)">` +
				`<p class="msp-stat-value">2</p>` +
				`<p class="msp-stat-label">Needs Attention</p></div>` +
				`<div class="msp-stat-card" style="border-left:3px solid var(--yr-danger)">` +
				`<p class="msp-stat-value">1</p>` +
				`<p class="msp-stat-label">Critical</p></div>` +
				`</div></div>` +

				`<div class="msp-grid msp-grid-2">` +

				// Alerts
				`<div class="msp-section">` +
				`<p class="msp-section-title">Alerts</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="danger">critical</yr-pill>` +
				` Workflow failed — Acme Corp` +
				`<span class="msp-feed-time">2m ago</span></div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="warning">warning</yr-pill>` +
				` Connector unhealthy — Beta LLC` +
				`<span class="msp-feed-time">18m ago</span></div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="warning">warning</yr-pill>` +
				` API key expires in 7d — Gamma Inc` +
				`<span class="msp-feed-time">1h ago</span></div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="info">info</yr-pill>` +
				` New user joined — Delta Co` +
				`<span class="msp-feed-time">3h ago</span></div>` +
				`</div></div>` +

				// Credential health
				`<div class="msp-section">` +
				`<p class="msp-section-title">Credential Health</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-stat-card">` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center">` +
				`<div><p class="msp-stat-value" style="font-size:var(--yr-font-size-lg)">` +
				`42</p><p class="msp-stat-label">Healthy</p></div>` +
				`<div><p class="msp-stat-value" style="font-size:var(--yr-font-size-lg);` +
				`color:var(--yr-warning)">3</p>` +
				`<p class="msp-stat-label">Expiring &lt;30d</p></div>` +
				`<div><p class="msp-stat-value" style="font-size:var(--yr-font-size-lg);` +
				`color:var(--yr-danger)">1</p>` +
				`<p class="msp-stat-label">Expired</p></div>` +
				`</div></div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="danger">expired</yr-pill>` +
				` Slack OAuth — Acme Corp</div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="warning">14d</yr-pill>` +
				` HubSpot API — Beta LLC</div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="warning">22d</yr-pill>` +
				` Stripe key — Gamma Inc</div>` +
				`</div></div>` +

				`</div>` + // close grid-2

				`<div class="msp-grid msp-grid-2">` +

				// Recent activity
				`<div class="msp-section">` +
				`<p class="msp-section-title">Recent Activity</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-feed-item">` +
				exampleIcon("user-plus", "", mspIconSize) +
				` User invited to Acme Corp` +
				`<span class="msp-feed-time">10m ago</span></div>` +
				`<div class="msp-feed-item">` +
				exampleIcon("play", "", mspIconSize) +
				` Workflow deployed to Beta LLC` +
				`<span class="msp-feed-time">45m ago</span></div>` +
				`<div class="msp-feed-item">` +
				exampleIcon("refresh-cw", "", mspIconSize) +
				` Credential rotated — Gamma Inc` +
				`<span class="msp-feed-time">2h ago</span></div>` +
				`<div class="msp-feed-item">` +
				exampleIcon("shield", "", mspIconSize) +
				` 2FA enforced on Delta Co` +
				`<span class="msp-feed-time">5h ago</span></div>` +
				`</div></div>` +

				// Onboarding progress
				`<div class="msp-section">` +
				`<p class="msp-section-title">Onboarding Progress</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-client-row">` +
				`<span class="msp-client-name">NewCo Industries</span>` +
				`<yr-pill>2 of 4 steps</yr-pill></div>` +
				`<div class="msp-client-row">` +
				`<span class="msp-client-name">FreshOrg Ltd</span>` +
				`<yr-pill>Users invited</yr-pill></div>` +
				`</div></div>` +

				`</div>` + // close grid-2

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
