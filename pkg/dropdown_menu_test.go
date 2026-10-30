package ui

import "testing"

func TestRenderDropdownMenu(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		DropdownMenu(&DropdownMenuProps{}),
		text("menu items"),
	))
	writeClassCoverage(t, "dropdown_menu", "default", html)
}
