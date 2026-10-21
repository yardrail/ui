package ui

import "github.com/a-h/templ"

// ResizeHandleProps configures a ResizeHandle.
type ResizeHandleProps struct {
	Target string
	Attrs  templ.Attributes
}

// ResizeHandle renders a draggable resize handle between layout regions.
func ResizeHandle(p *ResizeHandleProps) templ.Component {
	return resizeHandle(p)
}
