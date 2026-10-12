package ui

import "github.com/a-h/templ"

func portalWorkflowDetailExamples() []Example {
	return []Example{
		{Name: "default", DisplayURL: "platform.yardrail.com/spaces/marketing-ops/workflows/newsletter-digest", Component: templ.Raw(
			portalMockStyle +
				`<yr-page-workspace>` +
				portalNavbar() +
				portalSidebar("space-marketing") +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content><yr-content-main>` +

				// Breadcrumb
				`<nav class="yr-breadcrumb">` +
				`<a href="#">Home</a>` +
				`<span class="yr-breadcrumb-sep">/</span>` +
				`<a href="#">Marketing Ops</a>` +
				`<span class="yr-breadcrumb-sep">/</span>` +
				`<span class="yr-breadcrumb-current">Newsletter Digest</span>` +
				`</nav>` +

				`<h1 class="portal-page-title">Newsletter Digest</h1>` +

				// Status banner
				`<yr-status-banner tone="success" style="margin-bottom:var(--yr-space-6)">` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`<span class="yr-status-banner-message">Running on schedule. Last triggered 2h ago.</span>` +
				`<div class="yr-status-banner-actions">` +
				`<button class="yr-button yr-button--secondary">Pause</button>` +
				`<button class="yr-button yr-button--secondary">Trigger Now</button>` +
				`</div>` +
				`</yr-status-banner>` +

				// Tabs
				`<yr-tabs label="Workflow details">` +
				`<button slot="tab" data-tab="runs">Runs</button>` +
				`<button slot="tab" data-tab="config">Configuration</button>` +

				// --- Runs tab ---
				`<div slot="panel" data-tab="runs">` +

				// Needs attention
				`<yr-section-header>Needs Attention</yr-section-header>` +
				`<a class="yr-activity-row" href="#" style="margin-bottom:var(--yr-space-6)">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Approve newsletter content draft</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-49</span>` +
				`<yr-pill tone="warning">Waiting</yr-pill>` +
				`<span>Content review required</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Waiting</span>` +
				`<yr-live-duration started-at="2026-10-11T12:00:00Z"></yr-live-duration>` +
				`</div>` +
				`</a>` +

				// Recent runs
				`<yr-section-header>Recent Runs</yr-section-header>` +
				`<yr-timeline-entries>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-48</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>3 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">1m 23s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-47</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>3 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">58s</div>` +
				`</a>` +

				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-46</span>` +
				`<yr-pill tone="success">Completed</yr-pill>` +
				`<span>3 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">1m 05s</div>` +
				`</a>` +

				`<a class="yr-activity-row is-error" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-45</span>` +
				`<yr-pill tone="danger">Failed</yr-pill>` +
				`<span>SendGrid 429 — rate limit exceeded</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">12s</div>` +
				`</a>` +

				`</yr-timeline-entries>` +

				`</div>` + // end runs panel

				// --- Configuration tab ---
				`<div slot="panel" data-tab="config">` +

				// Triggers
				`<yr-section-header>Triggers</yr-section-header>` +
				`<div style="display:flex;flex-direction:column;gap:var(--yr-space-2);margin-bottom:var(--yr-space-6)">` +
				`<div class="yr-meta-row">` +
				`<span>Webhook</span>` +
				`<code style="font-family:var(--yr-font-mono);font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted)">https://api.yardrail.com/wh/abc123</code>` +
				`</div>` +
				`<div class="yr-meta-row">` +
				`<span>Manual</span>` +
				`<yr-pill tone="success">Enabled</yr-pill>` +
				`</div>` +
				`</div>` +

				// Connector Bindings
				`<yr-section-header>Connector Bindings</yr-section-header>` +
				`<div style="display:flex;flex-direction:column;gap:var(--yr-space-2);margin-bottom:var(--yr-space-6)">` +
				`<div class="yr-meta-row">` +
				`<span>email.send</span>` +
				`<span>SendGrid</span>` +
				`<yr-pill tone="success" shape="pill">Bound</yr-pill>` +
				`</div>` +
				`<div class="yr-meta-row">` +
				`<span>crm.read</span>` +
				`<span>HubSpot</span>` +
				`<yr-pill tone="success" shape="pill">Bound</yr-pill>` +
				`</div>` +
				`<div class="yr-meta-row">` +
				`<span>file.store</span>` +
				`<span>Google Drive</span>` +
				`<yr-pill tone="success" shape="pill">Bound</yr-pill>` +
				`</div>` +
				`</div>` +

				// Schedule
				`<yr-section-header>Schedule</yr-section-header>` +
				`<div style="margin-bottom:var(--yr-space-6)">` +
				`<div class="yr-meta-row">` +
				`<span>Frequency</span>` +
				`<span>Every weekday at 8:00 AM UTC</span>` +
				`</div>` +
				`</div>` +

				// Inference Backend
				`<yr-section-header>Inference Backend</yr-section-header>` +
				`<div class="yr-meta-row">` +
				`<span>Model</span>` +
				`<span>Claude Sonnet · temperature 0.3 · max 2048 tokens</span>` +
				`</div>` +

				`</div>` + // end config panel

				`</yr-tabs>` +

				`</yr-content-main></yr-content>` +
				`</yr-page-workspace>`,
		)},
	}
}
