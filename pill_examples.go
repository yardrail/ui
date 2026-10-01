package ui

import "github.com/a-h/templ"

func pillExamples() []Example {
	return []Example{
		{Name: "tones", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center">` +
				`<yr-pill tone="navy">Navy</yr-pill>` +
				`<yr-pill tone="amber">Amber</yr-pill>` +
				`<yr-pill tone="blue">Blue</yr-pill>` +
				`<yr-pill tone="green">Green</yr-pill>` +
				`<yr-pill tone="purple">Purple</yr-pill>` +
				`<yr-pill tone="pink">Pink</yr-pill>` +
				`<yr-pill tone="red">Red</yr-pill>` +
				`<yr-pill tone="slate">Slate</yr-pill>` +
				`</div>`,
		)},
		{Name: "signals", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.75rem;align-items:center">` +
				`<yr-pill tone="active" shape="pill">Running</yr-pill>` +
				`<yr-pill tone="pending" shape="pill">Deploying</yr-pill>` +
				`<yr-pill tone="stopped" shape="pill">Failed</yr-pill>` +
				`<yr-pill tone="idle" shape="pill">Idle</yr-pill>` +
				`</div>`,
		)},
		{Name: "interactive", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:0.5rem">` +
				`<a href="#" class="yr-pill yr-pill--pill yr-pill--tone-navy is-active">All</a>` +
				`<a href="#" class="yr-pill yr-pill--pill yr-pill--tone-slate">Marketing</a>` +
				`<a href="#" class="yr-pill yr-pill--pill yr-pill--tone-slate">Analytics</a>` +
				`<a href="#" class="yr-pill yr-pill--pill yr-pill--tone-slate">Sales</a>` +
				`</div>`,
		)},
	}
}
