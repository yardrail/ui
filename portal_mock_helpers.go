package ui

import "strings"

const portalIconSize = 16

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
	return mspNavbar()
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

func portalSidebar(activePage string) string {
	type link struct {
		icon, label, slug string
	}

	topLinks := []link{
		{"home", "Home", "home"},
		{"zap", "Automations", "automations"},
	}

	spaceLinks := []link{
		{"star", "Marketing Ops", "space-marketing"},
		{"star", "Client Onboarding", "space-onboarding"},
		{"folder", "My Space", "space-personal"},
		{"folder", "Sales Ops", "space-sales"},
	}

	var b strings.Builder

	b.WriteString(`<yr-sidebar>`)
	b.WriteString(`<yr-sidebar-header>`)
	b.WriteString(`<span style="font-family:var(--yr-font-body);`)
	b.WriteString(`font-weight:600;font-size:var(--yr-font-size-sm)">`)
	b.WriteString(`Yardrail</span>`)
	b.WriteString(`</yr-sidebar-header>`)

	b.WriteString(`<yr-sidebar-nav>`)
	for _, l := range topLinks {
		active := ""
		if l.slug == activePage {
			active = " active"
		}
		b.WriteString(`<yr-sidebar-link href="#"`)
		b.WriteString(active)
		b.WriteString(`>`)
		b.WriteString(exampleIcon(l.icon, "", portalIconSize))
		b.WriteString(` `)
		b.WriteString(l.label)
		b.WriteString(`</yr-sidebar-link>`)
	}
	b.WriteString(`</yr-sidebar-nav>`)

	b.WriteString(`<yr-sidebar-nav label="Spaces">`)
	for _, l := range spaceLinks {
		active := ""
		if l.slug == activePage {
			active = " active"
		}
		b.WriteString(`<yr-sidebar-link href="#"`)
		b.WriteString(active)
		b.WriteString(`>`)
		b.WriteString(exampleIcon(l.icon, "", portalIconSize))
		b.WriteString(` `)
		b.WriteString(l.label)
		b.WriteString(`</yr-sidebar-link>`)
	}
	b.WriteString(`</yr-sidebar-nav>`)

	b.WriteString(portalSidebarFooter())
	b.WriteString(`</yr-sidebar>`)

	return b.String()
}

func portalSidebarAutomations(activePage string) string {
	type link struct {
		icon, label, slug string
	}

	autoSubLinks := []link{
		{"git-branch", "Workflows", "workflows"},
		{"layout-grid", "Templates", "templates"},
		{"plug", "Connectors", "connectors"},
	}

	spaceLinks := []link{
		{"star", "Marketing Ops", "space-marketing"},
		{"star", "Client Onboarding", "space-onboarding"},
		{"folder", "My Space", "space-personal"},
		{"folder", "Sales Ops", "space-sales"},
	}

	var b strings.Builder

	b.WriteString(`<yr-sidebar>`)
	b.WriteString(`<yr-sidebar-header>`)
	b.WriteString(`<span style="font-family:var(--yr-font-body);`)
	b.WriteString(`font-weight:600;font-size:var(--yr-font-size-sm)">`)
	b.WriteString(`Yardrail</span>`)
	b.WriteString(`</yr-sidebar-header>`)

	b.WriteString(`<yr-sidebar-nav>`)
	b.WriteString(`<yr-sidebar-link href="#">`)
	b.WriteString(exampleIcon("home", "", portalIconSize))
	b.WriteString(` Home</yr-sidebar-link>`)

	b.WriteString(`<yr-sidebar-link href="#" active>`)
	b.WriteString(exampleIcon("zap", "", portalIconSize))
	b.WriteString(` Automations</yr-sidebar-link>`)
	b.WriteString(`</yr-sidebar-nav>`)

	b.WriteString(`<yr-sidebar-nav>`)
	for _, l := range autoSubLinks {
		active := ""
		if l.slug == activePage {
			active = " active"
		}
		b.WriteString(`<yr-sidebar-link href="#"`)
		b.WriteString(active)
		b.WriteString(`>`)
		b.WriteString(exampleIcon(l.icon, "", portalIconSize))
		b.WriteString(` `)
		b.WriteString(l.label)
		b.WriteString(`</yr-sidebar-link>`)
	}
	b.WriteString(`</yr-sidebar-nav>`)

	b.WriteString(`<yr-sidebar-nav label="Spaces">`)
	for _, l := range spaceLinks {
		b.WriteString(`<yr-sidebar-link href="#">`)
		b.WriteString(exampleIcon(l.icon, "", portalIconSize))
		b.WriteString(` `)
		b.WriteString(l.label)
		b.WriteString(`</yr-sidebar-link>`)
	}
	b.WriteString(`</yr-sidebar-nav>`)

	b.WriteString(portalSidebarFooter())
	b.WriteString(`</yr-sidebar>`)

	return b.String()
}
