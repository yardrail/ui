package ui

import (
	"fmt"
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

func TestRenderField(t *testing.T) {
	t.Parallel()

	i := 0

	for _, ft := range (FieldType{}).All() {
		for _, optional := range []bool{false, true} {
			for _, hideLabel := range []bool{false, true} {
				for _, focus := range []bool{false, true} {
					html := renderString(t, Field("Label", "name", FieldProps{
						Type:      ft,
						Optional:  optional,
						HideLabel: hideLabel,
						Focus:     focus,
					}))
					writeClassCoverage(t, "field", fmt.Sprintf("%03d", i), html)
					i++
				}
			}
		}
	}
}

func TestRenderFieldBranches(t *testing.T) {
	t.Parallel()

	t.Run("with value", func(t *testing.T) {
		t.Parallel()

		html := renderString(t, Field("Label", "name", FieldProps{Value: "hello"}))
		if !strings.Contains(html, `value="hello"`) {
			t.Error("missing value attribute")
		}

		writeClassCoverage(t, "field", "with_value", html)
	})

	t.Run("with hint", func(t *testing.T) {
		t.Parallel()

		html := renderString(t, Field("Label", "name", FieldProps{Hint: "Help text"}))
		if !strings.Contains(html, "Help text") {
			t.Error("missing hint text")
		}

		writeClassCoverage(t, "field", "with_hint", html)
	})

	t.Run("with error", func(t *testing.T) {
		t.Parallel()

		html := renderString(t, Field("Label", "name", FieldProps{Error: "Required"}))
		if !strings.Contains(html, "yr-field--invalid") {
			t.Error("missing invalid class")
		}

		if !strings.Contains(html, "yr-field-error") {
			t.Error("missing error element")
		}

		writeClassCoverage(t, "field", "with_error", html)
	})

	t.Run("with autocomplete", func(t *testing.T) {
		t.Parallel()

		for _, ac := range (Autocomplete{}).All() {
			html := renderString(t, Field("Label", "name", FieldProps{Autocomplete: ac}))
			writeClassCoverage(t, "field", "ac_"+ac.v, html)
		}
	})
}
