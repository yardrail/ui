package ui

import "github.com/a-h/templ"

func drawerExamples() []Example {
	return []Example{
		{Name: "open", Component: templ.Raw(
			`<aside class="yr-drawer yr-drawer--open" style="height:100vh">` +
				`<div class="yr-drawer-header">` +
				`<h3 class="yr-drawer-title">Details</h3>` +
				`<button style="background:none;border:none;padding:var(--yr-space-1);cursor:pointer;color:var(--yr-text-muted);display:flex;border-radius:var(--yr-radius-sm)" aria-label="Close drawer">` +
				exampleIcon("x", "", 16) +
				`</button>` +
				`</div>` +
				`<div class="yr-drawer-body">` +
				`<p style="font-size:var(--yr-font-size-sm);color:var(--yr-text-muted);margin:0 0 var(--yr-space-4)">Drawer content goes here. This panel pushes the main content narrower when open.</p>` +
				`<div class="yr-card">` +
				`<div class="yr-card-title">Status</div>` +
				`<span class="yr-pill yr-pill--success">Active</span>` +
				`</div>` +
				`</div>` +
				`</aside>`,
		)},
		{Name: "with-form", Component: templ.Raw(
			`<aside class="yr-drawer yr-drawer--open" style="height:100vh">` +
				`<div class="yr-drawer-header">` +
				`<h3 class="yr-drawer-title">Edit Session</h3>` +
				`<button style="background:none;border:none;padding:var(--yr-space-1);cursor:pointer;color:var(--yr-text-muted);display:flex;border-radius:var(--yr-radius-sm)" aria-label="Close drawer">` +
				exampleIcon("x", "", 16) +
				`</button>` +
				`</div>` +
				`<div class="yr-drawer-body">` +
				`<div class="yr-field" style="margin-bottom:var(--yr-space-4)">` +
				`<label class="yr-field-label">Name</label>` +
				`<input class="yr-field-input" type="text" value="Weekly Review">` +
				`</div>` +
				`<div class="yr-field" style="margin-bottom:var(--yr-space-4)">` +
				`<label class="yr-field-label">Description</label>` +
				`<textarea class="yr-field-input" rows="3">Automated weekly report generation</textarea>` +
				`</div>` +
				`<button class="yr-button yr-button--primary">Save</button>` +
				`</div>` +
				`</aside>`,
		)},
	}
}
