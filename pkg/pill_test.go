package ui

import (
	"fmt"
	"testing"
)

func TestRenderPill(t *testing.T) {
	t.Parallel()

	i := 0

	for _, tone := range (Tone{}).All() {
		for _, pill := range []bool{false, true} {
			html := renderString(t, Pill("Label", &PillProps{
				Tone: tone,
				Pill: pill,
			}))
			writeClassCoverage(t, "pill", fmt.Sprintf("%03d", i), html)
			i++
		}
	}
}
