package ui

import "github.com/a-h/templ"

func mspClientOverviewExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp/clients/acme-corp", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebarClientDetail("overview") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="msp-page-title">Acme Corp</h1>` +

				`<div class="msp-grid msp-grid-3" style="margin-bottom:var(--yr-space-8)">` +
				`<div class="msp-stat-card">` +
				`<p class="msp-stat-value">8</p>` +
				`<p class="msp-stat-label">Users</p></div>` +
				`<div class="msp-stat-card">` +
				`<p class="msp-stat-value">3</p>` +
				`<p class="msp-stat-label">Connectors</p></div>` +
				`<div class="msp-stat-card">` +
				`<p class="msp-stat-value">12</p>` +
				`<p class="msp-stat-label">Workflows</p></div>` +
				`</div>` +

				`<div class="msp-grid msp-grid-2">` +
				`<div class="msp-section">` +
				`<p class="msp-section-title">Recent Activity</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-feed-item">` +
				exampleIcon("user-plus", "", mspIconSize) +
				` Jane Smith invited` +
				`<span class="msp-feed-time">2h ago</span></div>` +
				`<div class="msp-feed-item">` +
				exampleIcon("play", "", mspIconSize) +
				` Invoice sync deployed` +
				`<span class="msp-feed-time">1d ago</span></div>` +
				`<div class="msp-feed-item">` +
				exampleIcon("refresh-cw", "", mspIconSize) +
				` Slack credential rotated` +
				`<span class="msp-feed-time">3d ago</span></div>` +
				`</div></div>` +

				`<div class="msp-section">` +
				`<p class="msp-section-title">Alerts</p>` +
				`<div class="msp-feed">` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="danger">critical</yr-pill>` +
				` Slack OAuth expired` +
				`<span class="msp-feed-time">2m ago</span></div>` +
				`<div class="msp-feed-item">` +
				`<yr-pill tone="info">info</yr-pill>` +
				` 2 workflows paused` +
				`<span class="msp-feed-time">1h ago</span></div>` +
				`</div></div>` +
				`</div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}

func mspClientUsersExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/users", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebarClientDetail("users") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center;margin-bottom:var(--yr-space-6)">` +
				`<h1 class="msp-page-title" style="margin:0">` +
				`Acme Corp — Users</h1>` +
				`<yr-button variant="primary">` +
				exampleIcon("user-plus", "", mspIconSize) +
				` Invite User</yr-button></div>` +

				`<table class="yr-table" style="width:100%">` +
				`<thead><tr>` +
				`<th>Name</th><th>Email</th><th>Role</th>` +
				`<th>Spaces</th><th>Status</th></tr></thead>` +
				`<tbody>` +
				`<tr><td>Alex Johnson</td>` +
				`<td>alex@acme.com</td>` +
				`<td><yr-pill>Admin</yr-pill></td>` +
				`<td>All spaces</td>` +
				`<td><yr-pill tone="success">Active</yr-pill></td></tr>` +
				`<tr><td>Maria Chen</td>` +
				`<td>maria@acme.com</td>` +
				`<td><yr-pill>Editor</yr-pill></td>` +
				`<td>Sales, Support</td>` +
				`<td><yr-pill tone="success">Active</yr-pill></td></tr>` +
				`<tr><td>Sam Patel</td>` +
				`<td>sam@acme.com</td>` +
				`<td><yr-pill>Viewer</yr-pill></td>` +
				`<td>Sales</td>` +
				`<td><yr-pill tone="success">Active</yr-pill></td></tr>` +
				`<tr><td>Taylor Kim</td>` +
				`<td>taylor@acme.com</td>` +
				`<td><yr-pill>Editor</yr-pill></td>` +
				`<td>Ops, Finance</td>` +
				`<td><yr-pill>Invited</yr-pill></td></tr>` +
				`</tbody></table>` +

				`<div class="msp-section" style="margin-top:var(--yr-space-8)">` +
				`<p class="msp-section-title">SCIM Sync</p>` +
				`<div class="msp-stat-card">` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center">` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text);margin:0">Azure AD</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:var(--yr-space-1) 0 0">` +
				`Last synced 4 hours ago · 8 users</p></div>` +
				`<yr-pill tone="success">Connected</yr-pill>` +
				`</div></div></div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}

func mspClientConnectorsExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/connectors", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebarClientDetail("connectors") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center;margin-bottom:var(--yr-space-6)">` +
				`<h1 class="msp-page-title" style="margin:0">` +
				`Acme Corp — Connectors</h1>` +
				`<yr-button variant="primary">` +
				exampleIcon("plug", "", mspIconSize) +
				` Add Connector</yr-button></div>` +

				`<div class="msp-feed">` +

				`<div class="msp-client-row">` +
				`<div style="display:flex;align-items:center;` +
				`gap:var(--yr-space-3);flex:1">` +
				exampleIcon("message-square", "", mspIconSize) +
				`<div><span class="msp-client-name">Slack</span>` +
				`<span class="msp-client-meta" style="display:block">` +
				`OAuth · Used by 4 workflows</span></div></div>` +
				`<yr-pill tone="danger">Expired</yr-pill></div>` +

				`<div class="msp-client-row">` +
				`<div style="display:flex;align-items:center;` +
				`gap:var(--yr-space-3);flex:1">` +
				exampleIcon("database", "", mspIconSize) +
				`<div><span class="msp-client-name">HubSpot</span>` +
				`<span class="msp-client-meta" style="display:block">` +
				`API Key · Used by 6 workflows</span></div></div>` +
				`<yr-pill tone="success">Healthy</yr-pill></div>` +

				`<div class="msp-client-row">` +
				`<div style="display:flex;align-items:center;` +
				`gap:var(--yr-space-3);flex:1">` +
				exampleIcon("credit-card", "", mspIconSize) +
				`<div><span class="msp-client-name">Stripe</span>` +
				`<span class="msp-client-meta" style="display:block">` +
				`API Key · Used by 2 workflows</span></div></div>` +
				`<yr-pill tone="success">Healthy</yr-pill></div>` +

				`</div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}

func mspClientSettingsExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/settings", Component: templ.Raw(
			mspMockStyle +
				`<yr-page-workspace>` +
				mspNavbar() +
				mspSidebarClientDetail("settings") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="msp-page-title">Acme Corp — Settings</h1>` +

				`<div class="msp-section">` +
				`<p class="msp-section-title">Authentication</p>` +
				`<div class="msp-stat-card" style="margin-bottom:var(--yr-space-3)">` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center">` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);font-weight:500;` +
				`color:var(--yr-text);margin:0">` +
				`Two-Factor Authentication</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:var(--yr-space-1) 0 0">` +
				`Require 2FA for all users in this organization</p></div>` +
				`<yr-pill tone="success">Required</yr-pill>` +
				`</div></div>` +

				`<div class="msp-stat-card" style="margin-bottom:var(--yr-space-3)">` +
				`<div style="display:flex;justify-content:space-between;` +
				`align-items:center">` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);font-weight:500;` +
				`color:var(--yr-text);margin:0">` +
				`Single Sign-On</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:var(--yr-space-1) 0 0">` +
				`SAML/OIDC integration with client identity provider</p></div>` +
				`<yr-pill>Not configured</yr-pill>` +
				`</div></div>` +
				`</div>` +

				`<div class="msp-section">` +
				`<p class="msp-section-title">Session Policies</p>` +
				`<div class="msp-stat-card">` +
				`<div style="display:flex;gap:var(--yr-space-8)">` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:0">Idle Timeout</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);font-weight:500;` +
				`color:var(--yr-text);margin:var(--yr-space-1) 0 0">` +
				`30 minutes</p></div>` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:0">Max Session</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);font-weight:500;` +
				`color:var(--yr-text);margin:var(--yr-space-1) 0 0">` +
				`8 hours</p></div>` +
				`<div>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted);margin:0">IP Allowlist</p>` +
				`<p style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);font-weight:500;` +
				`color:var(--yr-text);margin:var(--yr-space-1) 0 0">` +
				`Disabled</p></div>` +
				`</div></div>` +
				`</div>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
