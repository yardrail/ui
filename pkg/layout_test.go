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

func TestRenderStack(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(Stack(StackProps{}), text("item")))
	writeClassCoverage(t, "stack", "default", html)

	html = renderString(t, withChildren(Stack(StackProps{Gap: Space8}), text("item")))
	writeClassCoverage(t, "stack", "gap", html)

	html = renderString(t, withChildren(Stack(StackProps{Align: AlignCenter}), text("item")))
	writeClassCoverage(t, "stack", "align", html)

	html = renderString(t, withChildren(Stack(StackProps{As: AsList}), text("<li>item</li>")))
	writeClassCoverage(t, "stack", "as-list", html)

	html = renderString(
		t,
		withChildren(Stack(StackProps{Gap: Space4, Align: AlignEnd, As: AsList}), text("<li>item</li>")),
	)
	writeClassCoverage(t, "stack", "all-options", html)
}

func TestRenderCluster(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(Cluster(ClusterProps{}), text("item")))
	writeClassCoverage(t, "cluster", "default", html)

	html = renderString(t, withChildren(Cluster(ClusterProps{Gap: Space6}), text("item")))
	writeClassCoverage(t, "cluster", "gap", html)

	html = renderString(t, withChildren(Cluster(ClusterProps{Justify: JustifyBetween}), text("item")))
	writeClassCoverage(t, "cluster", "justify", html)

	html = renderString(t, withChildren(Cluster(ClusterProps{Align: AlignStretch}), text("item")))
	writeClassCoverage(t, "cluster", "align", html)

	html = renderString(t, withChildren(Cluster(ClusterProps{As: AsList}), text("<li>item</li>")))
	writeClassCoverage(t, "cluster", "as-list", html)

	html = renderString(
		t,
		withChildren(
			Cluster(ClusterProps{Gap: Space2, Justify: JustifyCenter, Align: AlignStart, As: AsList}),
			text("<li>item</li>"),
		),
	)
	writeClassCoverage(t, "cluster", "all-options", html)
}

func TestRenderGrid(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(Grid(GridProps{}), text("item")))
	writeClassCoverage(t, "grid", "default", html)

	html = renderString(t, withChildren(Grid(GridProps{MinItemWidth: ItemWidthSm}), text("item")))
	writeClassCoverage(t, "grid", "item-width", html)

	html = renderString(t, withChildren(Grid(GridProps{Gap: Space8}), text("item")))
	writeClassCoverage(t, "grid", "gap", html)

	html = renderString(t, withChildren(Grid(GridProps{As: AsList}), text("<li>item</li>")))
	writeClassCoverage(t, "grid", "as-list", html)

	html = renderString(
		t,
		withChildren(Grid(GridProps{MinItemWidth: ItemWidthMd, Gap: Space4, As: AsList}), text("<li>item</li>")),
	)
	writeClassCoverage(t, "grid", "all-options", html)
}

func TestRenderCenter(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(Center(CenterProps{}), text("content")))
	writeClassCoverage(t, "center", "default", html)

	html = renderString(t, withChildren(Center(CenterProps{Measure: MeasureNarrow}), text("content")))
	writeClassCoverage(t, "center", "measure-narrow", html)

	html = renderString(t, withChildren(Center(CenterProps{Measure: MeasureWide}), text("content")))
	writeClassCoverage(t, "center", "measure-wide", html)
}

func TestRenderSplit(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Split(SplitProps{}),
		withChildren(SplitAside(RegionProps{}), text("aside")),
		withChildren(SplitMain(RegionProps{}), text("main")),
	))
	writeClassCoverage(t, "split", "default", html)

	html = renderString(t, withChildren(
		Split(SplitProps{SideWidth: SideWidthLg, Gap: Space8}),
		withChildren(SplitAside(RegionProps{}), text("aside")),
		withChildren(SplitMain(RegionProps{}), text("main")),
	))
	writeClassCoverage(t, "split", "side-gap", html)
}

func TestRenderDivider(t *testing.T) {
	t.Parallel()

	html := renderString(t, Divider())
	writeClassCoverage(t, "divider", "default", html)

	// Divider inside a list-mode Stack renders as <li role="separator">
	html = renderString(t, withChildren(Stack(StackProps{As: AsList}), Divider()))
	writeClassCoverage(t, "divider", "in-list", html)
}
