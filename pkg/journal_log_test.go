package ui

import "testing"

func TestRenderJournalLog(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		JournalLog(&JournalLogProps{}),
		text("log entries"),
	))
	writeClassCoverage(t, "journal_log", "default", html)
}
