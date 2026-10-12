package ui

import "github.com/a-h/templ"

func portalWorkflowsExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/automations/workflows", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebarAutomations("workflows") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +

				// Title row
				`<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:var(--yr-space-6)">` +
				`<h1 class="portal-page-title" style="margin-bottom:0">Workflows</h1>` +
				`<button class="yr-button">+ New Workflow</button>` +
				`</div>` +

				// Search
				`<form class="yr-search yr-search--compact" method="get" action="#" role="search"` +
				` style="max-width:24rem;margin-bottom:var(--yr-space-4)">` +
				exampleIcon("search", `class="yr-search-icon"`, searchCompactIconSize) +
				`<input class="yr-search-input" type="search" name="q" placeholder="Search workflows…" autocomplete="off">` +
				`</form>` +

				// Stats row
				`<div class="portal-grid portal-grid-4" style="margin-bottom:var(--yr-space-6)">` +
				`<div class="portal-stat-card">` +
				`<p class="portal-stat-value">14</p>` +
				`<p class="portal-stat-label">Total</p></div>` +
				`<div class="portal-stat-card" style="border-left:3px solid var(--yr-success)">` +
				`<p class="portal-stat-value">10</p>` +
				`<p class="portal-stat-label">Active</p></div>` +
				`<div class="portal-stat-card" style="border-left:3px solid var(--yr-warning)">` +
				`<p class="portal-stat-value">3</p>` +
				`<p class="portal-stat-label">Paused</p></div>` +
				`<div class="portal-stat-card" style="border-left:3px solid var(--yr-danger)">` +
				`<p class="portal-stat-value">1</p>` +
				`<p class="portal-stat-label">Error</p></div>` +
				`</div>` +

				// Workflow rows
				`<div style="display:flex;flex-direction:column;gap:var(--yr-space-3)">` +

				// Newsletter Digest — Active
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Newsletter Digest</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Marketing Ops</span>` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`<span>Last run 2h ago</span>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Trigger</button>` +
				`</div></div>` +

				// Lead Scoring — Paused
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Lead Scoring</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Sales Ops</span>` +
				`<yr-pill shape="pill">Paused</yr-pill>` +
				`<span>Last run 3d ago</span>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Resume</button>` +
				`</div></div>` +

				// Invoice Sync — Active
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Invoice Sync</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Client Onboarding</span>` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`<span>Last run 15m ago</span>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Trigger</button>` +
				`</div></div>` +

				// Social Scheduler — Error
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Social Scheduler</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Marketing Ops</span>` +
				`<yr-pill tone="danger" shape="pill">Error</yr-pill>` +
				`<span>Failed 1h ago</span>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">View</button>` +
				`</div></div>` +

				// Onboarding Flow — Draft
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Onboarding Flow</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Client Onboarding</span>` +
				`<yr-pill tone="warning" shape="pill">Draft</yr-pill>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Edit</button>` +
				`</div></div>` +

				// Ticket Router — Active
				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` + exampleIcon("git-branch", "", portalIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Ticket Router</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>My Space</span>` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`<span>Last run 30m ago</span>` +
				`</div></div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Trigger</button>` +
				`</div></div>` +

				`</div>` + // close workflow rows

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
