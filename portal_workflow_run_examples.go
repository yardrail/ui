package ui

import "github.com/a-h/templ"

func portalWorkflowRunExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			DisplayURL: "platform.yardrail.com/spaces/marketing-ops/workflows/newsletter-digest/runs/42",
			Component: templ.Raw(
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
					`<a href="#">Newsletter Digest</a>` +
					`<span class="yr-breadcrumb-sep">/</span>` +
					`<span class="yr-breadcrumb-current">Run #42</span>` +
					`</nav>` +

					`<h1 class="portal-page-title">Run #42</h1>` +

					// Run metadata
					`<div class="yr-meta-row" style="margin-bottom:var(--yr-space-6)">` +
					`<yr-pill tone="warning" shape="pill">Waiting</yr-pill>` +
					`<span>Started Oct 11, 2026 at 2:30 PM</span>` +
					`<yr-live-duration started-at="2026-10-11T14:30:00Z"></yr-live-duration>` +
					`<span>Triggered by schedule</span>` +
					`</div>` +

					// HITL Prompt
					`<yr-hitl-prompt tone="warning" style="margin-bottom:var(--yr-space-8)">` +
					`<div class="yr-hitl-prompt-title">Approve the content draft for this week's newsletter?</div>` +
					`<p class="yr-hitl-prompt-desc">The workflow generated a newsletter draft from 12 recent blog ` +
					`posts and 3 product updates. Review the draft in the outputs below before approving.</p>` +
					`<div class="yr-hitl-prompt-actions">` +
					`<button class="yr-button">Approve</button>` +
					`<button class="yr-button yr-button--secondary">Request Changes</button>` +
					`</div>` +
					`</yr-hitl-prompt>` +

					portalWorkflowRunOutputs() +

					`</yr-content-main></yr-content>` +
					`</yr-page-workspace>`,
			),
		},
	}
}

func portalWorkflowRunOutputs() string {
	// Outputs
	return `<yr-section-header>Outputs</yr-section-header>` +
		`<yr-download-list style="margin-bottom:var(--yr-space-8)">` +
		`<a href="#" download>` +
		exampleIcon("file-text", "", portalIconSize) +
		` newsletter-draft.html` +
		`<span class="yr-download-size">84 KB</span>` +
		`</a>` +
		`<a href="#" download>` +
		exampleIcon("file-spreadsheet", "", portalIconSize) +
		` recipient-list.csv` +
		`<span class="yr-download-size">156 KB</span>` +
		`</a>` +
		`</yr-download-list>` +

		// Journal
		`<yr-section-header>Journal</yr-section-header>` +
		`<yr-journal-log>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Run started — triggered by schedule</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-1</span> trigger.schedule</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:00Z">14:30:00</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Fetched 12 blog posts from CMS</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-2</span> content.fetch</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:04Z">14:30:04</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Fetched 3 product updates from Notion</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-3</span> content.fetch</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:09Z">14:30:09</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Generated newsletter draft (1,240 words)</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-4</span> ai.generate</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:18Z">14:30:18</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Queried recipient list — 2,847 subscribers</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-5</span> list.query</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:22Z">14:30:22</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Wrote newsletter-draft.html</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">step-6</span> output.write</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:25Z">14:30:25</time>` +
		`</div>` +

		`<div class="yr-activity-row">` +
		`<div>` +
		`<div class="yr-activity-row-desc">Waiting for approval — content review required</div>` +
		`<div class="yr-activity-row-meta">` +
		`<span class="yr-activity-row-id-badge">wait</span> input.request</div>` +
		`</div>` +
		`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:26Z">14:30:26</time>` +
		`</div>` +

		`</yr-journal-log>`
}
