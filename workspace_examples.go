package ui

import "github.com/a-h/templ"

const workspaceBase = `<style>html,body{margin:0;height:100%}` +
	`.demo-region{display:flex;align-items:center;justify-content:center;` +
	`font-family:var(--yr-font-mono);font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-subtle)}</style>`

func workspaceExamples() []Example {
	return []Example{
		{Name: "expanded", Component: templ.Raw(
			workspaceBase +
				`<yr-page-workspace>` +
				`<yr-nav-shell>` +
				`<div class="demo-region" style="width:100%">nav</div>` +
				`</yr-nav-shell>` +
				`<yr-sidebar>` +
				`<div class="demo-region" style="height:100%">sidebar</div>` +
				`</yr-sidebar>` +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content>` +
				`<yr-content-main>` +
				`<div class="demo-region" style="height:100%">main</div>` +
				`</yr-content-main>` +
				`</yr-content>` +
				`</yr-page-workspace>`,
		), DisplayURL: ""},
		{Name: "with-drawer", Component: templ.Raw(
			workspaceBase +
				`<yr-page-workspace>` +
				`<yr-nav-shell>` +
				`<div class="demo-region" style="width:100%">nav</div>` +
				`</yr-nav-shell>` +
				`<yr-sidebar>` +
				`<div class="demo-region" style="height:100%">sidebar</div>` +
				`</yr-sidebar>` +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content>` +
				`<yr-content-main>` +
				`<div class="demo-region" style="height:100%">main</div>` +
				`</yr-content-main>` +
				`<yr-resize-handle target="drawer"></yr-resize-handle>` +
				`<yr-drawer open>` +
				`<div class="demo-region" style="height:100%">drawer</div>` +
				`</yr-drawer>` +
				`</yr-content>` +
				`</yr-page-workspace>`,
		), DisplayURL: ""},
		{Name: exampleCollapsed, Component: templ.Raw(
			workspaceBase +
				`<yr-page-workspace>` +
				`<yr-nav-shell>` +
				`<div class="demo-region" style="width:100%">nav</div>` +
				`</yr-nav-shell>` +
				`<yr-sidebar collapsed>` +
				`<div class="demo-region" style="height:100%">sidebar</div>` +
				`</yr-sidebar>` +
				`<yr-resize-handle target="sidebar"></yr-resize-handle>` +
				`<yr-content>` +
				`<yr-content-main>` +
				`<div class="demo-region" style="height:100%">main</div>` +
				`</yr-content-main>` +
				`</yr-content>` +
				`</yr-page-workspace>`,
		), DisplayURL: ""},
	}
}
