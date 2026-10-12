package ui

import "github.com/a-h/templ"

func portalTemplatesExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/automations/templates", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebarAutomations("templates") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="portal-page-title">Template Catalog</h1>` +

				// Hero search
				`<form class="yr-search yr-search--hero" method="get" action="#" role="search"` +
				` style="margin-bottom:var(--yr-space-4)">` +
				exampleIcon("search", `class="yr-search-icon"`, searchHeroIconSize) +
				`<input class="yr-search-input" type="search" name="q"` +
				` placeholder="Search 500+ templates…" autocomplete="off">` +
				`</form>` +

				// Category filter pills
				`<div style="display:flex;gap:var(--yr-space-2);margin-bottom:var(--yr-space-6)">` +
				`<yr-pill tone="accent" class="is-active">All</yr-pill>` +
				`<yr-pill>Marketing</yr-pill>` +
				`<yr-pill>Sales</yr-pill>` +
				`<yr-pill>Operations</yr-pill>` +
				`<yr-pill>HR</yr-pill>` +
				`<yr-pill>Finance</yr-pill>` +
				`</div>` +

				// Template grid
				`<div class="portal-grid portal-grid-3">` +

				// Refund Triage
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

				// Weekly Pipeline Report
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Weekly pipeline report</h3>` +
				`<p class="yr-template-card-desc">Summarizes open deals and pipeline movement every Monday` +
				` morning in a Slack digest.</p>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="purple">Community</yr-pill>` +
				`<span class="yr-template-card-install-count">1.2k installs</span>` +
				`</div>` +
				`<span class="yr-template-card-price">Free</span>` +
				`</div>` +
				`</a>` +

				// New Hire Onboarding
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">New hire onboarding</h3>` +
				`<p class="yr-template-card-desc">Creates accounts, sends welcome emails, and schedules` +
				` orientation for new team members.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="Google">GC</div>` +
				`<div class="yr-connector-pill--icon" title="Slack"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="accent">Verified</yr-pill>` +
				`<yr-rating value="4.8" reviews="128"></yr-rating>` +
				`</div>` +
				`<span class="yr-template-card-price">$18</span>` +
				`</div>` +
				`</a>` +

				// Lead Enrichment
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Lead enrichment</h3>` +
				`<p class="yr-template-card-desc">Enriches inbound leads with firmographic data and` +
				` scores them for sales readiness.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="HubSpot">HS</div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="purple">Community</yr-pill>` +
				`<span class="yr-template-card-install-count">890 installs</span>` +
				`</div>` +
				`<span class="yr-template-card-price">Free</span>` +
				`</div>` +
				`</a>` +

				// Invoice Reconciliation
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Invoice reconciliation</h3>` +
				`<p class="yr-template-card-desc">Matches invoices against purchase orders and flags` +
				` discrepancies for review.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="Stripe">ST</div>` +
				`<div class="yr-connector-pill--icon" title="Google Sheets">GS</div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="accent">Verified</yr-pill>` +
				`<yr-rating value="4.2" reviews="31"></yr-rating>` +
				`</div>` +
				`<span class="yr-template-card-price">$36</span>` +
				`</div>` +
				`</a>` +

				// Social Listening Digest
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Social listening digest</h3>` +
				`<p class="yr-template-card-desc">Aggregates brand mentions across social platforms` +
				` into a weekly Slack summary.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="Slack"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="purple">Community</yr-pill>` +
				`<span class="yr-template-card-install-count">2.1k installs</span>` +
				`</div>` +
				`<span class="yr-template-card-price">Free</span>` +
				`</div>` +
				`</a>` +

				`</div>` + // close portal-grid
				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
