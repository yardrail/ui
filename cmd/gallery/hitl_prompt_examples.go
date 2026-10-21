package main

func hitlPromptExamples() []Example {
	return []Example{
		{Name: "approve-reject", Component: hitlPromptApproveReject()},
		{Name: exampleDanger, Component: hitlPromptDanger()},
	}
}
