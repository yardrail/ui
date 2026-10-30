package ui

import "testing"

func TestRenderReviewCard(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(ReviewCard(&ReviewCardProps{}), text("review")))
	writeClassCoverage(t, "review_card", "default", html)
}
