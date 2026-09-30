package ui

import (
	"context"
	"io"
	"testing"

	"github.com/a-h/templ"
)

func TestTableRequiredArgs(t *testing.T) {
	t.Parallel()

	t.Run("empty caption", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, Table("", TableProps{}))
	})

	t.Run("empty summary", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, RowEdit("", RowEditProps{}))
	})

	t.Run("empty column label", func(t *testing.T) {
		t.Parallel()

		assertPanics(t, templ.ComponentFunc(func(_ context.Context, _ io.Writer) error {
			Col("")

			return nil
		}))
	})
}
