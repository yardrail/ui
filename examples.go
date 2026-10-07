package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// Example names shared by several components.
const (
	exampleDefault     = "default"
	exampleWithContent = "with-content"
	exampleCollapsed   = "collapsed"
)

// Example is one named rendering of a component. The snapshot tests compare it against
// testdata/<component>/<name>.html, and the gallery serves it at /<component>/<name>.
type Example struct {
	// Component is the rendered example.
	Component templ.Component
	// Name is a lowercase kebab slug, used as a file name and a URL segment.
	Name string
	// DisplayURL is an optional realistic URL shown in the gallery's browser chrome.
	DisplayURL string
}

// ExampleGroup is every example of one component.
type ExampleGroup struct {
	// Component is the component's lowercase kebab slug, used as a directory name and a URL segment.
	Component string
	// Examples are the component's examples, in display order.
	Examples []Example
}

// ExampleGroups returns every component's examples, in a fixed order.
func ExampleGroups() []ExampleGroup {
	return []ExampleGroup{
		{Component: "stack", Examples: stackExamples()},
		{Component: "cluster", Examples: clusterExamples()},
		{Component: "grid", Examples: gridExamples()},
		{Component: "center", Examples: centerExamples()},
		{Component: "split", Examples: splitExamples()},
		{Component: "button", Examples: buttonExamples()},
		{Component: "pill", Examples: pillExamples()},
		{Component: "avatar", Examples: avatarExamples()},
		{Component: "icon-badge", Examples: iconBadgeExamples()},
		{Component: "stat", Examples: statExamples()},
		{Component: "brand", Examples: brandExamples()},
		{Component: "rating", Examples: ratingExamples()},
		{Component: "field", Examples: fieldExamples()},
		{Component: "form", Examples: formExamples()},
		{Component: "table", Examples: tableExamples()},
		{Component: "row-edit", Examples: rowEditExamples()},
		{Component: "author-link", Examples: authorLinkExamples()},
		{Component: "action-card", Examples: actionCardExamples()},
		{Component: "activity-row", Examples: activityRowExamples()},
		{Component: "alert", Examples: alertExamples()},
		{Component: "auth-card", Examples: authCardExamples()},
		{Component: "breadcrumb", Examples: breadcrumbExamples()},
		{Component: "card", Examples: cardExamples()},
		{Component: "connector-pill", Examples: connectorPillExamples()},
		{Component: "copy-field", Examples: copyFieldExamples()},
		{Component: "create-panel", Examples: createPanelExamples()},
		{Component: "dropdown-menu", Examples: dropdownMenuExamples()},
		{Component: "empty-state", Examples: emptyStateExamples()},
		{Component: "list-row", Examples: listRowExamples()},
		{Component: "live-duration", Examples: liveDurationExamples()},
		{Component: "meta-row", Examples: metaRowExamples()},
		{Component: "nav-shell", Examples: navShellExamples()},
		{Component: "review-card", Examples: reviewCardExamples()},
		{Component: "search", Examples: searchExamples()},
		{Component: "settings-section", Examples: settingsSectionExamples()},
		{Component: "tabs", Examples: tabsExamples()},
		{Component: "template-card", Examples: templateCardExamples()},
		{Component: "page-workspace", Examples: workspaceExamples()},
		{Component: "page-auth", Examples: pageAuthExamples()},
		{Component: "page-setup", Examples: pageSetupExamples()},
		{Component: "navbar", Examples: navbarExamples()},
		{Component: "sidebar", Examples: sidebarExamples()},
		{Component: "drawer", Examples: drawerExamples()},
		{Component: "setup-card", Examples: setupCardExamples()},
		{Component: "msp-dashboard", Examples: mspDashboardExamples()},
		{Component: "msp-clients", Examples: mspClientsExamples()},
		{Component: "msp-client-overview", Examples: mspClientOverviewExamples()},
		{Component: "msp-client-users", Examples: mspClientUsersExamples()},
		{Component: "msp-client-connectors", Examples: mspClientConnectorsExamples()},
		{Component: "msp-client-settings", Examples: mspClientSettingsExamples()},
	}
}

// withChildren renders parent with children as its { children... }, the Go equivalent of
// @parent { children } in templ.
func withChildren(parent templ.Component, children ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return parent.Render(templ.WithChildren(ctx, templ.Join(children...)), w)
	})
}

// text renders s, escaped, as a child component.
func text(s string) templ.Component {
	return templ.Raw(templ.EscapeString(s))
}
