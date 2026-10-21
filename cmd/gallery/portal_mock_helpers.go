package main

import (
	"github.com/a-h/templ"

	ui "github.com/yardrail/ui/pkg"
)

const portalHome = "home"

// portalLink is one sidebar link in the portal mocks.
type portalLink struct {
	icon  templ.Component
	label string
	slug  string
}

// portalSpaceLinks returns the Spaces section links shared by every portal sidebar.
func portalSpaceLinks() []portalLink {
	return []portalLink{
		{ui.IconStar(), "Marketing Ops", "space-marketing"},
		{ui.IconStar(), "Client Onboarding", "space-onboarding"},
		{ui.IconFolder(), "My Space", "space-personal"},
		{ui.IconFolder(), "Sales Ops", "space-sales"},
	}
}
