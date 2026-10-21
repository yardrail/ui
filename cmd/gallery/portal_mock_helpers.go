package main

const portalHome = "home"

// portalLink is one sidebar link in the portal mocks.
type portalLink struct {
	icon, label, slug string
}

// portalSpaceLinks returns the Spaces section links shared by every portal sidebar.
func portalSpaceLinks() []portalLink {
	return []portalLink{
		{"star", "Marketing Ops", "space-marketing"},
		{"star", "Client Onboarding", "space-onboarding"},
		{"folder", "My Space", "space-personal"},
		{"folder", "Sales Ops", "space-sales"},
	}
}
