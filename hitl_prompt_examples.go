package ui

import "github.com/a-h/templ"

func hitlPromptExamples() []Example {
	return []Example{
		{Name: "approve-reject", Component: templ.Raw(
			`<div style="max-width:36rem">` +
				`<yr-hitl-prompt tone="warning">` +
				`<div class="yr-hitl-prompt-title">Approve the refund for order #1042?</div>` +
				`<p class="yr-hitl-prompt-desc">Customer requested a $84.50 refund. ` +
				`Order was delivered 3 days ago with no reported issues.</p>` +
				`<div class="yr-hitl-prompt-actions">` +
				`<button class="yr-button">Approve</button>` +
				`<button class="yr-button yr-button--secondary">Reject</button>` +
				`</div>` +
				`</yr-hitl-prompt>` +
				`</div>`,
		)},
		{Name: "danger", Component: templ.Raw(
			`<div style="max-width:36rem">` +
				`<yr-hitl-prompt tone="danger">` +
				`<div class="yr-hitl-prompt-title">Critical: Slack API key expired</div>` +
				`<p class="yr-hitl-prompt-desc">The Slack connector for Marketing Bot ` +
				`lost authentication. 3 workflows are blocked until reconnected.</p>` +
				`<div class="yr-hitl-prompt-actions">` +
				`<button class="yr-button">Reconnect</button>` +
				`</div>` +
				`</yr-hitl-prompt>` +
				`</div>`,
		)},
	}
}
