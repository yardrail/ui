package ui

import "github.com/a-h/templ"

func statExamples() []Example {
	return []Example{
		{Name: "row", Component: templ.Raw(
			`<div style="display:flex;gap:1rem">` +
				`<yr-stat number="500+" label="Templates"></yr-stat>` +
				`<yr-stat number="200+" label="Authors"></yr-stat>` +
				`<yr-stat number="10k+" label="Installs"></yr-stat>` +
				`</div>`,
		)},
	}
}
