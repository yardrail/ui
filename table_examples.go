package ui

import "github.com/a-h/templ"

// tableExamples covers the caption states, an actions column, and a Row rendered on its own in
// each of its states, as an htmx swap response would render it.
func tableExamples() []Example {
	return []Example{
		{Name: "caption", Component: withChildren(
			Table("Teams", TableProps{Columns: []Column{Col("Name"), Col("Members")}}),
			exampleTeamRow("Platform", "12"),
			exampleTeamRow("Design", "4"),
		)},
		{Name: "hidden-caption", Component: withChildren(
			Table("Teams", TableProps{Columns: []Column{Col("Name"), Col("Members")}, HideCaption: true}),
			exampleTeamRow("Platform", "12"),
		)},
		{Name: "actions", Component: withChildren(
			Table("Teams", TableProps{Columns: []Column{Col("Name"), ActionsCol()}}),
			withChildren(
				Row(RowProps{Attrs: templ.Attributes{"id": "team-7"}}),
				withChildren(Cell(CellProps{Kind: CellName}), text("Platform")),
				withChildren(ActionsCell(), Button("Remove", ButtonProps{Variant: DangerQuiet})),
			),
		)},
		{Name: "row-plain", Component: exampleTeamRow("Platform", "12")},
		{Name: "row-disabled", Component: withChildren(
			Row(RowProps{Disabled: true}),
			withChildren(Cell(CellProps{}), text("Archived")),
		)},
		{Name: "cell-name", Component: withChildren(
			Row(RowProps{}),
			withChildren(Cell(CellProps{Kind: CellName}), text("Platform")),
		)},
		{Name: "cell-mono", Component: withChildren(
			Row(RowProps{}),
			withChildren(Cell(CellProps{Kind: CellMono}), text("whk_7f3a9c")),
		)},
	}
}

// exampleTeamRow is a plain two-cell row.
func exampleTeamRow(name, members string) templ.Component {
	return withChildren(
		Row(RowProps{}),
		withChildren(Cell(CellProps{}), text(name)),
		withChildren(Cell(CellProps{}), text(members)),
	)
}
