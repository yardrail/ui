package ui

import "github.com/a-h/templ"

func activityRowExamples() []Example {
	return []Example{
		{Name: "rows", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:40rem">` +
				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Approve the refund for order 1042?</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">S-218</span> Refund triage</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Waiting</span>` +
				`<yr-live-duration started-at="2026-10-01T09:00:00Z"></yr-live-duration>` +
				`</div>` +
				`</a>` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Weekly pipeline report</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">S-217</span> Sales reporting</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Running</span>` +
				`<yr-live-duration started-at="2026-10-01T12:30:00Z"></yr-live-duration>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row is-error">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Sync contacts to the CRM</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">S-214</span> Connector returned 401</div>` +
				`</div>` +
				`<div class="yr-activity-row-duration">` +
				`<span class="yr-activity-row-duration-label">Failed</span>` +
				`<yr-live-duration></yr-live-duration>` +
				`</div>` +
				`</div>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: "journal", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:40rem">` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Run started — triggered by webhook</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">step-1</span> trigger.webhook</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:00Z">14:30:00</time>` +
				`</div>` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Fetched 24 records from HubSpot CRM</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">step-2</span> crm.query</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:12Z">14:30:12</time>` +
				`</div>` +
				`<div class="yr-activity-row is-error">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Connector authentication failed</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">step-3</span> connector.auth</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:15Z">14:30:15</time>` +
				`</div>` +
				`</div>`,
		)},
	}
}
