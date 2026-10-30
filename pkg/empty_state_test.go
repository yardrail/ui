package ui

import "testing"

func TestRenderEmptyState(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		EmptyState(&EmptyStateProps{}),
		text("No results"),
	))
	writeClassCoverage(t, "empty_state", "default", html)
}
