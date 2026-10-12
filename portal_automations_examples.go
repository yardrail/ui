package ui

import "github.com/a-h/templ"

func portalAutomationsExamples() []Example {
	return []Example{
		{Name: exampleDefault, DisplayURL: "platform.yardrail.com/automations", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebar("automations") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="portal-page-title">Automations</h1>` +

				// Navigation tiles
				`<div class="portal-grid portal-grid-3" style="margin-bottom:var(--yr-space-8)">` +
				`<div class="yr-action-card">` +
				`<div class="yr-icon-badge yr-icon-badge--accent">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Workflows</h3>` +
				`<p class="yr-action-card-desc">7 active workflows across 4 spaces</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<a href="#" class="yr-button yr-button--secondary">View all</a>` +
				`</div>` +
				`</div>` +
				`<div class="yr-action-card">` +
				`<div class="yr-icon-badge yr-icon-badge--purple">` + exampleIcon("layout-grid", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Template Catalog</h3>` +
				`<p class="yr-action-card-desc">Browse 500+ pre-built automation templates</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<a href="#" class="yr-button yr-button--secondary">Browse</a>` +
				`</div>` +
				`</div>` +
				`<div class="yr-action-card">` +
				`<div class="yr-icon-badge yr-icon-badge--amber">` + exampleIcon("plug", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Connectors</h3>` +
				`<p class="yr-action-card-desc">4 tools connected to your organization</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<a href="#" class="yr-button yr-button--secondary">Manage</a>` +
				`</div>` +
				`</div>` +
				`</div>` +

				// Connected Tools
				`<yr-section-header>Connected Tools <a href="#">Manage</a></yr-section-header>` +
				`<div style="display:flex;gap:var(--yr-space-3);margin-bottom:var(--yr-space-8)">` +
				`<div class="yr-connector-pill--icon" title="Slack">` +
				`<img src="` + exampleLogo + `" alt="Slack" width="24" height="24"></div>` +
				`<div class="yr-connector-pill--icon" title="HubSpot">HS</div>` +
				`<div class="yr-connector-pill--icon" title="Google Sheets">GS</div>` +
				`<div class="yr-connector-pill--icon" title="Stripe">ST</div>` +
				`</div>` +

				// Featured Templates
				`<yr-section-header>Featured Templates <a href="#">View all</a></yr-section-header>` +
				`<div class="portal-grid portal-grid-2">` +
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Refund triage</h3>` +
				`<p class="yr-template-card-desc">Routes refund requests to the right approver and posts` +
				` the decision back to the customer thread.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="Slack"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`<div class="yr-connector-pill--icon" title="HubSpot">HS</div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="accent">Verified</yr-pill>` +
				`<yr-rating value="4.5" reviews="42"></yr-rating>` +
				`</div>` +
				`<span class="yr-template-card-price">$24</span>` +
				`</div>` +
				`</a>` +
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Weekly pipeline report</h3>` +
				`<p class="yr-template-card-desc">Summarizes open deals every Monday and sends a` +
				` digest to the sales team channel.</p>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="purple">Community</yr-pill>` +
				`<span class="yr-template-card-install-count">1.2k installs</span>` +
				`</div>` +
				`<span class="yr-template-card-price">Free</span>` +
				`</div>` +
				`</a>` +
				`</div>` +
				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
