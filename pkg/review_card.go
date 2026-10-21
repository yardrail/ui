package ui

import "github.com/a-h/templ"

// ReviewCardProps configures a ReviewCard.
type ReviewCardProps struct {
	Attrs templ.Attributes
}

// ReviewCard renders a review card. Children are the review content.
func ReviewCard(p *ReviewCardProps) templ.Component {
	return reviewCard(p)
}
