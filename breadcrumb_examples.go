package ui

import "github.com/a-h/templ"

func breadcrumbExamples() []Example {
	return []Example{
		{Name: "trail", Component: templ.Raw(
			`<nav class="yr-breadcrumb">` +
				`<a href="#">Workflows</a>` +
				`<span class="yr-breadcrumb-sep">/</span>` +
				`<a href="#">Refund triage</a>` +
				`<span class="yr-breadcrumb-sep">/</span>` +
				`<span class="yr-breadcrumb-current">Run 8f3a2c1d</span>` +
				`</nav>`,
		)},
		{Name: "back-link", Component: templ.Raw(
			`<nav class="yr-breadcrumb"><a href="#">← Browse</a></nav>`,
		)},
	}
}
