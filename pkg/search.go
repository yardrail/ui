package ui

import "github.com/a-h/templ"

// Search variants.
const (
	SearchHero    = "hero"
	SearchCompact = "compact"
)

// SearchProps configures a Search form.
type SearchProps struct {
	Variant     string
	Placeholder string
	Action      string
	Name        string
	Attrs       templ.Attributes
}

func (p *SearchProps) class() string {
	if p.Variant != "" {
		return "yr-search yr-search--" + p.Variant
	}

	return "yr-search"
}

func (p *SearchProps) action() string {
	if p.Action != "" {
		return p.Action
	}

	return "#"
}

func (p *SearchProps) name() string {
	if p.Name != "" {
		return p.Name
	}

	return "q"
}

// Search renders a search form with icon and input.
func Search(p *SearchProps) templ.Component {
	return search(p)
}
