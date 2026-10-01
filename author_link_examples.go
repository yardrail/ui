package ui

import "github.com/a-h/templ"

func authorLinkExamples() []Example {
	return []Example{
		{Name: "sizes", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:1.5rem">` +
				`<div style="display:flex;flex-direction:column;gap:0.5rem">` +
				`<yr-author-link href="/authors/yardrail" name="Yardrail" handle="yardrail" avatar="YR" verified></yr-author-link>` +
				`<code style="font-size:0.75rem;color:#5C5A59">default</code>` +
				`</div>` +
				`<div style="display:flex;flex-direction:column;gap:0.5rem">` +
				`<yr-author-link href="/authors/yardrail" name="Yardrail" handle="yardrail" avatar="YR" verified size="lg"></yr-author-link>` +
				`<code style="font-size:0.75rem;color:#5C5A59">size="lg"</code>` +
				`</div>` +
				`</div>`,
		)},
	}
}
