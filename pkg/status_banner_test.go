package ui

import "testing"

func TestRenderStatusBanner(t *testing.T) {
	t.Parallel()

	for _, tone := range (Tone{}).All() {
		html := renderString(t, withChildren(
			StatusBanner(&StatusBannerProps{Tone: tone}),
			text("status"),
		))
		writeClassCoverage(t, "status_banner", tone.String(), html)
	}
}
