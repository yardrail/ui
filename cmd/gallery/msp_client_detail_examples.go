package main

func mspClientOverviewExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			DisplayURL: "platform.yardrail.com/msp/clients/acme-corp",
			Component:  mspClientOverviewPage(),
		},
	}
}

func mspClientUsersExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/users",
			Component:  mspClientUsersPage(),
		},
	}
}

func mspClientConnectorsExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/connectors",
			Component:  mspClientConnectorsPage(),
		},
	}
}

func mspClientSettingsExamples() []Example {
	return []Example{
		{
			Name:       exampleDefault,
			DisplayURL: "platform.yardrail.com/msp/clients/acme-corp/settings",
			Component:  mspClientSettingsPage(),
		},
	}
}
