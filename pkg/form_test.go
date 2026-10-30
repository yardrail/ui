package ui

import "testing"

func TestRenderForm(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Form(FormProps{Action: "/submit", ID: "test-form"}),
		text("form fields"),
	))
	writeClassCoverage(t, "form", "default", html)
}
