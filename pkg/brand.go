package ui

import "github.com/a-h/templ"

// BrandProps configures a Brand link.
type BrandProps struct {
	Href  string
	Attrs templ.Attributes
}

func (p *BrandProps) href() string {
	if p.Href != "" {
		return p.Href
	}

	return "#"
}

// Brand renders a branded logo link. Children are the logo content.
func Brand(p *BrandProps) templ.Component {
	return brand(p)
}
