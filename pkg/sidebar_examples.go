package ui

func sidebarExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: sidebarExampleDefault()},
		{Name: exampleWithContent, Component: sidebarExampleWithContent()},
		{Name: exampleCollapsed, Component: sidebarExampleCollapsed()},
	}
}
