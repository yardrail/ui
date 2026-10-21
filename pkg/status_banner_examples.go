package ui

func statusBannerExamples() []Example {
	return []Example{
		{Name: "active", Component: statusBannerActiveExample()},
		{Name: "paused", Component: statusBannerPausedExample()},
		{Name: "failed", Component: statusBannerFailedExample()},
	}
}
