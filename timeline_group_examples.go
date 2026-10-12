package ui

import "github.com/a-h/templ"

func timelineGroupExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<div style="max-width:44rem">` +
				`<yr-timeline-group aria-label="Today">` +
				`<time datetime="2026-10-11">Today</time>` +
				`<yr-timeline-entries>` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-45</span> Completed in 1m 23s</div>` +
				`</div>` +
				`</div>` +
				`<a class="yr-activity-row" href="#">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Lead Scoring</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-44</span> Waiting for input</div>` +
				`</div>` +
				`</a>` +
				`</yr-timeline-entries>` +
				`</yr-timeline-group>` +

				`<yr-timeline-group aria-label="Yesterday">` +
				`<time datetime="2026-10-10">Yesterday</time>` +
				`<yr-timeline-entries>` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Newsletter Digest</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-43</span> Completed in 58s</div>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row is-error">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Social Scheduler</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-42</span> Rate limit exceeded</div>` +
				`</div>` +
				`</div>` +
				`<div class="yr-activity-row">` +
				`<div>` +
				`<div class="yr-activity-row-desc">Lead Scoring</div>` +
				`<div class="yr-activity-row-meta">` +
				`<span class="yr-activity-row-id-badge">R-41</span> Completed in 4m 12s</div>` +
				`</div>` +
				`</div>` +
				`</yr-timeline-entries>` +
				`</yr-timeline-group>` +
				`</div>`,
		)},
	}
}
