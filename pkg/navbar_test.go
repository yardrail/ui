package ui

import (
	"fmt"
	"testing"
)

func TestRenderNavbar(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Navbar(&NavbarProps{Left: text("logo"), Center: text("search")}),
		text("right"),
	))
	writeClassCoverage(t, "navbar", "with_slots", html)

	html = renderString(t, withChildren(
		Navbar(&NavbarProps{}),
		text("right"),
	))
	writeClassCoverage(t, "navbar", "no_slots", html)
}

func TestRenderNavbarStd(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		props NavbarStdProps
	}{
		{"authed", NavbarStdProps{Authed: true}},
		{"unauthed", NavbarStdProps{}},
		{"unauthed_custom_hrefs", NavbarStdProps{LoginHref: "/auth/login", SignupHref: "/auth/signup"}},
		{"with_center", NavbarStdProps{Authed: true, Center: text("search")}},
		{"custom_brand", NavbarStdProps{Authed: true, Logo: "/custom.svg", Brand: "Acme"}},
	}

	for i, tc := range cases {
		html := renderString(t, withChildren(
			NavbarStd(&tc.props),
			text("children"),
		))
		writeClassCoverage(t, "navbar_std", fmt.Sprintf("%03d_%s", i, tc.name), html)
	}
}
