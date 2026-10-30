package ui

import (
	"testing"
)

func TestRenderAlert(t *testing.T) {
	t.Parallel()

	for _, tone := range (Tone{}).All() {
		html := renderString(t, withChildren(
			Alert(&AlertProps{Tone: tone}),
			text("message"),
		))
		writeClassCoverage(t, "alert", tone.String(), html)
	}
}
