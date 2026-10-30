package ui

import (
	"fmt"
	"testing"
)

func TestRenderCard(t *testing.T) {
	t.Parallel()

	for i, interactive := range []bool{false, true} {
		html := renderString(t, withChildren(
			Card(&CardProps{Interactive: interactive}),
			text("card content"),
		))
		writeClassCoverage(t, "card", fmt.Sprintf("%03d", i), html)
	}
}
