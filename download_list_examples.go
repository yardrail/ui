package ui

import "github.com/a-h/templ"

func downloadListExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<div style="max-width:28rem">` +
				`<yr-download-list>` +
				`<a href="#" download>` +
				exampleIcon("file-text", "", 16) +
				` newsletter-draft.html` +
				`<span class="yr-download-size">84 KB</span>` +
				`</a>` +
				`<a href="#" download>` +
				exampleIcon("file-spreadsheet", "", 16) +
				` recipient-list.csv` +
				`<span class="yr-download-size">156 KB</span>` +
				`</a>` +
				`<a href="#" download>` +
				exampleIcon("file-image", "", 16) +
				` hero-banner.png` +
				`<span class="yr-download-size">2.4 MB</span>` +
				`</a>` +
				`</yr-download-list>` +
				`</div>`,
		)},
	}
}
