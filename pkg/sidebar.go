package ui

import "github.com/a-h/templ"

// SidebarProps configures a Sidebar.
type SidebarProps struct {
	Collapsed bool
	Attrs     templ.Attributes
}

// Sidebar renders the main navigation sidebar.
func Sidebar(p *SidebarProps) templ.Component {
	return sidebarComponent(p)
}

// SidebarHeaderProps configures a SidebarHeader.
type SidebarHeaderProps struct {
	Attrs templ.Attributes
}

// SidebarHeader renders the header region of a sidebar.
func SidebarHeader(p *SidebarHeaderProps) templ.Component {
	return sidebarHeader(p)
}

// SidebarNavProps configures a SidebarNav.
type SidebarNavProps struct {
	Label  string
	Bottom bool
	Attrs  templ.Attributes
}

// SidebarNav renders a navigation group within a sidebar.
func SidebarNav(p *SidebarNavProps) templ.Component {
	return sidebarNav(p)
}

// SidebarLinkProps configures a SidebarLink.
type SidebarLinkProps struct {
	Href   string
	Active bool
	Attrs  templ.Attributes
}

// SidebarLink renders a single navigation link within a sidebar nav group.
func SidebarLink(p *SidebarLinkProps) templ.Component {
	return sidebarLink(p)
}

// SidebarOrgProps configures a SidebarOrg.
type SidebarOrgProps struct {
	Name    string
	Initial string
	Attrs   templ.Attributes
}

// SidebarOrg renders an organization badge within a sidebar.
func SidebarOrg(p *SidebarOrgProps) templ.Component {
	return sidebarOrg(p)
}

// SidebarFooterProps configures a SidebarFooter.
type SidebarFooterProps struct {
	Attrs templ.Attributes
}

// SidebarFooter renders the footer region of a sidebar.
func SidebarFooter(p *SidebarFooterProps) templ.Component {
	return sidebarFooter(p)
}
