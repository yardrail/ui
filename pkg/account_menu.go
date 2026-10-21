package ui

import "github.com/a-h/templ"

// AccountMenuProps configures an AccountMenu.
type AccountMenuProps struct {
	Name  string
	Email string
	Attrs templ.Attributes
}

// AccountMenu renders an account dropdown menu with avatar, name, and email.
func AccountMenu(p *AccountMenuProps) templ.Component {
	return accountMenu(p)
}
