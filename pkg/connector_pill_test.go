package ui

import "testing"

func TestRenderConnectorPill(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		ConnectorPill(&ConnectorPillProps{Title: "Slack"}),
		text("SL"),
	))
	writeClassCoverage(t, "connector_pill", "default", html)
}
