package main

func authCardExamples() []Example {
	return []Example{
		{Name: "sign-in", Component: authCardSignIn()},
		{Name: "with-error", Component: authCardWithError()},
	}
}
