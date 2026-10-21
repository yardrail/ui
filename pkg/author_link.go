package ui

import "github.com/a-h/templ"

// AuthorLinkProps configures an AuthorLink.
type AuthorLinkProps struct {
	Href     string
	Handle   string
	Verified bool
	Avatar   string
	Size     string
	Heading  bool
	Attrs    templ.Attributes
}

// AuthorLink renders an author identity: avatar, name, handle, and verified badge.
func AuthorLink(name string, p *AuthorLinkProps) templ.Component {
	return authorLink(name, p)
}
