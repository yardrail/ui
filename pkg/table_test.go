package ui

import (
	"context"
	"fmt"
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

func TestRenderTable(t *testing.T) {
	t.Parallel()

	for i, hideCaption := range []bool{false, true} {
		html := renderString(t, withChildren(
			Table("Results", TableProps{HideCaption: hideCaption}),
			text("rows"),
		))
		writeClassCoverage(t, "table", fmt.Sprintf("%03d", i), html)
	}
}

func TestRenderTableWithColumns(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Table("Users", TableProps{
			Columns: []Column{Col("Name"), Col("Email"), ActionsCol()},
		}),
		withChildren(Row(RowProps{}),
			withChildren(Cell(CellProps{Kind: CellName}), text("Alice")),
			withChildren(Cell(CellProps{Kind: CellNumeric}), text("42")),
			withChildren(ActionsCell(), text("edit")),
		),
		withChildren(Row(RowProps{Disabled: true}),
			withChildren(Cell(CellProps{}), text("Bob")),
			withChildren(Cell(CellProps{}), text("7")),
			withChildren(ActionsCell(), text("edit")),
		),
	))
	writeClassCoverage(t, "table", "with_columns", html)
}
