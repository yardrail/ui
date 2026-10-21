package ui

import "github.com/a-h/templ"

// CopyFieldProps configures a CopyField.
type CopyFieldProps struct {
	Attrs templ.Attributes
}

// CopyField renders a code display with a copy-to-clipboard button.
func CopyField(value string, p CopyFieldProps) templ.Component {
	return copyField(value, &p)
}
