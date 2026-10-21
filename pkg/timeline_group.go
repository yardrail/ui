package ui

import "github.com/a-h/templ"

// TimelineGroupProps configures a TimelineGroup.
type TimelineGroupProps struct {
	Attrs templ.Attributes
}

// TimelineGroup renders a date-labeled container with a vertical time rail. Children are the group content.
func TimelineGroup(p *TimelineGroupProps) templ.Component {
	return timelineGroup(p)
}

// TimelineEntriesProps configures a TimelineEntries.
type TimelineEntriesProps struct {
	Attrs templ.Attributes
}

// TimelineEntries renders a container for timeline entry rows. Children are the entries.
func TimelineEntries(p *TimelineEntriesProps) templ.Component {
	return timelineEntries(p)
}
