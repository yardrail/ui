package ui

import "github.com/a-h/templ"

// HitlPromptProps configures a HitlPrompt.
type HitlPromptProps struct {
	Tone  string
	Attrs templ.Attributes
}

// HitlPrompt renders a human-in-the-loop prompt card. Children are the prompt content.
func HitlPrompt(p *HitlPromptProps) templ.Component {
	return hitlPrompt(p)
}
