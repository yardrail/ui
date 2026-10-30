package ui

// Tone is the semantic color of a component.
type Tone struct{ v string }

// Tone values.
var (
	// ToneAccent is the primary brand accent.
	ToneAccent = Tone{v: "accent"}
	// ToneAmber is an amber category color.
	ToneAmber = Tone{v: "amber"}
	// ToneBlue is a blue category color.
	ToneBlue = Tone{v: "blue"}
	// ToneGreen is a green category color.
	ToneGreen = Tone{v: "green"}
	// TonePurple is a purple category color.
	TonePurple = Tone{v: "purple"}
	// TonePink is a pink category color.
	TonePink = Tone{v: "pink"}
	// ToneSuccess indicates a positive outcome.
	ToneSuccess = Tone{v: "success"}
	// ToneWarning indicates a cautionary state.
	ToneWarning = Tone{v: "warning"}
	// ToneDanger indicates a destructive or error state.
	ToneDanger = Tone{v: "danger"}
	// ToneInfo indicates a neutral informational state.
	ToneInfo = Tone{v: "info"}
)

// String returns the tone value for use in HTML attributes.
func (t Tone) String() string {
	return t.v
}

// All returns every Tone value.
func (Tone) All() []Tone {
	return []Tone{
		ToneAccent, ToneAmber, ToneBlue, ToneGreen, TonePurple,
		TonePink, ToneSuccess, ToneWarning, ToneDanger, ToneInfo,
	}
}
