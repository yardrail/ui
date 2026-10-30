package ui

import "testing"

func TestRenderTimelineGroup(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(TimelineGroup(&TimelineGroupProps{}), text("group")))
	writeClassCoverage(t, "timeline_group", "default", html)
}

func TestRenderTimelineEntries(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(TimelineEntries(&TimelineEntriesProps{}), text("entries")))
	writeClassCoverage(t, "timeline_entries", "default", html)
}
