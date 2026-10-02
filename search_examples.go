package ui

import "github.com/a-h/templ"

// The search icon sizes the apps use in the hero and compact search boxes.
const (
	searchHeroIconSize    = 18
	searchCompactIconSize = 15
)

func searchExample(modifier, placeholder string, iconSize int) templ.Component {
	class := "yr-search"
	if modifier != "" {
		class += " yr-search--" + modifier
	}

	return templ.Raw(
		`<form class="` + class + `" method="get" action="#" role="search">` +
			exampleIcon("search", `class="yr-search-icon"`, iconSize) +
			`<input class="yr-search-input" type="search" name="q" placeholder="` + placeholder + `" autocomplete="off">` +
			`</form>`,
	)
}

func searchExamples() []Example {
	return []Example{
		{Name: "plain", Component: searchExample("", "Search templates…", searchIconSize)},
		{Name: "hero", Component: searchExample("hero", "Search 500+ templates…", searchHeroIconSize)},
		{Name: "compact", Component: searchExample("compact", "Search connectors…", searchCompactIconSize)},
	}
}
