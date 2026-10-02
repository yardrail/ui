package ui

import "github.com/a-h/templ"

func pillExamples() []Example {
	return []Example{
		{Name: "tones", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center">` +
				`<yr-pill tone="accent">Accent</yr-pill>` +
				`<yr-pill tone="amber">Amber</yr-pill>` +
				`<yr-pill tone="blue">Blue</yr-pill>` +
				`<yr-pill tone="green">Green</yr-pill>` +
				`<yr-pill tone="purple">Purple</yr-pill>` +
				`<yr-pill tone="pink">Pink</yr-pill>` +
				`<yr-pill tone="danger">Danger</yr-pill>` +
				`<yr-pill>No tone</yr-pill>` +
				`</div>`,
		)},
		{Name: "status", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center">` +
				`<yr-pill tone="success" shape="pill">Running</yr-pill>` +
				`<yr-pill tone="warning" shape="pill">Deploying</yr-pill>` +
				`<yr-pill tone="danger" shape="pill">Failed</yr-pill>` +
				`<yr-pill shape="pill">Idle</yr-pill>` +
				`</div>`,
		)},
		{Name: "interactive", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.5rem">` +
				`<a href="#" class="yr-pill yr-pill--pill yr-pill--tone-accent is-active">All</a>` +
				`<a href="#" class="yr-pill yr-pill--pill">Marketing</a>` +
				`<a href="#" class="yr-pill yr-pill--pill">Analytics</a>` +
				`<a href="#" class="yr-pill yr-pill--pill">Sales</a>` +
				`</div>`,
		)},
	}
}
