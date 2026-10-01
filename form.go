package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

// FormMethod is the HTTP method of a Form. The zero value renders method="post".
type FormMethod struct{ v string }

// Form methods.
var (
	// MethodPost renders method="post", the same as the zero value.
	MethodPost = FormMethod{v: "post"}
	// MethodGet renders method="get".
	MethodGet = FormMethod{v: "get"}
)

// String returns the value of the method attribute.
func (m FormMethod) String() string {
	if m.v == "" {
		return MethodPost.v
	}

	return m.v
}

// FormLayout is how a Form lays out its fields. The zero value stacks them.
type FormLayout struct{ class string }

// FormStacked stacks fields vertically with consistent spacing, the default for sign-up and
// settings forms.
var FormStacked = FormLayout{class: "yr-form-stacked"}

// FormInline lays fields and buttons out on one line, for editing a table row in place.
var FormInline = FormLayout{class: "yr-row-edit-form"}

// layoutClass returns the CSS class for the form layout, defaulting to stacked.
func (p *FormProps) layoutClass() string {
	if p.Layout.class != "" {
		return p.Layout.class
	}

	return FormStacked.class
}

// FormProps holds a Form's optional settings; every zero value is the default.
type FormProps struct {
	// Attrs holds htmx and script hooks, filtered by the Attrs allow-list.
	Attrs templ.Attributes
	// Method is the HTTP method; the zero value is post.
	Method FormMethod
	// Layout is the field layout; the zero value stacks fields.
	Layout FormLayout
	// Action is the URL the form submits to; empty submits to the current URL.
	Action string
	// ID is the form's id and the id prefix of every Field inside it. Forms repeated on one page,
	// such as one per table row, must set it.
	ID string
}

// formIDKey is the context key under which a Form passes its ID to the Fields inside it.
type formIDKey struct{}

// Form renders a <form novalidate> around its children. Server-side validation reports errors
// through Field's Error prop, so the browser's own validation bubbles are off.
func Form(p FormProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if p.ID != "" {
			ctx = context.WithValue(ctx, formIDKey{}, p.ID)
		}

		return formElement(p).Render(ctx, w)
	})
}

// formID returns the ID of the enclosing Form, or "" outside a Form or inside one without an ID.
func formID(ctx context.Context) string {
	id, ok := ctx.Value(formIDKey{}).(string)
	if !ok {
		return ""
	}

	return id
}
