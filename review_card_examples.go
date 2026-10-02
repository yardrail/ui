package ui

import "github.com/a-h/templ"

func reviewCardExamples() []Example {
	return []Example{
		{Name: "reviews", Component: templ.Raw(
			`<div style="display:flex;flex-direction:column;gap:0.75rem;max-width:36rem">` +
				`<div class="yr-review-card">` +
				`<div class="yr-review-card-header">` +
				`<yr-rating value="5"></yr-rating>` +
				`<span class="yr-review-card-author">Jordan Ellery</span>` +
				`</div>` +
				`<p class="yr-review-card-body">Set up in ten minutes and it has run every night since.` +
				` The documentation covers the two connector scopes you need.</p>` +
				`</div>` +
				`<div class="yr-review-card">` +
				`<div class="yr-review-card-header">` +
				`<yr-rating value="3"></yr-rating>` +
				`<span class="yr-review-card-author">Sam Okafor</span>` +
				`</div>` +
				`<p class="yr-review-card-body">Works, but the Slack step needs a manual channel ID.</p>` +
				`</div>` +
				`</div>`,
		)},
	}
}
