package ui

import "github.com/a-h/templ"

// RatingProps configures a Rating.
type RatingProps struct {
	Reviews string
	Attrs   templ.Attributes
}

// Rating renders a 5-star rating with optional review count.
func Rating(value string, p RatingProps) templ.Component {
	return rating(value, &p)
}
