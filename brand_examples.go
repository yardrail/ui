package ui

import "github.com/a-h/templ"

func brandExamples() []Example {
	return []Example{
		{Name: "default", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:1.5rem">` +
				`<a href="#" class="yr-brand">Yardrail</a>` +
				`<a href="#" class="yr-brand">Yardrail <span class="yr-brand-sub">Depot</span></a>` +
				`</div>`,
		)},
	}
}
