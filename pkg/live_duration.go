package ui

import "github.com/a-h/templ"

// LiveDurationProps configures a LiveDuration.
type LiveDurationProps struct {
	Attrs templ.Attributes
}

// LiveDuration renders a live HH:MM:SS counter from the given ISO 8601 start time.
func LiveDuration(startedAt string, p *LiveDurationProps) templ.Component {
	return liveDuration(startedAt, p)
}
