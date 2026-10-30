package ui

import "testing"

func TestRenderWorkflowCard(t *testing.T) {
	t.Parallel()

	html := renderString(t, WorkflowCard("Invoice Sync", "Last run 2h ago", &WorkflowCardProps{
		Status:      text("<span>Active</span>"),
		Runs:        "127",
		SuccessRate: "98%",
		Action:      text("<button>Trigger</button>"),
	}))
	writeClassCoverage(t, "workflow_card", "default", html)
}
