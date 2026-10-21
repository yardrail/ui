package ui

// Space is a step of the spacing scale. The zero value carries no token and means the primitive's
// default.
type Space struct{ v string }

// Spacing steps, named as multiples of 0.25rem.
var (
	// Space1 is 0.25rem.
	Space1 = Space{v: "--yr-space-1"}
	// Space2 is 0.5rem.
	Space2 = Space{v: "--yr-space-2"}
	// Space3 is 0.75rem.
	Space3 = Space{v: "--yr-space-3"}
	// Space4 is 1rem.
	Space4 = Space{v: "--yr-space-4"}
	// Space6 is 1.5rem.
	Space6 = Space{v: "--yr-space-6"}
	// Space8 is 2rem.
	Space8 = Space{v: "--yr-space-8"}
	// Space12 is 3rem.
	Space12 = Space{v: "--yr-space-12"}
	// Space16 is 4rem.
	Space16 = Space{v: "--yr-space-16"}
)

// FontSize is a step of the text size scale. The zero value carries no token and means the
// primitive's default.
type FontSize struct{ v string }

// Text sizes.
var (
	// FontSizeXs is 0.75rem, for captions and table headers.
	FontSizeXs = FontSize{v: "--yr-font-size-xs"}
	// FontSizeSm is 0.875rem, for interface text.
	FontSizeSm = FontSize{v: "--yr-font-size-sm"}
	// FontSizeMd is 1rem, for body text and inputs.
	FontSizeMd = FontSize{v: "--yr-font-size-md"}
	// FontSizeLg is 1.125rem, for card titles and lead text.
	FontSizeLg = FontSize{v: "--yr-font-size-lg"}
	// FontSizeXl is 1.5rem, for section headings.
	FontSizeXl = FontSize{v: "--yr-font-size-xl"}
	// FontSize2xl is 2rem, for page titles and stat numbers.
	FontSize2xl = FontSize{v: "--yr-font-size-2xl"}
)

// Radius is a step of the corner radius scale. The zero value carries no token and means the
// primitive's default.
type Radius struct{ v string }

// Corner radii.
var (
	// RadiusSm is 4px.
	RadiusSm = Radius{v: "--yr-radius-sm"}
	// RadiusMd is 8px.
	RadiusMd = Radius{v: "--yr-radius-md"}
	// RadiusLg is 12px.
	RadiusLg = Radius{v: "--yr-radius-lg"}
	// RadiusFull is 999px, a pill or circle.
	RadiusFull = Radius{v: "--yr-radius-full"}
)

// Measure is a content width. The zero value carries no token and means the primitive's default.
type Measure struct{ v string }

// Content widths.
var (
	// MeasureNarrow is 24rem.
	MeasureNarrow = Measure{v: "--yr-measure-narrow"}
	// MeasureProse is 36rem.
	MeasureProse = Measure{v: "--yr-measure-prose"}
	// MeasureWide is 48rem.
	MeasureWide = Measure{v: "--yr-measure-wide"}
)

// SideWidth is the width of a sidebar or side panel. The zero value carries no token and means the
// primitive's default.
type SideWidth struct{ v string }

// Side widths.
var (
	// SideWidthSm is 15rem.
	SideWidthSm = SideWidth{v: "--yr-side-sm"}
	// SideWidthMd is 17.5rem.
	SideWidthMd = SideWidth{v: "--yr-side-md"}
	// SideWidthLg is 20rem.
	SideWidthLg = SideWidth{v: "--yr-side-lg"}
)

// ItemWidth is the minimum width of an item in a grid. The zero value carries no token and means
// the primitive's default.
type ItemWidth struct{ v string }

// Item widths.
var (
	// ItemWidthSm is 15rem.
	ItemWidthSm = ItemWidth{v: "--yr-item-width-sm"}
	// ItemWidthMd is 20rem.
	ItemWidthMd = ItemWidth{v: "--yr-item-width-md"}
)
