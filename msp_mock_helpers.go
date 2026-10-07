package ui

import "strings"

const mspIconSize = 16

const mspMockStyle = `<style>` +
	`html,body{margin:0;height:100%}` +
	`.msp-page-title{font-family:var(--yr-font-display);` +
	`font-size:var(--yr-font-size-xl);font-weight:600;` +
	`color:var(--yr-text);margin:0 0 var(--yr-space-6)}` +
	`.msp-section{margin-bottom:var(--yr-space-8)}` +
	`.msp-section-title{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-sm);font-weight:600;` +
	`text-transform:uppercase;letter-spacing:0.04em;` +
	`color:var(--yr-text-subtle);margin:0 0 var(--yr-space-3)}` +
	`.msp-grid{display:grid;gap:var(--yr-space-4)}` +
	`.msp-grid-2{grid-template-columns:1fr 1fr}` +
	`.msp-grid-3{grid-template-columns:1fr 1fr 1fr}` +
	`.msp-grid-4{grid-template-columns:1fr 1fr 1fr 1fr}` +
	`.msp-stat-card{background:var(--yr-surface);` +
	`border:1px solid var(--yr-border);` +
	`border-radius:var(--yr-radius-md);` +
	`padding:var(--yr-space-4) var(--yr-space-6)}` +
	`.msp-stat-value{font-family:var(--yr-font-display);` +
	`font-size:var(--yr-font-size-2xl);font-weight:600;` +
	`color:var(--yr-text);margin:0}` +
	`.msp-stat-label{font-family:var(--yr-font-ui);` +
	`font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-muted);margin:var(--yr-space-1) 0 0}` +
	`.msp-feed{display:flex;flex-direction:column;` +
	`gap:var(--yr-space-2)}` +
	`.msp-feed-item{display:flex;align-items:center;` +
	`gap:var(--yr-space-3);padding:var(--yr-space-2) var(--yr-space-3);` +
	`font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm);` +
	`color:var(--yr-text-muted);border-radius:var(--yr-radius-sm)}` +
	`.msp-feed-item:hover{background:var(--yr-surface-muted)}` +
	`.msp-feed-time{font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-disabled);margin-left:auto;white-space:nowrap}` +
	`.msp-client-row{display:flex;align-items:center;` +
	`gap:var(--yr-space-4);padding:var(--yr-space-3) var(--yr-space-4);` +
	`border:1px solid var(--yr-border);border-radius:var(--yr-radius-md);` +
	`background:var(--yr-surface);cursor:pointer;` +
	`font-family:var(--yr-font-ui);font-size:var(--yr-font-size-sm)}` +
	`.msp-client-row:hover{border-color:var(--yr-accent);` +
	`background:var(--yr-accent-subtle)}` +
	`.msp-client-name{font-weight:600;color:var(--yr-text);flex:1}` +
	`.msp-client-meta{color:var(--yr-text-muted);` +
	`font-size:var(--yr-font-size-xs)}` +
	`</style>`

func mspNavbar() string {
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

func mspSidebarFooter() string {
	return `<yr-sidebar-nav bottom>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("train-front", "", mspIconSize) +
		` Yardrail</yr-sidebar-link>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("user", "", mspIconSize) +
		` Account</yr-sidebar-link>` +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-footer>` +
		`<span style="font-size:var(--yr-font-size-xs);` +
		`color:var(--yr-text-subtle)">v0.1.0</span>` +
		`</yr-sidebar-footer>`
}

func mspSidebar() string {
	return `<yr-sidebar>` +
		`<yr-sidebar-header>` +
		`<span style="font-family:var(--yr-font-body);` +
		`font-weight:600;font-size:var(--yr-font-size-sm)">` +
		`MSP Portal</span>` +
		`</yr-sidebar-header>` +
		`<yr-sidebar-nav>` +
		`<yr-sidebar-link href="#" active>` +
		exampleIcon("layout-dashboard", "", mspIconSize) +
		` Dashboard</yr-sidebar-link>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("building-2", "", mspIconSize) +
		` Clients</yr-sidebar-link>` +
		`</yr-sidebar-nav>` +
		mspSidebarFooter() +
		`</yr-sidebar>`
}

func mspSidebarClients() string {
	return `<yr-sidebar>` +
		`<yr-sidebar-header>` +
		`<span style="font-family:var(--yr-font-body);` +
		`font-weight:600;font-size:var(--yr-font-size-sm)">` +
		`MSP Portal</span>` +
		`</yr-sidebar-header>` +
		`<yr-sidebar-nav>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("layout-dashboard", "", mspIconSize) +
		` Dashboard</yr-sidebar-link>` +
		`<yr-sidebar-link href="#" active>` +
		exampleIcon("building-2", "", mspIconSize) +
		` Clients</yr-sidebar-link>` +
		`</yr-sidebar-nav>` +
		mspSidebarFooter() +
		`</yr-sidebar>`
}

func mspSidebarClientDetail(activePage string) string {
	clientLinks := []struct {
		icon, label, slug string
	}{
		{"layout-dashboard", "Overview", "overview"},
		{"users", "Users", "users"},
		{"plug", "Connectors", "connectors"},
		{"play", "Workflows", "workflows"},
		{"settings", "Settings", "settings"},
	}

	links := make([]string, 0, len(clientLinks))
	for _, l := range clientLinks {
		active := ""
		if l.slug == activePage {
			active = " active"
		}

		links = append(links, `<yr-sidebar-link href="#"`+active+`>`+
			exampleIcon(l.icon, "", mspIconSize)+
			` `+l.label+`</yr-sidebar-link>`)
	}

	return `<yr-sidebar>` +
		`<yr-sidebar-header>` +
		`<span style="font-family:var(--yr-font-body);` +
		`font-weight:600;font-size:var(--yr-font-size-sm)">` +
		`MSP Portal</span>` +
		`</yr-sidebar-header>` +
		`<yr-sidebar-nav>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("layout-dashboard", "", mspIconSize) +
		` Dashboard</yr-sidebar-link>` +
		`<yr-sidebar-link href="#">` +
		exampleIcon("building-2", "", mspIconSize) +
		` Clients</yr-sidebar-link>` +
		`</yr-sidebar-nav>` +
		`<yr-sidebar-nav>` +
		`<yr-sidebar-org name="Acme Corp"></yr-sidebar-org>` +
		strings.Join(links, "") +
		`</yr-sidebar-nav>` +
		mspSidebarFooter() +
		`</yr-sidebar>`
}
