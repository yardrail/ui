package ui

import "testing"

func TestRenderLiveDuration(t *testing.T) {
	t.Parallel()

	html := renderString(t, LiveDuration("2026-01-01T00:00:00Z", &LiveDurationProps{}))
	writeClassCoverage(t, "live_duration", "default", html)
}
