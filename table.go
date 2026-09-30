package ui

import "github.com/a-h/templ"

// visuallyHiddenClass hides an element visually while keeping it in the accessibility tree.
const visuallyHiddenClass = "yr-visually-hidden"

// Column is one header cell of a Table. Build it with Col or ActionsCol.
type Column struct {
	label   string
	actions bool
}

// Col returns a column headed label.
func Col(label string) Column {
	requireArg("Col", "label", label)

	return Column{label: label, actions: false}
}

// ActionsCol returns the header of a column of ActionsCell cells: "Actions" for screen readers
// only.
func ActionsCol() Column {
	return Column{label: "Actions", actions: true}
}

// TableProps holds a Table's optional settings; every zero value is the default.
type TableProps struct {
	// Columns are the header cells, in order; none renders no <thead>.
	Columns []Column
	// HideCaption keeps the caption for screen readers only.
	HideCaption bool
}

// RowProps holds a Row's optional settings; every zero value is the default.
type RowProps struct {
	// Attrs holds htmx and script hooks, filtered by the Attrs allow-list. A row that htmx swaps
	// needs an id here.
	Attrs templ.Attributes
	// Disabled dims the row's cells.
	Disabled bool
}

// CellKind is the style of a Cell. The zero value is unstyled.
type CellKind struct{ class string }

// Cell kinds beyond the zero value.
var (
	// CellName emphasises the cell that names the row.
	CellName = CellKind{class: "yr-table-name"}
	// CellMono sets the cell in the monospace font, for ids and keys.
	CellMono = CellKind{class: "yr-table-mono"}
)

// CellProps holds a Cell's optional settings; every zero value is the default.
type CellProps struct {
	// Kind is the cell style; the zero value is unstyled.
	Kind CellKind
}
