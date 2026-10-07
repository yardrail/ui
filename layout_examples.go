package ui

import "github.com/a-h/templ"

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

// placeholderItem renders a list item card labelled label, for examples rendered with AsList.
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
		{Name: exampleDefault, Component: withChildren(Stack(StackProps{}), placeholders(three...)...), DisplayURL: ""},
		{
			Name:       "gap-8",
			Component:  withChildren(Stack(StackProps{Gap: Space8}), placeholders(three...)...),
			DisplayURL: "",
		},
		{
			Name:       "align-center",
			Component:  withChildren(Stack(StackProps{Align: AlignCenter}), placeholders(three...)...),
			DisplayURL: "",
		},
		{
			Name:       exampleAsList,
			Component:  withChildren(Stack(StackProps{As: AsList}), placeholderItems(three...)...),
			DisplayURL: "",
		},
	}
}

func clusterExamples() []Example {
	tags := []string{"Billing", "Support", "Onboarding", "Security", "Reporting"}

	return []Example{
		{
			Name:       exampleDefault,
			Component:  withChildren(Cluster(ClusterProps{}), placeholders(tags...)...),
			DisplayURL: "",
		},
		{
			Name: "justify-between",
			Component: withChildren(Cluster(ClusterProps{Justify: JustifyBetween}),
				placeholders(placeholderOne, placeholderTwo, placeholderThree)...),
			DisplayURL: "",
		},
		{
			Name:       exampleAsList,
			Component:  withChildren(Cluster(ClusterProps{As: AsList}), placeholderItems(tags...)...),
			DisplayURL: "",
		},
	}
}
