package ui

import "github.com/a-h/templ"

const setupSlotStyle = `<style>.demo-slot{background:var(--yr-bg);` +
	`border:1px dashed var(--yr-border);` +
	`border-radius:var(--yr-radius-sm);` +
	`padding:var(--yr-space-1) var(--yr-space-3);` +
	`color:var(--yr-text-disabled);` +
	`font-family:var(--yr-font-mono);` +
	`font-size:var(--yr-font-size-xs)}</style>`

func setupCardExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			setupSlotStyle +
				`<yr-setup-card>` +
				`<yr-setup-card-progress steps="4" current="1">` +
				`</yr-setup-card-progress>` +
				`<yr-setup-card-header>` +
				`<div class="demo-slot">header</div>` +
				`</yr-setup-card-header>` +
				`<yr-setup-card-body>` +
				`<div class="demo-slot">body</div>` +
				`</yr-setup-card-body>` +
				`<yr-setup-card-footer>` +
				`<div class="demo-slot">footer</div>` +
				`</yr-setup-card-footer>` +
				`</yr-setup-card>`,
		)},
		{Name: "form-step", Component: templ.Raw(
			`<yr-setup-card>` +
				`<yr-setup-card-progress steps="4" current="1">` +
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
				`<label for="sc-name">Organization name</label>` +
				` <input type="text" id="sc-name" name="name"` +
				` placeholder="Acme Corp"></div>` +
				`</yr-setup-card-body>` +
				`<yr-setup-card-footer>` +
				`<span></span>` +
				`<yr-button variant="primary" type="submit">` +
				`Continue</yr-button>` +
				`</yr-setup-card-footer>` +
				`</yr-setup-card>`,
		)},
		{Name: "selection-step", Component: templ.Raw(
			`<style>.setup-option{display:flex;align-items:center;` +
				`gap:var(--yr-space-3);padding:var(--yr-space-3)` +
				` var(--yr-space-4);border:1px solid var(--yr-border);` +
				`border-radius:var(--yr-radius-md);cursor:pointer;` +
				`font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text)}` +
				`.setup-option:hover{border-color:var(--yr-accent);` +
				`background:var(--yr-accent-subtle)}` +
				`.setup-option input{margin:0}` +
				`.setup-option-text{display:flex;flex-direction:column;` +
				`gap:var(--yr-space-1)}` +
				`.setup-option-label{font-weight:500}` +
				`.setup-option-desc{font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-muted)}</style>` +
				`<yr-setup-card>` +
				`<yr-setup-card-progress steps="4" current="2">` +
				`</yr-setup-card-progress>` +
				`<yr-setup-card-header>` +
				`<h2 class="yr-setup-card-title">` +
				`Choose deployment type</h2>` +
				`<p class="yr-setup-card-subtitle">` +
				`Select how you want to run your instance.</p>` +
				`</yr-setup-card-header>` +
				`<yr-setup-card-body>` +
				`<label class="setup-option">` +
				`<input type="radio" name="deploy" value="saas"` +
				` checked>` +
				`<div class="setup-option-text">` +
				`<span class="setup-option-label">Cloud (SaaS)</span>` +
				`<span class="setup-option-desc">` +
				`Managed hosting, automatic updates</span>` +
				`</div></label>` +
				`<label class="setup-option">` +
				`<input type="radio" name="deploy" value="self">` +
				`<div class="setup-option-text">` +
				`<span class="setup-option-label">Self-hosted</span>` +
				`<span class="setup-option-desc">` +
				`Run on your own infrastructure</span>` +
				`</div></label>` +
				`</yr-setup-card-body>` +
				`<yr-setup-card-footer>` +
				`<a class="yr-setup-card-back" href="#">Back</a>` +
				`<yr-button variant="primary">Continue</yr-button>` +
				`</yr-setup-card-footer>` +
				`</yr-setup-card>`,
		)},
		{Name: "confirmation", Component: templ.Raw(
			`<yr-setup-card>` +
				`<yr-setup-card-progress steps="4" current="4">` +
				`</yr-setup-card-progress>` +
				`<yr-setup-card-header>` +
				`<h2 class="yr-setup-card-title">` +
				`You&#39;re all set</h2>` +
				`<p class="yr-setup-card-subtitle">` +
				`Your workspace is ready. You can invite teammates` +
				` or start exploring on your own.</p>` +
				`</yr-setup-card-header>` +
				`<yr-setup-card-body>` +
				`<yr-alert tone="success">Organization created` +
				` successfully.</yr-alert>` +
				`</yr-setup-card-body>` +
				`<yr-setup-card-footer>` +
				`<span></span>` +
				`<yr-button variant="primary">` +
				`Go to dashboard</yr-button>` +
				`</yr-setup-card-footer>` +
				`</yr-setup-card>`,
		)},
	}
}
