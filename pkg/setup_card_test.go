package ui

import "testing"

func TestRenderSetupCard(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SetupCard(&SetupCardProps{}), text("content")))
	writeClassCoverage(t, "setup_card", "default", html)
}

func TestRenderSetupCardProgress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		steps   int
		current int
	}{
		{"step1of3", 3, 1},
		{"step2of3", 3, 2},
		{"step3of3", 3, 3},
	}

	for _, tt := range tests {
		html := renderString(t, SetupCardProgress(&SetupCardProgressProps{
			Steps:   tt.steps,
			Current: tt.current,
		}))
		writeClassCoverage(t, "setup_card_progress", tt.name, html)
	}
}

func TestRenderSetupCardHeader(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SetupCardHeader(&SetupCardHeaderProps{}), text("header")))
	writeClassCoverage(t, "setup_card_header", "default", html)
}

func TestRenderSetupCardBody(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SetupCardBody(&SetupCardBodyProps{}), text("body")))
	writeClassCoverage(t, "setup_card_body", "default", html)
}

func TestRenderSetupCardFooter(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SetupCardFooter(&SetupCardFooterProps{}), text("footer")))
	writeClassCoverage(t, "setup_card_footer", "default", html)
}
