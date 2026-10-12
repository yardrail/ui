package ui

import "strings"

const (
	portalIconSize = 16
	portalHome     = "home"
)

const portalMockStyle = `<style>` +
	`html,body{margin:0;height:100%}` +
	`.portal-page-title{font-family:var(--yr-font-display);` +
	`font-size:var(--yr-font-size-xl);font-weight:600;` +
	`color:var(--yr-text);margin:0 0 var(--yr-space-6)}` +
	`.portal-section{margin-bottom:var(--yr-space-8)}` +
	`.portal-section-title{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-sm);font-weight:600;` +
	`text-transform:uppercase;letter-spacing:0.04em;` +
	`color:var(--yr-text-subtle);margin:0 0 var(--yr-space-3)}` +
	`.portal-grid{display:grid;gap:var(--yr-space-4)}` +
	`.portal-grid-2{grid-template-columns:1fr 1fr}` +
	`.portal-grid-3{grid-template-columns:1fr 1fr 1fr}` +
	`.portal-grid-4{grid-template-columns:1fr 1fr 1fr 1fr}` +
	`.portal-stat-card{background:var(--yr-surface);` +
	`border:1px solid var(--yr-border);` +
	`border-radius:var(--yr-radius-md);` +
	`padding:var(--yr-space-4) var(--yr-space-6)}` +
	`.portal-stat-value{font-family:var(--yr-font-display);` +
	`font-size:var(--yr-font-size-2xl);font-weight:600;` +
	`color:var(--yr-text);margin:0}` +
	`.portal-stat-label{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-muted);margin:var(--yr-space-1) 0 0}` +
	`.portal-feed{display:flex;flex-direction:column;` +
	`gap:var(--yr-space-2)}` +
	`.portal-feed-item{display:flex;align-items:center;` +
	`gap:var(--yr-space-3);padding:var(--yr-space-2) var(--yr-space-3);` +
	`font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);` +
	`color:var(--yr-text-muted);border-radius:var(--yr-radius-sm)}` +
	`.portal-feed-item:hover{background:var(--yr-surface-muted)}` +
	`.portal-feed-time{font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-disabled);margin-left:auto;white-space:nowrap}` +
	`.portal-placeholder{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-sm);color:var(--yr-text-disabled);` +
	`padding:var(--yr-space-8);text-align:center;` +
	`border:2px dashed var(--yr-border);border-radius:var(--yr-radius-md)}` +
	`</style>`

func portalNavbar() string {
	return `<yr-navbar-std logo="/logo.png" brand="Yardrail" authed>` +
		`<div class="yr-search yr-search--compact"` +
		` slot="center" style="max-width:320px">` +
		exampleIcon("search",
			`class="yr-search-icon" style="left:0.5rem"`,
			navSearchIconSize) +
		`<input class="yr-search-input" type="search"` +
		` placeholder="Search…"` +
		` style="padding:var(--yr-space-1) var(--yr-space-2)` +
		` var(--yr-space-1) 1.75rem;` +
		`font-size:var(--yr-font-size-sm)">` +
		`</div>` +
		`<yr-account-menu name="Brahm Lower"` +
		` email="brahm@yardrail.com">` +
		`<span slot="avatar">BL</span>` +
		`<a class="yr-dropdown-menu-item" href="#">Account settings</a>` +
		`<a class="yr-dropdown-menu-item" href="#">Sign out</a>` +
		`</yr-account-menu>` +
		`</yr-navbar-std>`
}

func portalSidebarFooter() string {
	return `<yr-sidebar-nav bottom>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("settings", "", portalIconSize) +
		` Settings</yr-sidebar-link>` +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-footer>` +
		`<span style="font-size:var(--yr-font-size-xs);` +
		`color:var(--yr-text-subtle)">v0.1.0</span>` +
		`</yr-sidebar-footer>`
}

// portalLink is one sidebar link in the portal mocks.
type portalLink struct {
	icon, label, slug string
}

// portalSpaceLinks returns the Spaces section links shared by every portal sidebar.
func portalSpaceLinks() []portalLink {
	return []portalLink{
		{"star", "Marketing Ops", "space-marketing"},
		{"star", "Client Onboarding", "space-onboarding"},
		{"folder", "My Space", "space-personal"},
		{"folder", "Sales Ops", "space-sales"},
	}
}

func portalSidebarHeader() string {
	return `<yr-sidebar>` +
		`<yr-sidebar-header>` +
		`<span style="font-family:var(--yr-font-body);` +
		`font-weight:600;font-size:var(--yr-font-size-sm)">` +
		`Yardrail</span>` +
		`</yr-sidebar-header>`
}

func portalSidebarLink(l portalLink, activePage string) string {
	active := ""
	if l.slug == activePage {
		active = " active"
	}

	return `<yr-sidebar-link href="#"` + active + `>` +
		exampleIcon(l.icon, "", portalIconSize) +
		` ` + l.label +
		`</yr-sidebar-link>`
}

func portalSidebarLinks(links []portalLink, activePage string) string {
	parts := make([]string, 0, len(links))
	for _, l := range links {
		parts = append(parts, portalSidebarLink(l, activePage))
	}

	return strings.Join(parts, "")
}

func portalSidebar(activePage string) string {
	topLinks := []portalLink{
		{portalHome, "Home", portalHome},
		{"zap", "Automations", "automations"},
	}

	return portalSidebarHeader() +
		`<yr-sidebar-nav>` +
		portalSidebarLinks(topLinks, activePage) +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-nav label="Spaces">` +
		portalSidebarLinks(portalSpaceLinks(), activePage) +
		`</yr-sidebar-nav>` +
		portalSidebarFooter() +
		`</yr-sidebar>`
}

func portalSidebarAutomations(activePage string) string {
	topLinks := []portalLink{
		{portalHome, "Home", ""},
		{"zap", "Automations", "automations"},
	}

	autoSubLinks := []portalLink{
		{"git-branch", "Workflows", "workflows"},
		{"layout-grid", "Templates", "templates"},
		{"plug", "Connectors", "connectors"},
	}

	return portalSidebarHeader() +
		`<yr-sidebar-nav>` +
		portalSidebarLinks(topLinks, "automations") +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-nav>` +
		portalSidebarLinks(autoSubLinks, activePage) +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-nav label="Spaces">` +
		portalSidebarLinks(portalSpaceLinks(), "") +
		`</yr-sidebar-nav>` +
		portalSidebarFooter() +
		`</yr-sidebar>`
}
