package ui

import "github.com/a-h/templ"

const outputRowIconSize = 20

func outputRowExamples() []Example {
	return []Example{
		{Name: "rows", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:44rem">` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("file-text", "", outputRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">newsletter-digest-2026-10-11.pdf</div>` +
				`<div class="yr-list-row-desc">Generated report, 2.4 MB</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button yr-button--secondary" download>Download</a>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("file-spreadsheet", "", outputRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">lead-scores-export.csv</div>` +
				`<div class="yr-list-row-desc">Scored leads, 156 KB</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button yr-button--secondary" download>Download</a>` +
				`</div>` +
				`</div>` +

				`<div class="yr-list-row">` +
				`<div class="yr-list-row-icon">` +
				exampleIcon("file", "", outputRowIconSize) + `</div>` +
				`<div class="yr-list-row-main">` +
				`<div class="yr-list-row-name">social-posts-batch.json</div>` +
				`<div class="yr-list-row-desc">Post queue, 48 KB</div>` +
				`</div>` +
				`<div class="yr-list-row-actions">` +
				`<a href="#" class="yr-button yr-button--secondary" download>Download</a>` +
				`</div>` +
				`</div>` +

				`</div>`,
		)},
	}
}
