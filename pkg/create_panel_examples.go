package ui

func createPanelExamples() []Example {
	return []Example{
		{Name: "expanded", Component: createPanelExpanded()},
		{Name: exampleCollapsed, Component: createPanelCollapsed()},
	}
}
