package ui

import "github.com/a-h/templ"

// AlertProps configures an Alert.
type AlertProps struct {
	Tone  string
	Attrs templ.Attributes
}

// Alert renders a banner/flash message. Children are the message content.
func Alert(p AlertProps) templ.Component {
	return alert(&p)
}
