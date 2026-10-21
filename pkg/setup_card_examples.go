package ui

func setupCardExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: setupCardDefault()},
		{Name: "form-step", Component: setupCardFormStep()},
		{Name: "selection-step", Component: setupCardSelectionStep()},
		{Name: "confirmation", Component: setupCardConfirmation()},
	}
}
