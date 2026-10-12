package ui

import "github.com/a-h/templ"

func portalHomeExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebar("home") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +
				`<h1 class="portal-page-title">Home</h1>` +

				// Stats row
				`<div class="portal-grid portal-grid-4" style="margin-bottom:var(--yr-space-8)">` +
				`<div class="portal-stat-card">` +
				`<p class="portal-stat-value">8</p>` +
				`<p class="portal-stat-label">Active Workflows</p></div>` +
				`<div class="portal-stat-card" style="border-left:3px solid var(--yr-warning)">` +
				`<p class="portal-stat-value">3</p>` +
				`<p class="portal-stat-label">Waiting on You</p></div>` +
				`<div class="portal-stat-card">` +
				`<p class="portal-stat-value">127</p>` +
				`<p class="portal-stat-label">Runs This Week</p></div>` +
				`<div class="portal-stat-card">` +
				`<p class="portal-stat-value">96%</p>` +
				`<p class="portal-stat-label">Success Rate</p></div>` +
				`</div>` +

				// Needs You
				`<yr-section-header>Needs You <a href="#">View all</a></yr-section-header>` +
				`<div style="margin-bottom:var(--yr-space-8)">` +
				`<div class="yr-action-card yr-action-card--urgent">` +
				`<div class="yr-icon-badge yr-icon-badge--amber">` + exampleIcon("clock", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Approve budget over $5k for Acme lead</h3>` +
				`<p class="yr-action-card-desc">Waiting 2h 14m · Lead Scoring · Marketing Ops</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
				`</div></div>` +
				`<div class="yr-action-card yr-action-card--urgent">` +
				`<div class="yr-icon-badge yr-icon-badge--amber">` + exampleIcon("clock", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Review generated newsletter draft</h3>` +
				`<p class="yr-action-card-desc">Waiting 45m · Newsletter Digest · Marketing Ops</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
				`</div></div>` +
				`</div>` +

				// Problems
				`<yr-section-header>Problems</yr-section-header>` +
				`<div style="margin-bottom:var(--yr-space-8)">` +
				`<div class="yr-action-card yr-action-card--urgent-danger">` +
				`<div class="yr-icon-badge yr-icon-badge--danger">` + exampleIcon("alert-triangle", "", portalIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Social Scheduler failed</h3>` +
				`<p class="yr-action-card-desc">Rate limit exceeded · 3 consecutive failures · 1h ago</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				`</div></div>` +
				`</div>` +

				// Recent Activity
				`<yr-section-header>Recent Activity <a href="#">See more</a></yr-section-header>` +
				`<yr-timeline-entries>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-48</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>3 outputs</span></div></div>` +
				`<div class="yr-activity-row-duration">1m 23s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Lead Scoring</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-47</span>` +
				`<yr-pill tone="warning">Waiting</yr-pill>` +
				`<span>Approve budget &gt;$5k</span></div></div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Waiting</span>` +
				`<yr-live-duration started-at="2026-10-01T09:00:00Z"></yr-live-duration>` +
				`</div></a>` +

				`<div class="yr-activity-row is-error">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Social Scheduler</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-46</span>` +
				`<yr-pill tone="danger">Failed</yr-pill>` +
				`<span>Rate limit exceeded</span></div></div>` +
				`<div class="yr-activity-row-duration">12s</div>` +
				`</div>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Contact Sync</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-45</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>24 records synced</span></div></div>` +
				`<div class="yr-activity-row-duration">4m 8s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Invoice Processor</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-44</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>12 invoices</span></div></div>` +
				`<div class="yr-activity-row-duration">2m 51s</div>` +
				`</a>` +

				`</yr-timeline-entries>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
