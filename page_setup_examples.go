package ui

import "github.com/a-h/templ"

const pageSetupBase = `<style>html,body{margin:0;height:100%}` +
	`.demo-region{display:flex;align-items:center;justify-content:center;` +
	`font-family:var(--yr-font-mono);font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-subtle)}</style>`

func pageSetupExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			pageSetupBase +
				`<yr-page-setup>` +
				`<yr-navbar>` +
				`<div class="demo-region" style="width:100%">nav</div>` +
				`</yr-navbar>` +
				`<yr-page-setup-content>` +
				`<div class="demo-region" style="width:720px;height:400px;` +
				`border:1px dashed var(--yr-border);` +
				`border-radius:var(--yr-radius-md)">` +
				`content` +
				`</div>` +
				`</yr-page-setup-content>` +
				`</yr-page-setup>`,
		)},
		{Name: "with-setup-card", Component: templ.Raw(
			pageSetupBase +
				`<yr-page-setup>` +
				`<yr-navbar-std logo="/logo.png" brand="Yardrail">` +
				`</yr-navbar-std>` +
				`<yr-page-setup-content>` +
				`<yr-setup-card>` +
				`<yr-setup-card-progress steps="3" current="1">` +
				`</yr-setup-card-progress>` +
				`<yr-setup-card-header>` +
				`<h2 class="yr-setup-card-title">` +
				`Create your organization</h2>` +
				`<p class="yr-setup-card-subtitle">` +
				`This is the workspace where your team will` +
				` collaborate.</p>` +
				`</yr-setup-card-header>` +
				`<yr-setup-card-body>` +
				`<div class="yr-field">` +
				`<label for="ps-name">Organization name</label>` +
				` <input type="text" id="ps-name" name="name"` +
				` placeholder="Acme Corp"></div>` +
				`</yr-setup-card-body>` +
				`<yr-setup-card-footer>` +
				`<span></span>` +
				`<yr-button variant="primary">Continue</yr-button>` +
				`</yr-setup-card-footer>` +
				`</yr-setup-card>` +
				`</yr-page-setup-content>` +
				`</yr-page-setup>`,
		)},
	}
}
