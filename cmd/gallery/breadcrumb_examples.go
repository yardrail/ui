package main

func breadcrumbExamples() []Example {
	return []Example{
		{Name: "trail", Component: breadcrumbExampleTrail()},
		{Name: "back-link", Component: breadcrumbExampleBackLink()},
	}
}
