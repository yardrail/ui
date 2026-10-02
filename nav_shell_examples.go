package ui

import "github.com/a-h/templ"

// searchIconSize is the search icon size the apps use in the default search box.
const searchIconSize = 16

func navShellExamples() []Example {
	return []Example{
		{Name: "top-nav", Component: templ.Raw(
			`<nav class="yr-nav-shell">` +
				`<a href="#" class="yr-brand">Yardrail</a>` +
				`<div class="yr-search">` +
				exampleIcon("search", `class="yr-search-icon"`, searchIconSize) +
				`<input class="yr-search-input" type="search" placeholder="Search workflows, connectors…">` +
				`</div>` +
				`<div style="display:flex;align-items:center;gap:0.75rem">` +
				`<a href="#" class="yr-button yr-button--secondary">Docs</a>` +
				`<yr-avatar>BL</yr-avatar>` +
				`</div>` +
				`</nav>`,
		)},
	}
}
