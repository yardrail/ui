package ui

import (
	"fmt"
	"testing"
)

func TestRenderActionCard(t *testing.T) {
	t.Parallel()

	i := 0

	for _, urgent := range []bool{false, true} {
		for _, urgentDanger := range []bool{false, true} {
			html := renderString(t, withChildren(
				ActionCard(&ActionCardProps{
					Urgent:       urgent,
					UrgentDanger: urgentDanger,
				}),
				text("action content"),
			))
			writeClassCoverage(t, "action_card", fmt.Sprintf("%03d", i), html)
			i++
		}
	}
}
