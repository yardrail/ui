package ui

import (
	"errors"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

// shellTestTitle is the document title used by the Shell tests.
const shellTestTitle = "Clients"

// shellPage renders a Shell titled shellTestTitle with children as its regions.
func shellPage(children ...templ.Component) templ.Component {
	return withChildren(Shell(shellTestTitle, ShellProps{}), children...)
}

// TestShellReportsMisuse checks that each region placement mistake is reported with its
// sentinel error, which panics under go test.
func TestShellReportsMisuse(t *testing.T) {
	t.Parallel()

	idAttrs := templ.Attributes{"id": "custom"}

	cases := []struct {
		name string
		c    templ.Component
		want error
	}{
		{"top bar outside shell", ShellTopBar(RegionProps{}), errRegionOutsideShell},
		{"nav outside shell", ShellNav(RegionProps{}), errRegionOutsideShell},
		{"main outside shell", ShellMain(ShellMainProps{}), errRegionOutsideShell},
		{"footer outside shell", ShellFooter(RegionProps{}), errRegionOutsideShell},
		{"band outside shell", ShellBand(RegionProps{}), errRegionOutsideShell},
		{"band outside main", shellPage(ShellMain(ShellMainProps{}), ShellBand(RegionProps{})), errBandOutsideMain},
		{
			"duplicate top bar",
			shellPage(ShellTopBar(RegionProps{}), ShellTopBar(RegionProps{}), ShellMain(ShellMainProps{})),
			errDuplicateRegion,
		},
		{
			"duplicate nav",
			shellPage(ShellNav(RegionProps{}), ShellNav(RegionProps{}), ShellMain(ShellMainProps{})),
			errDuplicateRegion,
		},
		{"duplicate main", shellPage(ShellMain(ShellMainProps{}), ShellMain(ShellMainProps{})), errDuplicateRegion},
		{
			"duplicate footer",
			shellPage(ShellMain(ShellMainProps{}), ShellFooter(RegionProps{}), ShellFooter(RegionProps{})),
			errDuplicateRegion,
		},
		{"missing main", shellPage(ShellTopBar(RegionProps{})), errMissingMain},
		{"id on main", shellPage(ShellMain(ShellMainProps{Attrs: idAttrs})), errFixedID},
		{"id on nav", shellPage(ShellNav(RegionProps{Attrs: idAttrs}), ShellMain(ShellMainProps{})), errFixedID},
		{"empty title", withChildren(Shell("", ShellProps{}), ShellMain(ShellMainProps{})), errRequiredArgEmpty},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := renderPanic(t, tc.c)
			if !errors.Is(err, tc.want) {
				t.Errorf("panic = %v, want %v", err, tc.want)
			}
		})
	}
}

// TestShellAllRegionsRender checks that a Shell holding every region, with repeated bands inside
// its main, renders without reporting a mistake.
func TestShellAllRegionsRender(t *testing.T) {
	t.Parallel()

	page := withChildren(
		Shell(shellTestTitle, ShellProps{Lang: "de", Head: text("head"), Attrs: templ.Attributes{"data-page": "x"}}),
		ShellTopBar(RegionProps{}),
		ShellNav(RegionProps{Attrs: templ.Attributes{"aria-label": "Primary"}}),
		withChildren(
			ShellMain(ShellMainProps{Width: ShellWidthFull, Messages: text("saved")}),
			ShellBand(RegionProps{}),
			ShellBand(RegionProps{}),
		),
		ShellFooter(RegionProps{}),
	)

	err := renderPanic(t, page)
	if err != nil {
		t.Errorf("panic = %v, want none", err)
	}
}

// TestShellDocumentOrder checks the order of the document parts a minimal Shell renders.
func TestShellDocumentOrder(t *testing.T) {
	t.Parallel()

	out := renderString(t, shellPage(ShellMain(ShellMainProps{})))

	want := []string{
		`<html lang="en">`,
		`<title>Clients</title>`,
		`href="/static/ui/ui.css"`,
		`src="/static/ui/htmx.min.js"`,
		`src="/static/ui/ui.js"`,
		`href="#yr-shell-main"`,
		`<main id="yr-shell-main" class="yr-shell__main">`,
	}

	rest := out

	for _, s := range want {
		_, after, ok := strings.Cut(rest, s)
		if !ok {
			t.Fatalf("output lacks %q after the previous parts:\n%s", s, out)
		}

		rest = after
	}
}

// TestShellMainWidth checks that ShellWidthFull adds the full-width modifier and the zero width
// adds none.
func TestShellMainWidth(t *testing.T) {
	t.Parallel()

	const fullClass = "yr-shell__main--width-full"

	full := renderString(t, shellPage(ShellMain(ShellMainProps{Width: ShellWidthFull})))
	if !strings.Contains(full, fullClass) {
		t.Errorf("ShellWidthFull output lacks %q:\n%s", fullClass, full)
	}

	standard := renderString(t, shellPage(ShellMain(ShellMainProps{})))
	if strings.Contains(standard, fullClass) {
		t.Errorf("zero width output contains %q:\n%s", fullClass, standard)
	}
}

func TestRenderShell(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Shell("Test Page", ShellProps{}),
		ShellMain(ShellMainProps{}),
	))
	writeClassCoverage(t, "shell", "default", html)

	html = renderString(t, withChildren(
		Shell("German Page", ShellProps{Lang: "de", Head: text("<link rel=\"icon\">")}),
		ShellMain(ShellMainProps{}),
	))
	writeClassCoverage(t, "shell", "lang_and_head", html)

	html = renderString(t, withChildren(
		Shell("Body Attrs", ShellProps{Attrs: templ.Attributes{"data-page": "x"}}),
		ShellMain(ShellMainProps{Attrs: templ.Attributes{"data-main": "y"}}),
	))
	writeClassCoverage(t, "shell", "with_attrs", html)
}

func TestRenderShellWithRegions(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		Shell("Full Shell", ShellProps{}),
		withChildren(ShellTopBar(RegionProps{Attrs: templ.Attributes{"data-bar": "1"}}), text("topbar")),
		withChildren(ShellNav(RegionProps{Attrs: templ.Attributes{"aria-label": "Primary"}}), text("nav")),
		withChildren(ShellMain(ShellMainProps{
			Width:    ShellWidthFull,
			Messages: text("saved"),
		}),
			text("main"),
			withChildren(ShellBand(RegionProps{Attrs: templ.Attributes{"data-band": "1"}}), text("band")),
		),
		withChildren(ShellFooter(RegionProps{Attrs: templ.Attributes{"data-footer": "1"}}), text("footer")),
	))
	writeClassCoverage(t, "shell", "with_regions", html)
}
