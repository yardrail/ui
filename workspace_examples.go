package ui

import "github.com/a-h/templ"

const workspaceBase = `<style>html,body{margin:0;height:100%}` +
	`.demo-region{display:flex;align-items:center;justify-content:center;` +
	`font-family:var(--font-mono);font-size:0.8rem;color:var(--slate-light)}</style>`

func workspaceExamples() []Example {
	return []Example{
		{Name: "expanded", Component: templ.Raw(
			workspaceBase +
				`<div class="yr-sidebar-expanded" style="height:100vh">` +
				`<nav class="yr-nav-shell"><div class="demo-region" style="width:100%">nav</div></nav>` +
				`<aside class="yr-sidebar"><div class="demo-region" style="height:100%">sidebar</div></aside>` +
				`<div class="yr-content">` +
				`<main class="yr-content-main"><div class="demo-region" style="height:100%">main</div></main>` +
				`</div>` +
				`</div>`,
		)},
		{Name: "with-drawer", Component: templ.Raw(
			workspaceBase +
				`<div class="yr-sidebar-expanded" style="height:100vh">` +
				`<nav class="yr-nav-shell"><div class="demo-region" style="width:100%">nav</div></nav>` +
				`<aside class="yr-sidebar"><div class="demo-region" style="height:100%">sidebar</div></aside>` +
				`<div class="yr-content">` +
				`<main class="yr-content-main"><div class="demo-region" style="height:100%">main</div></main>` +
				`<aside class="yr-drawer yr-drawer--open"><div class="demo-region" style="height:100%">drawer</div></aside>` +
				`</div>` +
				`</div>`,
		)},
		{Name: "collapsed", Component: templ.Raw(
			workspaceBase +
				`<div class="yr-sidebar-collapsed" style="height:100vh">` +
				`<nav class="yr-nav-shell"><div class="demo-region" style="width:100%">nav</div></nav>` +
				`<aside class="yr-sidebar"><div class="demo-region" style="height:100%">sidebar</div></aside>` +
				`<div class="yr-content">` +
				`<main class="yr-content-main"><div class="demo-region" style="height:100%">main</div></main>` +
				`</div>` +
				`</div>`,
		)},
	}
}
