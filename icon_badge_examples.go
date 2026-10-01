package ui

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func iconBadgeItem(tone, iconName string) string {
	var buf bytes.Buffer

	_ = Icon(iconName).Render(context.Background(), &buf)

	return `<div style="display:flex;flex-direction:column;align-items:center;gap:0.5rem">` +
		`<div class="yr-icon-badge yr-icon-badge--` + tone + `">` + buf.String() + `</div>` +
		`<code style="font-size:0.75rem;color:#5C5A59">` + tone + `</code></div>`
}

func iconBadgeExamples() []Example {
	return []Example{
		{Name: "tones", Component: templ.Raw(
			`<div style="display:flex;flex-wrap:wrap;gap:1rem">` +
				iconBadgeItem("navy", "pencil") +
				iconBadgeItem("amber", "mail") +
				iconBadgeItem("blue", "star") +
				iconBadgeItem("green", "check-circle") +
				iconBadgeItem("purple", "shield") +
				iconBadgeItem("pink", "user") +
				iconBadgeItem("red", "zap") +
				iconBadgeItem("slate", "settings") +
				`</div>`,
		)},
	}
}
