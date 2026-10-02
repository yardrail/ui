package ui

import "github.com/a-h/templ"

func actionCardExamples() []Example {
	return []Example{
		{Name: "stacked", Component: templ.Raw(
			`<div style="max-width:40rem">` +
				`<div class="yr-action-card">` +
				`<div class="yr-icon-badge yr-icon-badge--accent">` + exampleIcon("pencil", "", exampleIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Finish setting up</h3>` +
				`<p class="yr-action-card-desc">Complete your organization setup to get started.</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<a href="#" class="yr-button yr-button--secondary">Resume setup</a>` +
				`</div>` +
				`</div>` +
				`<div class="yr-action-card">` +
				`<div class="yr-icon-badge yr-icon-badge--amber">` + exampleIcon("mail", "", exampleIconSize) + `</div>` +
				`<div class="yr-action-card-body">` +
				`<h3 class="yr-action-card-title">Verify your email</h3>` +
				`<p class="yr-action-card-desc">We sent a link to pat@example.com. Purchases stay locked` +
				` until the address is verified.</p>` +
				`</div>` +
				`<div class="yr-action-card-action">` +
				`<a href="#" class="yr-button yr-button--secondary">Resend email</a>` +
				`</div>` +
				`</div>` +
				`</div>`,
		)},
	}
}
