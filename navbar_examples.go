package ui

import "github.com/a-h/templ"

func navbarExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<style>.demo-slot{background:var(--yr-bg);border:1px dashed var(--yr-border);border-radius:var(--yr-radius-sm);` +
				`padding:var(--yr-space-1) var(--yr-space-3);color:var(--yr-text-disabled);font-family:var(--yr-font-mono);font-size:var(--yr-font-size-xs)}</style>` +
				`<yr-nav-shell>` +
				`<div class="demo-slot" slot="left">left</div>` +
				`<div class="demo-slot" slot="center">center</div>` +
				`<div class="demo-slot">right</div>` +
				`</yr-nav-shell>`,
		)},
		{Name: "with-content", Component: templ.Raw(
			`<yr-nav-shell>` +
				`<a href="#" class="yr-brand" slot="left">Yardrail</a>` +
				`<div class="yr-search yr-search--compact" slot="center" style="max-width:320px">` +
				exampleIcon("search", `class="yr-search-icon" style="left:0.5rem"`, 14) +
				`<input class="yr-search-input" type="search" placeholder="Search…" style="padding:var(--yr-space-1) var(--yr-space-2) var(--yr-space-1) 1.75rem;font-size:var(--yr-font-size-sm)">` +
				`</div>` +
				`<a href="#" style="font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);color:var(--yr-text-muted);text-decoration:none">Docs</a>` +
				`<yr-account-menu name="Brahm Lower" email="brahm@yardrail.com">` +
				`<span slot="avatar">BL</span>` +
				`<a class="yr-dropdown-menu-item" href="#">Account settings</a>` +
				`<a class="yr-dropdown-menu-item" href="#">Sign out</a>` +
				`</yr-account-menu>` +
				`</yr-nav-shell>`,
		)},
	}
}
