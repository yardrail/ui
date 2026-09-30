package ui

import (
	"context"
	"strings"

	"github.com/a-h/templ"
)

// FieldType is the input type of a Field. The zero value renders type="text".
type FieldType struct{ v string }

// Field types beyond the zero value.
var (
	// FieldEmail renders type="email".
	FieldEmail = FieldType{v: "email"}
	// FieldPassword renders type="password".
	FieldPassword = FieldType{v: "password"}
	// FieldURL renders type="url".
	FieldURL = FieldType{v: "url"}
	// FieldSearch renders type="search".
	FieldSearch = FieldType{v: "search"}
)

// String returns the value of the type attribute.
func (t FieldType) String() string {
	if t.v == "" {
		return "text"
	}

	return t.v
}

// Autocomplete is the autofill token of a Field. The zero value renders no autocomplete attribute.
type Autocomplete struct{ v string }

// Autocomplete tokens.
var (
	// AutocompleteOff turns autofill off.
	AutocompleteOff = Autocomplete{v: "off"}
	// AutocompleteName is a person's full name.
	AutocompleteName = Autocomplete{v: "name"}
	// AutocompleteEmail is an email address.
	AutocompleteEmail = Autocomplete{v: "email"}
	// AutocompleteUsername is a sign-in username.
	AutocompleteUsername = Autocomplete{v: "username"}
	// AutocompleteCurrentPassword is the user's existing password, on sign-in forms.
	AutocompleteCurrentPassword = Autocomplete{v: "current-password"}
	// AutocompleteNewPassword is a password being set, on sign-up and reset forms.
	AutocompleteNewPassword = Autocomplete{v: "new-password"}
	// AutocompleteOrganization is a company or organization name.
	AutocompleteOrganization = Autocomplete{v: "organization"}
)

// FieldProps holds a Field's optional settings; every zero value is the default.
type FieldProps struct {
	// Attrs holds htmx and script hooks for the <input>, filtered by the Attrs allow-list. An
	// aria-describedby value is merged after the hint and error ids.
	Attrs templ.Attributes
	// Type is the input type; the zero value is text.
	Type FieldType
	// Autocomplete is the autofill token; the zero value renders none.
	Autocomplete Autocomplete
	// Value is the input's current value.
	Value string
	// Hint is help text shown under the input and linked with aria-describedby.
	Hint string
	// Error is a validation message. Non-empty marks the input invalid, shows the message and
	// links it with aria-describedby.
	Error string
	// Optional appends " (optional)" to the label.
	Optional bool
	// HideLabel keeps the label for screen readers only.
	HideLabel bool
	// Focus renders autofocus, for example on the invalid field of an error response.
	Focus bool
}

// fieldIDs are the element ids one Field renders.
type fieldIDs struct {
	input string
	hint  string
	error string
}

// newFieldIDs derives a Field's ids: <form ID>-<name> inside a Form with an ID, otherwise <name>,
// with -hint and -error suffixes for the hint and error.
func newFieldIDs(ctx context.Context, name string) fieldIDs {
	input := name

	prefix := formID(ctx)
	if prefix != "" {
		input = prefix + "-" + name
	}

	return fieldIDs{input: input, hint: input + "-hint", error: input + "-error"}
}

// fieldLabel returns the label text, marked optional when p.Optional is set.
func fieldLabel(label string, p *FieldProps) string {
	if p.Optional {
		return label + " (optional)"
	}

	return label
}

// fieldClass returns the wrapper's class attribute.
func fieldClass(p *FieldProps) string {
	if p.Error != "" {
		return "yr-field yr-field--invalid"
	}

	return "yr-field"
}

// fieldDescribedBy merges the hint id, the error id and any aria-describedby from Attrs, in that
// order, space-separated.
func fieldDescribedBy(ids fieldIDs, p *FieldProps, fromAttrs any) string {
	parts := make([]string, 0, 3) //nolint:mnd // hint, error, Attrs.
	if p.Hint != "" {
		parts = append(parts, ids.hint)
	}

	if p.Error != "" {
		parts = append(parts, ids.error)
	}

	s, ok := fromAttrs.(string)
	if ok && s != "" {
		parts = append(parts, s)
	}

	return strings.Join(parts, " ")
}

// fieldInputAttrs returns the <input>'s filtered Attrs plus the library-owned aria-invalid and
// merged aria-describedby.
func fieldInputAttrs(ids fieldIDs, p *FieldProps) templ.Attributes {
	attrs := filterAttrs("Field", p.Attrs)

	describedBy := fieldDescribedBy(ids, p, attrs["aria-describedby"])
	delete(attrs, "aria-describedby")

	if describedBy != "" {
		attrs["aria-describedby"] = describedBy
	}

	if p.Error != "" {
		attrs["aria-invalid"] = "true"
	}

	return attrs
}
