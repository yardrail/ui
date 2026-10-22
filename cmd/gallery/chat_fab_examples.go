package main

func chatFabExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: chatFabDefault(), DisplayURL: "platform.yardrail.com/automations"},
		{Name: "variants", Component: chatFabVariants()},
		{Name: "badge", Component: chatFabWithBadge()},
	}
}
