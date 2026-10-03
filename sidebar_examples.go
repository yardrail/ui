package ui

import "github.com/a-h/templ"

func sidebarExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<yr-sidebar style="position:relative;top:0;height:100vh">` +
				`<yr-sidebar-header>` +
				`<span style="font-family:var(--yr-font-body);font-weight:600;font-size:var(--yr-font-size-sm)">Workspace</span>` +
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
				`<span style="font-size:var(--yr-font-size-xs);color:var(--yr-text-subtle)">v0.1.0</span>` +
				`</yr-sidebar-footer>` +
				`</yr-sidebar>`,
		)},
		{Name: "collapsed", Component: templ.Raw(
			`<yr-sidebar collapsed style="position:relative;top:0;height:100vh">` +
				`<yr-sidebar-header>` +
				`<span style="font-family:var(--yr-font-body);font-weight:600;font-size:var(--yr-font-size-sm)">W</span>` +
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
