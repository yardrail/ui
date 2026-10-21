package main

func workspaceExamples() []Example {
	return []Example{
		{Name: "expanded", Component: workspaceExampleExpanded()},
		{Name: "with-drawer", Component: workspaceExampleWithDrawer()},
		{Name: exampleCollapsed, Component: workspaceExampleCollapsed()},
	}
}
