package ui

import "github.com/a-h/templ"

func templateCardExamples() []Example {
	return []Example{
		{Name: "grid", Component: templ.Raw(
			`<div style="display:grid;grid-template-columns:repeat(2,minmax(0,20rem));gap:1rem">` +
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Refund triage</h3>` +
				`<p class="yr-template-card-desc">Routes refund requests to the right approver and posts` +
				` the decision back to the customer thread.</p>` +
				`<div class="yr-template-card-connectors">` +
				`<div class="yr-connector-pill--icon" title="Slack"><img src="` + exampleLogo + `" alt="Slack"></div>` +
				`<div class="yr-connector-pill--icon" title="HubSpot">HS</div>` +
				`</div>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="accent">Verified</yr-pill>` +
				`<yr-rating value="4.5" reviews="42"></yr-rating>` +
				`</div>` +
				`<span class="yr-template-card-price">$24</span>` +
				`</div>` +
				`</a>` +
				`<a href="#" class="yr-card yr-card--interactive yr-template-card">` +
				`<h3 class="yr-template-card-name">Weekly pipeline report</h3>` +
				`<p class="yr-template-card-desc">Summarizes open deals every Monday.</p>` +
				`<div class="yr-card-footer">` +
				`<div class="yr-template-card-footer-left">` +
				`<yr-pill tone="purple">Community</yr-pill>` +
				`<span class="yr-template-card-install-count">1.2k installs</span>` +
				`</div>` +
				`<span class="yr-template-card-price">Free</span>` +
				`</div>` +
				`</a>` +
				`</div>`,
		)},
	}
}
