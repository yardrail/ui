package ui

import "github.com/a-h/templ"

// AvatarProps configures an Avatar.
type AvatarProps struct {
	Size  string
	Attrs templ.Attributes
}

// Avatar renders a circular initials badge. The text argument is the initials.
func Avatar(text string, p AvatarProps) templ.Component {
	return avatar(text, &p)
}
