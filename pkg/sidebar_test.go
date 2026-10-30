package ui

import (
	"fmt"
	"testing"
)

func TestRenderSidebar(t *testing.T) {
	t.Parallel()

	for i, collapsed := range []bool{false, true} {
		html := renderString(t, withChildren(
			Sidebar(&SidebarProps{Collapsed: collapsed}),
			text("nav"),
		))
		writeClassCoverage(t, "sidebar", fmt.Sprintf("%03d", i), html)
	}
}

func TestRenderSidebarLink(t *testing.T) {
	t.Parallel()

	for i, active := range []bool{false, true} {
		html := renderString(t, withChildren(
			SidebarLink(&SidebarLinkProps{Href: "#", Active: active}),
			text("Link"),
		))
		writeClassCoverage(t, "sidebar_link", fmt.Sprintf("%03d", i), html)
	}
}

func TestRenderSidebarHeader(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		SidebarHeader(&SidebarHeaderProps{}),
		text("header"),
	))
	writeClassCoverage(t, "sidebar_header", "default", html)
}

func TestRenderSidebarNav(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		SidebarNav(&SidebarNavProps{Label: "Navigation"}),
		text("nav items"),
	))
	writeClassCoverage(t, "sidebar_nav", "default", html)

	html = renderString(t, withChildren(
		SidebarNav(&SidebarNavProps{Label: "Secondary", Bottom: true}),
		text("bottom items"),
	))
	writeClassCoverage(t, "sidebar_nav", "bottom", html)
}

func TestRenderSidebarOrg(t *testing.T) {
	t.Parallel()

	html := renderString(t, SidebarOrg(&SidebarOrgProps{Name: "Acme Corp"}))
	writeClassCoverage(t, "sidebar_org", "name_only", html)

	html = renderString(t, SidebarOrg(&SidebarOrgProps{Name: "Acme Corp", Initial: "A"}))
	writeClassCoverage(t, "sidebar_org", "with_initial", html)

	html = renderString(t, SidebarOrg(&SidebarOrgProps{}))
	writeClassCoverage(t, "sidebar_org", "empty", html)
}

func TestRenderSidebarFooter(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		SidebarFooter(&SidebarFooterProps{}),
		text("footer"),
	))
	writeClassCoverage(t, "sidebar_footer", "default", html)
}
