package ui

import "testing"

func TestRenderHitlPrompt(t *testing.T) {
	t.Parallel()

	for _, tone := range (Tone{}).All() {
		html := renderString(t, withChildren(
			HitlPrompt(&HitlPromptProps{Tone: tone}),
			text("prompt"),
		))
		writeClassCoverage(t, "hitl_prompt", tone.String(), html)
	}
}
