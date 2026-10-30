package ui

import "testing"

func TestRenderAccountMenu(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		AccountMenu(&AccountMenuProps{Name: "Test User", Email: "test@test.com", Avatar: "TU"}),
		text("<a>Settings</a>"),
	))
	writeClassCoverage(t, "account_menu", "default", html)
}
