package ui

func workflowCardExamples() []Example {
	return []Example{
		{Name: "active", Component: workflowCardActive()},
		{Name: "paused", Component: workflowCardPaused()},
	}
}
