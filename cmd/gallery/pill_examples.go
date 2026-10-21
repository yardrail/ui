package main

func pillExamples() []Example {
	return []Example{
		{Name: "tones", Component: pillTonesExample()},
		{Name: "status", Component: pillStatusExample()},
		{Name: "interactive", Component: pillInteractiveExample()},
	}
}
