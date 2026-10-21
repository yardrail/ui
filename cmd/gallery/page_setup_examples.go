package main

func pageSetupExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: pageSetupExampleDefault()},
		{Name: "with-setup-card", Component: pageSetupExampleWithSetupCard()},
	}
}
