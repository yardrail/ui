package ui

import "github.com/a-h/templ"

func workflowCardExamples() []Example {
	return []Example{
		{Name: "active", Component: templ.Raw(
			`<div style="max-width:22rem">` +
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
				`</div>`,
		), DisplayURL: ""},
		{Name: "paused", Component: templ.Raw(
			`<div style="max-width:22rem">` +
				`<yr-workflow-card>` +
				`<div class="yr-workflow-card-header">` +
				`<span class="yr-workflow-card-name">Lead Scoring</span>` +
				`<yr-pill shape="pill">Paused</yr-pill>` +
				`</div>` +
				`<div class="yr-workflow-card-stats">` +
				`<yr-stat number="42" label="Runs"></yr-stat>` +
				`<yr-stat number="91%" label="Success"></yr-stat>` +
				`</div>` +
				`<div class="yr-workflow-card-footer">` +
				`<span>Last run 3d ago</span>` +
				`<button class="yr-button yr-button--secondary">Resume</button>` +
				`</div>` +
				`</yr-workflow-card>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
