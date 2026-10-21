package main

import (
	"github.com/a-h/templ"
	ui "github.com/yardrail/ui/pkg"
)

// Example names shared by the layout primitives.
const (
	exampleAsList = "as-list"
)

// Placeholder labels for layout example children.
const (
	placeholderOne   = "One"
	placeholderTwo   = "Two"
	placeholderThree = "Three"
)

// placeholder renders a card labelled label, so a layout example shows its gaps.
func placeholder(label string) templ.Component {
	return templ.Raw(`<div class="yr-card">` + templ.EscapeString(label) + `</div>`)
}

// placeholderItem renders a list item card labelled label, for examples rendered with ui.AsList.
func placeholderItem(label string) templ.Component {
	return templ.Raw(`<li class="yr-card">` + templ.EscapeString(label) + `</li>`)
}

// placeholders renders one placeholder card per label.
func placeholders(labels ...string) []templ.Component {
	out := make([]templ.Component, 0, len(labels))
	for _, l := range labels {
		out = append(out, placeholder(l))
	}

	return out
}

// placeholderItems renders one placeholder list item per label.
func placeholderItems(labels ...string) []templ.Component {
	out := make([]templ.Component, 0, len(labels))
	for _, l := range labels {
		out = append(out, placeholderItem(l))
	}

	return out
}

func stackExamples() []Example {
	three := []string{placeholderOne, placeholderTwo, placeholderThree}

	return []Example{
		{
			Name:       exampleDefault,
			Component:  withChildren(ui.Stack(ui.StackProps{}), placeholders(three...)...),
			DisplayURL: "",
		},
		{
			Name:       "gap-8",
			Component:  withChildren(ui.Stack(ui.StackProps{Gap: ui.Space8}), placeholders(three...)...),
			DisplayURL: "",
		},
		{
			Name:       "align-center",
			Component:  withChildren(ui.Stack(ui.StackProps{Align: ui.AlignCenter}), placeholders(three...)...),
			DisplayURL: "",
		},
		{
			Name:       exampleAsList,
			Component:  withChildren(ui.Stack(ui.StackProps{As: ui.AsList}), placeholderItems(three...)...),
			DisplayURL: "",
		},
	}
}

func clusterExamples() []Example {
	tags := []string{"Billing", "Support", "Onboarding", "Security", "Reporting"}

	return []Example{
		{
			Name:       exampleDefault,
			Component:  withChildren(ui.Cluster(ui.ClusterProps{}), placeholders(tags...)...),
			DisplayURL: "",
		},
		{
			Name: "justify-between",
			Component: withChildren(ui.Cluster(ui.ClusterProps{Justify: ui.JustifyBetween}),
				placeholders(placeholderOne, placeholderTwo, placeholderThree)...),
			DisplayURL: "",
		},
		{
			Name:       exampleAsList,
			Component:  withChildren(ui.Cluster(ui.ClusterProps{As: ui.AsList}), placeholderItems(tags...)...),
			DisplayURL: "",
		},
	}
}

func gridExamples() []Example {
	six := []string{placeholderOne, placeholderTwo, placeholderThree, "Four", "Five", "Six"}

	return []Example{
		{
			Name:       exampleDefault,
			Component:  withChildren(ui.Grid(ui.GridProps{}), placeholders(six...)...),
			DisplayURL: "",
		},
		{
			Name:       "item-sm",
			Component:  withChildren(ui.Grid(ui.GridProps{MinItemWidth: ui.ItemWidthSm}), placeholders(six...)...),
			DisplayURL: "",
		},
		{
			Name:       exampleAsList,
			Component:  withChildren(ui.Grid(ui.GridProps{As: ui.AsList}), placeholderItems(six...)...),
			DisplayURL: "",
		},
	}
}

// centerProse is the placeholder copy for Center examples.
const centerProse = "Connect a data source to start syncing. Each connector runs on its own schedule, and " +
	"you can pause or remove it at any time from the connector settings."

func centerExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			Component:  withChildren(ui.Center(ui.CenterProps{}), placeholder(centerProse)),
			DisplayURL: "",
		},
		{
			Name:       "measure-narrow",
			Component:  withChildren(ui.Center(ui.CenterProps{Measure: ui.MeasureNarrow}), placeholder(centerProse)),
			DisplayURL: "",
		},
	}
}

// Placeholder labels for the regions of Split examples.
const (
	placeholderAside = "Aside"
	placeholderMain  = "Main"
)

func splitExamples() []Example {
	aside := withChildren(ui.SplitAside(ui.RegionProps{}), placeholder(placeholderAside))
	main := withChildren(ui.SplitMain(ui.RegionProps{}), placeholder(placeholderMain))

	return []Example{
		{Name: "aside-first", Component: withChildren(ui.Split(ui.SplitProps{}), aside, main), DisplayURL: ""},
		{Name: "main-first", Component: withChildren(ui.Split(ui.SplitProps{}), main, aside), DisplayURL: ""},
		{
			Name:       "side-lg",
			Component:  withChildren(ui.Split(ui.SplitProps{SideWidth: ui.SideWidthLg}), aside, main),
			DisplayURL: "",
		},
	}
}

func dividerExamples() []Example {
	return []Example{
		{
			Name: "in-stack",
			Component: withChildren(ui.Stack(ui.StackProps{}),
				placeholder(placeholderOne), ui.Divider(), placeholder(placeholderTwo)),
			DisplayURL: "",
		},
		{
			Name: "in-list",
			Component: withChildren(ui.Stack(ui.StackProps{As: ui.AsList}),
				placeholderItem(placeholderOne), ui.Divider(), placeholderItem(placeholderTwo)),
			DisplayURL: "",
		},
	}
}
