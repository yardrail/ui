package ui

import (
	"strings"
	"testing"
)

const (
	fieldLabelTeamName = "Team name"
	fieldNameName      = "name"
)

func TestFieldIDs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		component func() string
		want      []string
	}{
		{
			name: "alone",
			component: func() string {
				return renderString(t, Field(fieldLabelTeamName, fieldNameName, FieldProps{}))
			},
			want: []string{`<label for="name">`, `id="name"`},
		},
		{
			name: "inside a form with an ID",
			component: func() string {
				return renderString(t, withChildren(
					Form(FormProps{ID: "rename-team-7"}),
					Field(fieldLabelTeamName, fieldNameName, FieldProps{}),
				))
			},
			want: []string{`<label for="rename-team-7-name">`, `id="rename-team-7-name"`},
		},
		{
			name: "hint and error",
			component: func() string {
				return renderString(t, Field(fieldLabelTeamName, fieldNameName, FieldProps{
					Hint:  "Shown in the sidebar.",
					Error: "Team name is required.",
				}))
			},
			want: []string{
				`aria-describedby="name-hint name-error"`,
				`aria-invalid="true"`,
				`<span id="name-hint">`,
				`<span class="yr-field-error" id="name-error">`,
			},
		},
		{
			name: "aria-describedby from Attrs merges after hint and error",
			component: func() string {
				return renderString(t, Field(fieldLabelTeamName, fieldNameName, FieldProps{
					Hint:  "Shown in the sidebar.",
					Error: "Team name is required.",
					Attrs: map[string]any{"aria-describedby": "naming-rules"},
				}))
			},
			want: []string{`aria-describedby="name-hint name-error naming-rules"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			html := tt.component()
			for _, want := range tt.want {
				if !strings.Contains(html, want) {
					t.Errorf("render is missing %s:\n%s", want, html)
				}
			}
		})
	}
}

func TestFieldRequiredArgs(t *testing.T) {
	t.Parallel()

	t.Run("empty label", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, Field("", fieldNameName, FieldProps{}))
	})

	t.Run("empty name", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, Field(fieldLabelTeamName, "", FieldProps{}))
	})
}
