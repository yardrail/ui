package ui

import "testing"

func TestRenderStat(t *testing.T) {
	t.Parallel()

	html := renderString(t, Stat("42", &StatProps{Label: "Runs"}))
	writeClassCoverage(t, "stat", "default", html)
}
