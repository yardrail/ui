package ui

func mspClientsExamples() []Example {
	return []Example{
		{Name: exampleDefault, DisplayURL: "platform.yardrail.com/msp/clients", Component: mspClients()},
	}
}
