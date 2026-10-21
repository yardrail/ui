package main

func drawerExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: drawerExampleDefault()},
		{Name: exampleWithContent, Component: drawerExampleWithContent()},
		{Name: "with-form", Component: drawerExampleWithForm()},
	}
}
