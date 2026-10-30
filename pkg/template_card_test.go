package ui

import "testing"

func TestRenderTemplateCard(t *testing.T) {
	t.Parallel()

	html := renderString(t, TemplateCard("Invoice Sync", "Automates invoicing", "Free", &TemplateCardProps{
		Href:       "/templates/invoice",
		Connectors: text("<span>Slack</span>"),
		FooterLeft: text("<span>100 installs</span>"),
	}))
	writeClassCoverage(t, "template_card", "default", html)
}
