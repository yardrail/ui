package ui

import "github.com/a-h/templ"

const workflowRowIconSize = 20

func workflowRowExamples() []Example {
	return []Example{
		{Name: exampleRows, Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:44rem">` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("git-branch", "", workflowRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Newsletter Digest</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Marketing Ops</span>` +
				` <yr-pill tone="success" shape="pill">Active</yr-pill>` +
				` <span>Last run 2h ago</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Trigger</button>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("git-branch", "", workflowRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Lead Scoring</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Sales Ops</span>` +
				` <yr-pill shape="pill">Paused</yr-pill>` +
				` <span>Last run 3d ago</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Resume</button>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("git-branch", "", workflowRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Social Scheduler</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Marketing Ops</span>` +
				` <yr-pill tone="danger" shape="pill">Error</yr-pill>` +
				` <span>Failed 1h ago</span>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">View</button>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("git-branch", "", workflowRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">Onboarding Flow</div>` +
				`<div class="yr-list-row-meta">` +
				`<span>Client Onboarding</span>` +
				` <yr-pill tone="warning" shape="pill">Draft</yr-pill>` +
				`</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<button class="yr-button yr-button--secondary">Edit</button>` +
				`</div>` +
				`</div>` +

				`</div>`,
		), DisplayURL: ""},
	}
}
