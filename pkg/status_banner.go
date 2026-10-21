package ui

import "github.com/a-h/templ"

// StatusBannerProps configures a StatusBanner.
type StatusBannerProps struct {
	Tone  string
	Attrs templ.Attributes
}

// StatusBanner renders a persistent status display bar. Children are the banner content.
func StatusBanner(p StatusBannerProps) templ.Component {
	return statusBanner(&p)
}
