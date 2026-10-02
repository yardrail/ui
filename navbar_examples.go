package ui

import "github.com/a-h/templ"

func navbarExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<nav class="yr-nav-shell">` +
				`<div style="display:flex;align-items:center;gap:var(--yr-space-3)">` +
				`<button style="background:none;border:none;padding:var(--yr-space-1);cursor:pointer;color:var(--yr-text-muted);display:flex;border-radius:var(--yr-radius-sm)" aria-label="Toggle sidebar">` +
				exampleIcon("panel-left", "", exampleIconSize) +
				`</button>` +
				`<a href="#" class="yr-brand">Yardrail</a>` +
				`</div>` +
				`<div class="yr-search yr-search--compact" style="max-width:320px">` +
				exampleIcon("search", `class="yr-search-icon" style="left:0.5rem"`, 14) +
				`<input class="yr-search-input" type="search" placeholder="Search…" style="padding:var(--yr-space-1) var(--yr-space-2) var(--yr-space-1) 1.75rem;font-size:var(--yr-font-size-sm)">` +
				`</div>` +
				`<div style="display:flex;align-items:center;gap:var(--yr-space-4)">` +
				`<a href="#" style="font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);color:var(--yr-text-muted);text-decoration:none">Docs</a>` +
				`<yr-account-menu name="Brahm Lower" email="brahm@yardrail.com">` +
				`<span slot="avatar">BL</span>` +
				`<a class="yr-dropdown-menu-item" href="#">Account settings</a>` +
				`<a class="yr-dropdown-menu-item" href="#">Sign out</a>` +
				`</yr-account-menu>` +
				`</div>` +
				`</nav>`,
		)},
		{Name: "with-sub-brand", Component: templ.Raw(
			`<nav class="yr-nav-shell">` +
				`<div style="display:flex;align-items:center;gap:var(--yr-space-3)">` +
				`<button style="background:none;border:none;padding:var(--yr-space-1);cursor:pointer;color:var(--yr-text-muted);display:flex;border-radius:var(--yr-radius-sm)" aria-label="Toggle sidebar">` +
				exampleIcon("panel-left", "", exampleIconSize) +
				`</button>` +
				`<a href="#" class="yr-brand">Yardrail <span class="yr-brand-sub">Depot</span></a>` +
				`</div>` +
				`<div class="yr-search yr-search--compact" style="max-width:320px">` +
				exampleIcon("search", `class="yr-search-icon" style="left:0.5rem"`, 14) +
				`<input class="yr-search-input" type="search" placeholder="Search connectors…" style="padding:var(--yr-space-1) var(--yr-space-2) var(--yr-space-1) 1.75rem;font-size:var(--yr-font-size-sm)">` +
				`</div>` +
				`<div style="display:flex;align-items:center;gap:var(--yr-space-4)">` +
				`<a href="#" style="font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);color:var(--yr-text-muted);text-decoration:none">Docs</a>` +
				`<yr-account-menu name="Brahm Lower" email="brahm@yardrail.com">` +
				`<span slot="avatar">BL</span>` +
				`<a class="yr-dropdown-menu-item" href="#">Account settings</a>` +
				`<a class="yr-dropdown-menu-item" href="#">Sign out</a>` +
				`</yr-account-menu>` +
				`</div>` +
				`</nav>`,
		)},
	}
}
