package ui

import (
	"fmt"
	"testing"
)

func TestRenderAuthorLink(t *testing.T) {
	t.Parallel()

	i := 0

	for _, large := range []bool{false, true} {
		for _, verified := range []bool{false, true} {
			for _, heading := range []bool{false, true} {
				html := renderString(t, AuthorLink("Author", &AuthorLinkProps{
					Large:    large,
					Verified: verified,
					Heading:  heading,
					Avatar:   "AU",
					Handle:   "author",
				}))
				writeClassCoverage(t, "author_link", fmt.Sprintf("%03d", i), html)
				i++
			}
		}
	}

	html := renderString(t, AuthorLink("Author", &AuthorLinkProps{
		Href:   "/profile",
		Avatar: "AU",
	}))
	writeClassCoverage(t, "author_link", "with-href", html)

	html = renderString(t, AuthorLink("", &AuthorLinkProps{Handle: "author"}))
	writeClassCoverage(t, "author_link", "handle-only", html)

	html = renderString(t, AuthorLink("Author", &AuthorLinkProps{}))
	writeClassCoverage(t, "author_link", "name-only", html)
}
