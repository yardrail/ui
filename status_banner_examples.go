package ui

import "github.com/a-h/templ"

func statusBannerExamples() []Example {
	return []Example{
		{Name: "active", Component: templ.Raw(
			`<div style="max-width:44rem">` +
				`<yr-status-banner tone="success">` +
				`<yr-pill tone="success" shape="pill">Active</yr-pill>` +
				`<span class="yr-status-banner-message">Workflow is running. Last triggered 2h ago.</span>` +
				`<div class="yr-status-banner-actions">` +
				`<button class="yr-button yr-button--secondary">Pause</button>` +
				`</div>` +
				`</yr-status-banner>` +
				`</div>`,
		)},
		{Name: "paused", Component: templ.Raw(
			`<div style="max-width:44rem">` +
				`<yr-status-banner tone="warning">` +
				`<yr-pill tone="warning" shape="pill">Paused</yr-pill>` +
				`<span class="yr-status-banner-message">Workflow is paused. No runs will be triggered.</span>` +
				`<div class="yr-status-banner-actions">` +
				`<button class="yr-button yr-button--secondary">Resume</button>` +
				`</div>` +
				`</yr-status-banner>` +
				`</div>`,
		)},
		{Name: "failed", Component: templ.Raw(
			`<div style="max-width:44rem">` +
				`<yr-status-banner tone="danger">` +
				`<yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				`<span class="yr-status-banner-message">Last run failed: connector returned 401.</span>` +
				`<div class="yr-status-banner-actions">` +
				`<button class="yr-button yr-button--secondary">Retry</button>` +
				`</div>` +
				`</yr-status-banner>` +
				`</div>`,
		)},
	}
}
