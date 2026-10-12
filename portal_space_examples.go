package ui

import "github.com/a-h/templ"

const spaceDetailStyle = `<style>` +
	`.space-header{display:flex;justify-content:space-between;` +
	`align-items:flex-start;margin-bottom:var(--yr-space-6)}` +
	`.space-header-title{display:flex;align-items:center;gap:var(--yr-space-3)}` +
	`.space-header-meta{display:flex;align-items:center;gap:var(--yr-space-4)}` +
	`.space-header-members{display:flex;align-items:center;gap:var(--yr-space-1)}` +
	`.space-header-member-count{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-xs);color:var(--yr-text-muted);` +
	`margin-left:var(--yr-space-2)}` +
	`</style>`

func portalSpaceExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/spaces/marketing-ops", Component: templ.Raw(
			portalMockStyle + spaceDetailStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebar("space-marketing") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +

				// Space header
				`<div class="space-header">` +
				`<div class="space-header-title">` +
				`<h1 class="portal-page-title" style="margin-bottom:0">Marketing Ops</h1>` +
				`</div>` +
				`<div class="space-header-meta">` +
				`<div class="space-header-members">` +
				`<yr-avatar>AL</yr-avatar>` +
				`<yr-avatar>MK</yr-avatar>` +
				`<yr-avatar>JS</yr-avatar>` +
				`<span class="space-header-member-count">3 members</span>` +
				`</div>` +
				`<button class="yr-button yr-button--secondary">Settings</button>` +
				`</div>` +
				`</div>` +

				// Installed Workflows
				`<yr-section-header>Installed Workflows <a href="#">Install new</a></yr-section-header>` +
				`<div class="portal-grid portal-grid-3" style="margin-bottom:var(--yr-space-8)">` +

				// Lead Scoring — waiting (needs-attention first)
				`<yr-workflow-card>` +
				`<div class="yr-workflow-card-header">` +
				`<span class="yr-workflow-card-name">Lead Scoring</span>` +
				`<yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
				`</div>` +
				`<div class="yr-workflow-card-stats">` +
				`<yr-stat number="42" label="Runs"></yr-stat>` +
				`<yr-stat number="91%" label="Success"></yr-stat>` +
				`</div>` +
				`<div class="yr-workflow-card-footer">` +
				`<span>Waiting for input</span>` +
				`<a href="#" class="yr-button yr-button--secondary">View</a>` +
				`</div>` +
				`</yr-workflow-card>` +

				// Newsletter Digest — active
				`<yr-workflow-card>` +
				`<div class="yr-workflow-card-header">` +
				`<span class="yr-workflow-card-name">Newsletter Digest</span>` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`</div>` +
				`<div class="yr-workflow-card-stats">` +
				`<yr-stat number="127" label="Runs"></yr-stat>` +
				`<yr-stat number="98%" label="Success"></yr-stat>` +
				`</div>` +
				`<div class="yr-workflow-card-footer">` +
				`<span>Last run 2h ago</span>` +
				`<button class="yr-button yr-button--secondary">Trigger</button>` +
				`</div>` +
				`</yr-workflow-card>` +

				// Social Scheduler — paused
				`<yr-workflow-card>` +
				`<div class="yr-workflow-card-header">` +
				`<span class="yr-workflow-card-name">Social Scheduler</span>` +
				`<yr-pill shape="pill">Paused</yr-pill>` +
				`</div>` +
				`<div class="yr-workflow-card-stats">` +
				`<yr-stat number="18" label="Runs"></yr-stat>` +
				`<yr-stat number="85%" label="Success"></yr-stat>` +
				`</div>` +
				`<div class="yr-workflow-card-footer">` +
				`<span>Last run 1d ago</span>` +
				`<button class="yr-button yr-button--secondary">Resume</button>` +
				`</div>` +
				`</yr-workflow-card>` +

				`</div>` + // close grid

				// Recent Activity
				`<yr-section-header>Recent Activity <a href="#">View all</a></yr-section-header>` +
				`<yr-timeline-entries>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-45</span>` +
				` <yr-pill tone="success" shape="pill">Completed</yr-pill>` +
				` <span>3 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">1m 23s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Lead Scoring</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-44</span>` +
				` <yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
				` <span>Approve budget &gt;$5k</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Waiting</span>` +
				`<yr-live-duration started-at="2026-10-11T08:00:00Z"></yr-live-duration>` +
				`</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-43</span>` +
				` <yr-pill tone="success" shape="pill">Completed</yr-pill>` +
				` <span>2 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">58s</div>` +
				`</a>` +

				`<a class="yr-activity-row is-error" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Social Scheduler</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-42</span>` +
				` <yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				` <span>Rate limit exceeded</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">12s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Lead Scoring</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-41</span>` +
				` <yr-pill tone="success" shape="pill">Completed</yr-pill>` +
				` <span>1 output</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">4m 12s</div>` +
				`</a>` +

				`</yr-timeline-entries>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
