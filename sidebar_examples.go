package ui

import "github.com/a-h/templ"

const sidebarSlotStyle = `<style>.demo-slot{background:var(--yr-bg);` +
	`border:1px dashed var(--yr-border);` +
	`border-radius:var(--yr-radius-sm);` +
	`padding:var(--yr-space-1) var(--yr-space-3);` +
	`color:var(--yr-text-disabled);` +
	`font-family:var(--yr-font-mono);` +
	`font-size:var(--yr-font-size-xs)}</style>`

func sidebarExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			sidebarSlotStyle +
				`<yr-sidebar style="position:relative;top:0;` +
				`height:100vh">` +
				`<yr-sidebar-header>` +
				`<div class="demo-slot">header</div>` +
				`</yr-sidebar-header>` +
				`<yr-sidebar-nav>` +
				`<div class="demo-slot">nav</div>` +
				`</yr-sidebar-nav>` +
				`<yr-sidebar-footer>` +
				`<div class="demo-slot">footer</div>` +
				`</yr-sidebar-footer>` +
				`</yr-sidebar>`,
		)},
		{Name: "with-content", Component: templ.Raw(
			`<yr-sidebar style="position:relative;top:0;` +
				`height:100vh">` +
				`<yr-sidebar-header>` +
				`<span style="font-family:var(--yr-font-body);` +
				`font-weight:600;` +
				`font-size:var(--yr-font-size-sm)">` +
				`Workspace</span>` +
				`</yr-sidebar-header>` +
				`<yr-sidebar-nav>` +
				`<yr-sidebar-link href="#" active>` +
				exampleIcon("layout-dashboard", "", searchIconSize) +
				` Dashboard</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("list", "", searchIconSize) +
				` Sessions</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("file-text", "", searchIconSize) +
				` Templates</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("settings", "", searchIconSize) +
				` Settings</yr-sidebar-link>` +
				`</yr-sidebar-nav>` +
				`<yr-sidebar-footer>` +
				`<span style="font-size:var(--yr-font-size-xs);` +
				`color:var(--yr-text-subtle)">v0.1.0</span>` +
				`</yr-sidebar-footer>` +
				`</yr-sidebar>`,
		)},
		{Name: "collapsed", Component: templ.Raw(
			`<yr-sidebar collapsed style="position:relative;` +
				`top:0;height:100vh">` +
				`<yr-sidebar-header>` +
				`<span style="font-family:var(--yr-font-body);` +
				`font-weight:600;` +
				`font-size:var(--yr-font-size-sm)">W</span>` +
				`</yr-sidebar-header>` +
				`<yr-sidebar-nav>` +
				`<yr-sidebar-link href="#" active>` +
				exampleIcon("layout-dashboard", "", searchIconSize) +
				` Dashboard</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("list", "", searchIconSize) +
				` Sessions</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("file-text", "", searchIconSize) +
				` Templates</yr-sidebar-link>` +
				`<yr-sidebar-link href="#">` +
				exampleIcon("settings", "", searchIconSize) +
				` Settings</yr-sidebar-link>` +
				`</yr-sidebar-nav>` +
				`</yr-sidebar>`,
		)},
	}
}
