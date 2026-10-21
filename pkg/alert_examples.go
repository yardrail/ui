package ui

func alertExamples() []Example {
	return []Example{
		{Name: "all-tones", Component: alertTonesExample()},
		{Name: "with-icon", Component: alertWithIconExample()},
	}
}
