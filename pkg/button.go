package ui

import "github.com/a-h/templ"

// ButtonType is the native type of a Button. The zero value renders type="button", so a Button
// inside a form submits only when it asks to.
type ButtonType struct{ v string }

// Button types beyond the zero value.
var (
	// ButtonSubmit renders type="submit".
	ButtonSubmit = ButtonType{v: "submit"}
	// ButtonReset renders type="reset".
	ButtonReset = ButtonType{v: "reset"}
)

// String returns the value of the type attribute.
func (t ButtonType) String() string {
	if t.v == "" {
		return "button"
	}

	return t.v
}

// ButtonVariant is the visual style of a Button. The zero value is ButtonPrimary.
type ButtonVariant struct{ v string }

// Button variants.
var (
	// ButtonPrimary is the filled default action.
	ButtonPrimary = ButtonVariant{}
	// ButtonSecondary is the outlined alternative action.
	ButtonSecondary = ButtonVariant{v: "secondary"}
	// ButtonLink looks like a text link.
	ButtonLink = ButtonVariant{v: "link"}
	// ButtonDanger is a filled destructive action.
	ButtonDanger = ButtonVariant{v: "danger"}
	// ButtonDangerQuiet is a low-emphasis destructive action, for example Remove in a table row.
	ButtonDangerQuiet = ButtonVariant{v: "danger-quiet"}
)

// buttonClass is the base class every Button carries.
const buttonClass = "yr-button"

// class returns the class attribute for the variant.
func (v ButtonVariant) class() string {
	if v.v == "" {
		return buttonClass
	}

	return buttonClass + " " + buttonClass + "--" + v.v
}

// ButtonProps holds a Button's optional settings; every zero value is the default.
type ButtonProps struct {
	// Attrs holds htmx and script hooks, filtered by the Attrs allow-list.
	Attrs templ.Attributes
	// Type is the native button type; the zero value is type="button".
	Type ButtonType
	// Variant is the visual style; the zero value is ButtonPrimary.
	Variant ButtonVariant
	// Name is the name submitted with the form, when set.
	Name string
	// Value is the value submitted with the form, when set.
	Value string
	// Confirm, when set, asks the user to confirm before the htmx request is sent.
	Confirm string
	// Disabled renders the button disabled.
	Disabled bool
}
