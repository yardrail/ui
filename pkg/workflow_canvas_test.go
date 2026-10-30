package ui

import "testing"

func TestRenderWorkflowCanvas(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(WorkflowCanvas(WorkflowCanvasProps{}), text("canvas")))
	writeClassCoverage(t, "workflow_canvas", "default", html)
}
