package ui

import (
	"bytes"
	"errors"
	"testing"

	"github.com/a-h/templ"
)

// TestSplitRegionOutsideSplitPanics checks that a Split region rendered outside a Split is
// reported, which panics under go test.
func TestSplitRegionOutsideSplitPanics(t *testing.T) {
	t.Parallel()

	err := renderPanic(t, SplitAside(RegionProps{}))
	if !errors.Is(err, errRegionOutsideSplit) {
		t.Errorf("panic = %v, want %v", err, errRegionOutsideSplit)
	}
}

// renderPanic renders c and returns the error it panicked with, or nil when it did not panic.
func renderPanic(t *testing.T, c templ.Component) (err error) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		var ok bool

		err, ok = r.(error)
		if !ok {
			t.Errorf("panic value %v is not an error", r)
		}
	}()

	renderErr := c.Render(t.Context(), &bytes.Buffer{})
	if renderErr != nil {
		t.Errorf("render: %v", renderErr)
	}

	return nil
}
