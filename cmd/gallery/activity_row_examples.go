package main

func activityRowExamples() []Example {
	return []Example{
		{Name: exampleRows, Component: activityRowExampleRows()},
		{Name: "journal", Component: activityRowExampleJournal()},
	}
}
