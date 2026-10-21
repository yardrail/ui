package ui

import "github.com/a-h/templ"

// JournalLogProps configures a JournalLog.
type JournalLogProps struct {
	Attrs templ.Attributes
}

// JournalLog renders a scrollable container for workflow run journal entries. Children are the entries.
func JournalLog(p *JournalLogProps) templ.Component {
	return journalLog(p)
}
