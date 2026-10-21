package main

func attentionCardExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: attentionCardDefault()},
		{Name: exampleDanger, Component: attentionCardDanger()},
	}
}
