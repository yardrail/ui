package ui

import "github.com/a-h/templ"

const downloadListIconSize = 16

func downloadListExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: templ.Raw(
			`<div style="max-width:28rem">` +
				`<yr-download-list>` +
				`<a href="#" download>` +
				exampleIcon("file-text", "", downloadListIconSize) +
				` newsletter-draft.html` +
				`<span class="yr-download-size">84 KB</span>` +
				`</a>` +
				`<a href="#" download>` +
				exampleIcon("file-spreadsheet", "", downloadListIconSize) +
				` recipient-list.csv` +
				`<span class="yr-download-size">156 KB</span>` +
				`</a>` +
				`<a href="#" download>` +
				exampleIcon("file-image", "", downloadListIconSize) +
				` hero-banner.png` +
				`<span class="yr-download-size">2.4 MB</span>` +
				`</a>` +
				`</yr-download-list>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
