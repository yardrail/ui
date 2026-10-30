package ui

import "testing"

func TestRenderResizeHandle(t *testing.T) {
	t.Parallel()

	html := renderString(t, ResizeHandle(&ResizeHandleProps{Target: "sidebar"}))
	writeClassCoverage(t, "resize_handle", "default", html)
}
