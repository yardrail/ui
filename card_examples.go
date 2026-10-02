package ui

import "github.com/a-h/templ"

func cardExamples() []Example {
	return []Example{
		{Name: "with-title", Component: templ.Raw(
			`<div style="max-width:22rem">` +
				`<div class="yr-card">` +
				`<h3 class="yr-card-title">Ready to go?</h3>` +
				`<p style="margin:0">Everything is connected. Finish setup to start using this workflow.</p>` +
				`</div>` +
				`</div>`,
		)},
		{Name: "header-footer", Component: templ.Raw(
			`<div style="max-width:22rem">` +
				`<div class="yr-card">` +
				`<div class="yr-card-header">` +
				`<h3 class="yr-card-title">Order summary</h3>` +
				`<yr-pill tone="green">Paid</yr-pill>` +
				`</div>` +
				`<p style="margin:0">2 templates, billed to Ellery Logistics.</p>` +
				`<div class="yr-card-footer">` +
				`<span>Total</span>` +
				`<strong>$48.00</strong>` +
				`</div>` +
				`</div>` +
				`</div>`,
		)},
		{Name: "interactive", Component: templ.Raw(
			`<div style="display:grid;grid-template-columns:repeat(2,minmax(0,16rem));gap:1rem">` +
				`<a href="#" class="yr-card yr-card--interactive">Marketing</a>` +
				`<a href="#" class="yr-card yr-card--interactive">Analytics</a>` +
				`</div>`,
		)},
	}
}
