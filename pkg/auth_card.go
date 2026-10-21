package ui

import "github.com/a-h/templ"

// AuthCardProps configures an AuthCard.
type AuthCardProps struct {
	Attrs templ.Attributes
}

// AuthCard renders a centered auth card (login, signup, invite). Children are the card content.
func AuthCard(p *AuthCardProps) templ.Component {
	return authCard(p)
}
