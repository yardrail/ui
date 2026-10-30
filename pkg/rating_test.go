package ui

import "testing"

func TestRenderRating(t *testing.T) {
	t.Parallel()

	html := renderString(t, Rating("4.5", &RatingProps{Reviews: "12"}))
	writeClassCoverage(t, "rating", "default", html)
}
