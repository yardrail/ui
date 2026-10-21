package ui

import (
	"bytes"
	"context"
	"strconv"
	"strings"
)

// exampleLogo is a placeholder connector logo for examples that need an <img>.
const exampleLogo = `data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'%3E` +
	`%3Crect width='24' height='24' rx='5' fill='%231A5959'/%3E%3C/svg%3E`

// exampleIconSize is the width and height the icon files are drawn at.
const exampleIconSize = 24

// exampleIcon returns the named icon's SVG markup at size pixels. A non-empty attrs replaces the
// icon's own class attribute, for markup that styles the <svg> element itself.
func exampleIcon(name, attrs string, size int) string {
	var buf bytes.Buffer

	err := Icon(name).Render(context.Background(), &buf)
	if err != nil {
		panic("render icon " + name + ": " + err.Error())
	}

	svg := buf.String()

	if attrs != "" {
		start := strings.Index(svg, `class="`)
		if start < 0 {
			panic("icon " + name + " has no class attribute")
		}

		end := start + len(`class="`) + strings.Index(svg[start+len(`class="`):], `"`) + 1
		svg = svg[:start] + attrs + svg[end:]
	}

	if size != exampleIconSize {
		from, to := strconv.Itoa(exampleIconSize), strconv.Itoa(size)
		svg = strings.Replace(svg, `width="`+from+`"`, `width="`+to+`"`, 1)
		svg = strings.Replace(svg, `height="`+from+`"`, `height="`+to+`"`, 1)
	}

	return svg
}
