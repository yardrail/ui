package ui

import "github.com/a-h/templ"

// DropdownMenuProps configures a DropdownMenu.
type DropdownMenuProps struct {
	Attrs templ.Attributes
}

// DropdownMenu renders a dropdown menu with trigger and panel slots.
func DropdownMenu(p *DropdownMenuProps) templ.Component {
	return dropdownMenu(p)
}
