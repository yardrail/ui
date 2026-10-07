package ui

import "github.com/a-h/templ"

// searchIconSize is the search icon size the apps use in the default search box.
const searchIconSize = 16

func navShellExamples() []Example {
	return []Example{
		{Name: "top-nav", Component: templ.Raw(
			`<yr-nav-shell>` +
				`<a href="#" class="yr-brand" slot="left">Yardrail</a>` +
				`<div class="yr-search" slot="center">` +
				exampleIcon("search", `class="yr-search-icon"`, searchIconSize) +
				`<input class="yr-search-input" type="search" placeholder="Search workflows, connectors…">` +
				`</div>` +
				`<a href="#" style="font-family:var(--yr-font-ui);` +
				`font-size:var(--yr-font-size-sm);` +
				`color:var(--yr-text-muted);` +
				`text-decoration:none">Docs</a>` +
				`<yr-avatar>BL</yr-avatar>` +
				`</yr-nav-shell>`,
		), DisplayURL: ""},
	}
}
