package ui

const navSearchIconSize = 14

func navbarExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: navbarExampleDefault()},
		{Name: exampleWithContent, Component: navbarExampleWithContent()},
		{Name: "standard-unauthed", Component: navbarExampleStandardUnauthed()},
		{Name: "standard-authed", Component: navbarExampleStandardAuthed()},
	}
}
