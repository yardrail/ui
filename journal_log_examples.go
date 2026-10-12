package ui

import "github.com/a-h/templ"

func journalLogExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: templ.Raw(
			`<div style="max-width:44rem">` +
				`<yr-journal-log>` +

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

				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Filtered to 8 leads matching criteria</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">step-3</span> filter.apply</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:13Z">14:30:13</time>` +
				`</div>` +

				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Sent summary to #sales-leads channel</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">step-4</span> chat.post</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:18Z">14:30:18</time>` +
				`</div>` +

				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Run completed — 1 output generated</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">done</span> run.complete</div>` +
				`</div>` +
				`<time class="yr-activity-row-timestamp" datetime="2026-10-11T14:30:19Z">14:30:19</time>` +
				`</div>` +

				`</yr-journal-log>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
