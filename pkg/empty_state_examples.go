package ui

// emptyStateIconSize is the icon size the apps use in an empty state.
const emptyStateIconSize = 36

func emptyStateExamples() []Example {
	return []Example{
		{Name: "text", Component: emptyStateTextExample()},
		{Name: "with-icon", Component: emptyStateWithIconExample()},
	}
}
