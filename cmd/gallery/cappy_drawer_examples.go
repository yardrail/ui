package main

func cappyDrawerExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: cappyDrawerDefault()},
		{Name: "with-tool-calls", Component: cappyDrawerToolCalls()},
		{Name: "escalated", Component: cappyDrawerEscalated()},
	}
}
