package ui

import "github.com/a-h/templ"

func runSummaryRowExamples() []Example {
	return []Example{
		{Name: "states", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:44rem">` +

				// Success
				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-45</span>` +
				` <yr-pill tone="success" shape="pill">Completed</yr-pill>` +
				` <span>3 outputs</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">1m 23s</span>` +
				`</div>` +
				`</a>` +

				// Failed
				`<a class="yr-activity-row is-error" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Social Scheduler</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-42</span>` +
				` <yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				` <span>Rate limit exceeded</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">12s</span>` +
				`</div>` +
				`</a>` +

				// Waiting
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

				`</div>`,
		), DisplayURL: ""},
	}
}
