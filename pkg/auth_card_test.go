package ui

import "testing"

func TestRenderAuthCard(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		AuthCard(&AuthCardProps{}),
		text("content"),
	))
	writeClassCoverage(t, "auth_card", "default", html)
}
