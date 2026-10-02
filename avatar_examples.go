package ui

import "github.com/a-h/templ"

func avatarExamples() []Example {
	return []Example{
		{Name: "sizes", Component: templ.Raw(
			`<div style="display:flex;align-items:start;gap:2rem">` +
				`<div style="display:flex;flex-direction:column;align-items:center;gap:0.5rem">` +
				`<yr-avatar>BL</yr-avatar>` +
				`<code style="font-size:0.75rem;color:#5C5A59">default</code>` +
				`</div>` +
				`<div style="display:flex;flex-direction:column;align-items:center;gap:0.5rem">` +
				`<yr-avatar size="lg">BL</yr-avatar>` +
				`<code style="font-size:0.75rem;color:#5C5A59">size="lg"</code>` +
				`</div>` +
				`</div>`,
		)},
	}
}
