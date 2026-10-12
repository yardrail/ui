package ui

import "github.com/a-h/templ"

func sectionHeaderExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<div style="max-width:40rem">` +
				`<yr-section-header>Recent Runs</yr-section-header>` +
				`<yr-section-header>Needs You</yr-section-header>` +
				`</div>`,
		)},
		{Name: "with-action", Component: templ.Raw(
			`<div style="max-width:40rem">` +
				`<yr-section-header>Needs You <a href="#">View all</a></yr-section-header>` +
				`<yr-section-header>Recent Activity <a href="#">See more</a></yr-section-header>` +
				`</div>`,
		)},
	}
}
