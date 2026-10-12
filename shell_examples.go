package ui

// shellExamples covers the four page kinds a Shell serves.
func shellExamples() []Example {
	return []Example{
		{Name: "app-page", Component: shellAppPage(), DisplayURL: "platform.yardrail.com/clients"},
		{Name: "settings-page", Component: shellSettingsPage(), DisplayURL: ""},
		{Name: "focus-page", Component: shellFocusPage(), DisplayURL: ""},
		{Name: "public-page", Component: shellPublicPage(), DisplayURL: ""},
	}
}
