package main

const dropdownMenuChevronSize = 12

func dropdownMenuExamples() []Example {
	return []Example{
		{Name: "closed", Component: dropdownMenuExampleClosed()},
		{Name: "open", Component: dropdownMenuExampleOpen()},
		{Name: "account-menu", Component: dropdownMenuExampleAccountMenu()},
	}
}
