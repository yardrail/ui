package main

func cardExamples() []Example {
	return []Example{
		{Name: "with-title", Component: cardWithTitle()},
		{Name: "header-footer", Component: cardHeaderFooter()},
		{Name: "interactive", Component: cardInteractive()},
	}
}
