package ui

import "github.com/a-h/templ"

// TemplateCardProps configures a TemplateCard.
type TemplateCardProps struct {
	Href       string
	Connectors templ.Component
	FooterLeft templ.Component
	Attrs      templ.Attributes
}

func (p *TemplateCardProps) href() string {
	if p.Href != "" {
		return p.Href
	}

	return "#"
}

// TemplateCard renders a template catalog card with name, description, price, and optional connectors.
func TemplateCard(name, desc, price string, p *TemplateCardProps) templ.Component {
	return templateCard(name, desc, price, p)
}
