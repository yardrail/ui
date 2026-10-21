package ui

func sectionHeaderExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: sectionHeaderDefaultExample()},
		{Name: "with-action", Component: sectionHeaderWithActionExample()},
	}
}
