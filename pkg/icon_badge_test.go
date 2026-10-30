package ui

import "testing"

func TestRenderIconBadge(t *testing.T) {
	t.Parallel()

	for _, tone := range (Tone{}).All() {
		html := renderString(t, withChildren(
			IconBadge(tone, &IconBadgeProps{}),
			text("icon"),
		))
		writeClassCoverage(t, "icon_badge", tone.String(), html)
	}
}
