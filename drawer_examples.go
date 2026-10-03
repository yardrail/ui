package ui

import "github.com/a-h/templ"

const closeIconSize = 16

const drawerSlotStyle = `<style>.demo-slot{background:var(--yr-bg);` +
	`border:1px dashed var(--yr-border);` +
	`border-radius:var(--yr-radius-sm);` +
	`padding:var(--yr-space-1) var(--yr-space-3);` +
	`color:var(--yr-text-disabled);` +
	`font-family:var(--yr-font-mono);` +
	`font-size:var(--yr-font-size-xs)}</style>`

func drawerExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			drawerSlotStyle +
				`<yr-drawer open style="height:100vh">` +
				`<yr-drawer-header>` +
				`<div class="demo-slot">title</div>` +
				`<yr-drawer-close>` +
				exampleIcon("x", "", closeIconSize) +
				`</yr-drawer-close>` +
				`</yr-drawer-header>` +
				`<yr-drawer-body>` +
				`<div class="demo-slot">body</div>` +
				`</yr-drawer-body>` +
				`</yr-drawer>`,
		)},
		{Name: "with-content", Component: templ.Raw(
			`<yr-drawer open style="height:100vh">` +
				`<yr-drawer-header>` +
				`<h3 class="yr-drawer-title">Details</h3>` +
				`<yr-drawer-close>` +
				exampleIcon("x", "", closeIconSize) +
				`</yr-drawer-close>` +
				`</yr-drawer-header>` +
				`<yr-drawer-body>` +
				`<p style="font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted);` +
				`margin:0 0 var(--yr-space-4)">` +
				`Drawer content goes here.</p>` +
				`<div class="yr-card">` +
				`<div class="yr-card-title">Status</div>` +
				`<span class="yr-pill yr-pill--success">` +
				`Active</span>` +
				`</div>` +
				`</yr-drawer-body>` +
				`</yr-drawer>`,
		)},
		{Name: "with-form", Component: templ.Raw(
			`<yr-drawer open style="height:100vh">` +
				`<yr-drawer-header>` +
				`<h3 class="yr-drawer-title">Edit Session</h3>` +
				`<yr-drawer-close>` +
				exampleIcon("x", "", closeIconSize) +
				`</yr-drawer-close>` +
				`</yr-drawer-header>` +
				`<yr-drawer-body>` +
				`<div class="yr-field"` +
				` style="margin-bottom:var(--yr-space-4)">` +
				`<label class="yr-field-label">Name</label>` +
				`<input class="yr-field-input" type="text"` +
				` value="Weekly Review">` +
				`</div>` +
				`<div class="yr-field"` +
				` style="margin-bottom:var(--yr-space-4)">` +
				`<label class="yr-field-label">` +
				`Description</label>` +
				`<textarea class="yr-field-input" rows="3">` +
				`Automated weekly report generation</textarea>` +
				`</div>` +
				`<button class="yr-button yr-button--primary">` +
				`Save</button>` +
				`</yr-drawer-body>` +
				`</yr-drawer>`,
		)},
	}
}
