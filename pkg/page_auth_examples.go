package ui

func pageAuthExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: pageAuthExampleDefault()},
		{Name: "with-auth-card", Component: pageAuthExampleWithAuthCard()},
	}
}
