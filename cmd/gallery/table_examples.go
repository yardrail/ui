package main

import (
	"github.com/a-h/templ"
	ui "github.com/yardrail/ui/pkg"
)

// tableExamples covers the caption states, an actions column, and a Row rendered on its own in
// each of its states, as an htmx swap response would render it.
func tableExamples() []Example {
	return []Example{
		{Name: "caption", Component: withChildren(
			ui.Table("Teams", ui.TableProps{Columns: []ui.Column{ui.Col("Name"), ui.Col("Members")}}),
			exampleTeamRow("Platform", "12"),
			exampleTeamRow("Design", "4"),
		), DisplayURL: ""},
		{Name: "hidden-caption", Component: withChildren(
			ui.Table(
				"Teams",
				ui.TableProps{Columns: []ui.Column{ui.Col("Name"), ui.Col("Members")}, HideCaption: true},
			),
			exampleTeamRow("Platform", "12"),
		), DisplayURL: ""},
		{Name: "actions", Component: withChildren(
			ui.Table("Teams", ui.TableProps{Columns: []ui.Column{ui.Col("Name"), ui.ActionsCol()}}),
			withChildren(
				ui.Row(ui.RowProps{Attrs: templ.Attributes{"id": "team-7"}}),
				withChildren(ui.Cell(ui.CellProps{Kind: ui.CellName}), text("Platform")),
				withChildren(ui.ActionsCell(), ui.Button("Remove", ui.ButtonProps{Variant: ui.ButtonDangerQuiet})),
			),
		), DisplayURL: ""},
		{Name: "row-plain", Component: exampleTeamRow("Platform", "12"), DisplayURL: ""},
		{Name: "row-disabled", Component: withChildren(
			ui.Row(ui.RowProps{Disabled: true}),
			withChildren(ui.Cell(ui.CellProps{}), text("Archived")),
		), DisplayURL: ""},
		{Name: "cell-name", Component: withChildren(
			ui.Row(ui.RowProps{}),
			withChildren(ui.Cell(ui.CellProps{Kind: ui.CellName}), text("Platform")),
		), DisplayURL: ""},
		{Name: "cell-numeric", Component: withChildren(
			ui.Row(ui.RowProps{}),
			withChildren(ui.Cell(ui.CellProps{Kind: ui.CellNumeric}), text("1m 42s")),
		), DisplayURL: ""},
	}
}

// exampleTeamRow is a plain two-cell row.
func exampleTeamRow(name, members string) templ.Component {
	return withChildren(
		ui.Row(ui.RowProps{}),
		withChildren(ui.Cell(ui.CellProps{}), text(name)),
		withChildren(ui.Cell(ui.CellProps{}), text(members)),
	)
}
