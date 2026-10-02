package ui

import "github.com/a-h/templ"

func sidebarExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<aside class="yr-sidebar" style="position:relative;top:0;height:100vh">` +
				`<div class="yr-sidebar-header">` +
				`<span style="font-family:var(--yr-font-body);font-weight:600;font-size:var(--yr-font-size-sm)">Workspace</span>` +
				`</div>` +
				`<ul class="yr-sidebar-nav">` +
				`<li><a class="yr-sidebar-link yr-sidebar-link--active" href="#">` +
				exampleIcon("layout-dashboard", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Dashboard</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("list", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Sessions</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("file-text", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Templates</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("settings", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Settings</span></a></li>` +
				`</ul>` +
				`<div class="yr-sidebar-footer">` +
				`<span style="font-size:var(--yr-font-size-xs);color:var(--yr-text-subtle)">v0.1.0</span>` +
				`</div>` +
				`</aside>`,
		)},
		{Name: "collapsed", Component: templ.Raw(
			`<div class="yr-sidebar-collapsed">` +
				`<aside class="yr-sidebar" style="position:relative;top:0;height:100vh">` +
				`<div class="yr-sidebar-header">` +
				`<span style="font-family:var(--yr-font-body);font-weight:600;font-size:var(--yr-font-size-sm)">W</span>` +
				`</div>` +
				`<ul class="yr-sidebar-nav">` +
				`<li><a class="yr-sidebar-link yr-sidebar-link--active" href="#">` +
				exampleIcon("layout-dashboard", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Dashboard</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("list", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Sessions</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("file-text", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Templates</span></a></li>` +
				`<li><a class="yr-sidebar-link" href="#">` +
				exampleIcon("settings", "", searchIconSize) +
				`<span class="yr-sidebar-link-label">Settings</span></a></li>` +
				`</ul>` +
				`</aside>` +
				`</div>`,
		)},
	}
}
