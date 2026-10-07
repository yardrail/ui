package ui

import (
	"context"
	"io"
	"strings"

	"github.com/a-h/templ"
)

// Align is how a layout primitive aligns its children across its main axis. The zero value means
// the primitive's default.
type Align struct{ v string }

// Cross-axis alignments.
var (
	// AlignStart packs children against the start edge.
	AlignStart = Align{v: "start"}
	// AlignCenter centres children.
	AlignCenter = Align{v: "center"}
	// AlignEnd packs children against the end edge.
	AlignEnd = Align{v: "end"}
	// AlignStretch stretches children to fill the cross axis.
	AlignStretch = Align{v: "stretch"}
)

// Justify is how a layout primitive distributes its children along its main axis. The zero value
// means the primitive's default.
type Justify struct{ v string }

// Main-axis distributions.
var (
	// JustifyStart packs children against the start edge.
	JustifyStart = Justify{v: "start"}
	// JustifyCenter centres children.
	JustifyCenter = Justify{v: "center"}
	// JustifyEnd packs children against the end edge.
	JustifyEnd = Justify{v: "end"}
	// JustifyBetween spreads children so the first and last touch the edges.
	JustifyBetween = Justify{v: "between"}
)

// As is the element a layout primitive renders as. The zero value renders a <div>.
type As struct{ v string }

// AsList renders the primitive as <ul role="list">, with list styling reset. Callers supply the
// <li> children.
var AsList = As{v: "list"}

// list reports whether the primitive renders as a list.
func (a As) list() bool {
	return a == AsList
}

// Block classes of the layout primitives.
const (
	stackBlock   = "yr-stack"
	clusterBlock = "yr-cluster"
)

// Modifier names shared by several primitives.
const (
	modGap   = "gap"
	modAlign = "align"
)

// step returns the scale step of s, such as "4" for Space4, or "" for the zero value.
func (s Space) step() string {
	return strings.TrimPrefix(s.v, "--yr-space-")
}

// modifier returns the class block--name-value, or "" when value is empty.
func modifier(block, name, value string) string {
	if value == "" {
		return ""
	}

	return block + "--" + name + "-" + value
}

// classes joins block and every non-empty modifier into a class attribute.
func classes(block string, modifiers ...string) string {
	parts := []string{block}

	for _, m := range modifiers {
		if m != "" {
			parts = append(parts, m)
		}
	}

	return strings.Join(parts, " ")
}

// layoutInfo describes the layout primitive enclosing a component.
type layoutInfo struct {
	// block is the enclosing primitive's block class, or "" outside any primitive.
	block string
	// list reports whether the enclosing primitive rendered as a list.
	list bool
}

// layoutParentKey is the context key under which a layout primitive passes its layoutInfo to its
// children.
type layoutParentKey struct{}

// withLayoutParent returns ctx carrying info as the enclosing layout primitive.
func withLayoutParent(ctx context.Context, info layoutInfo) context.Context {
	return context.WithValue(ctx, layoutParentKey{}, info)
}

// asParent renders el with info as the enclosing layout primitive of its children.
func asParent(info layoutInfo, el templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return el.Render(withLayoutParent(ctx, info), w)
	})
}

// StackProps holds a Stack's optional settings; every zero value is the default.
type StackProps struct {
	// Attrs holds htmx and script hooks, filtered by the Attrs allow-list.
	Attrs templ.Attributes
	// Gap is the space between children; the zero value is Space4.
	Gap Space
	// Align is the horizontal alignment of children; the zero value stretches them.
	Align Align
	// As is the rendered element; the zero value is a <div>.
	As As
}

// class returns the class attribute of a Stack.
func (p *StackProps) class() string {
	return classes(stackBlock, modifier(stackBlock, modGap, p.Gap.step()), modifier(stackBlock, modAlign, p.Align.v))
}

// Stack lays its children out in a column with a consistent gap between them.
func Stack(p StackProps) templ.Component {
	return asParent(layoutInfo{block: stackBlock, list: p.As.list()},
		layoutBox("Stack", p.class(), p.As, p.Attrs))
}

// ClusterProps holds a Cluster's optional settings; every zero value is the default.
type ClusterProps struct {
	// Attrs holds htmx and script hooks, filtered by the Attrs allow-list.
	Attrs templ.Attributes
	// Gap is the space between children, in both directions; the zero value is Space2.
	Gap Space
	// Justify is the horizontal distribution of children; the zero value packs them at the start.
	Justify Justify
	// Align is the vertical alignment of children; the zero value centres them.
	Align Align
	// As is the rendered element; the zero value is a <div>.
	As As
}

// class returns the class attribute of a Cluster.
func (p *ClusterProps) class() string {
	return classes(clusterBlock,
		modifier(clusterBlock, modGap, p.Gap.step()),
		modifier(clusterBlock, "justify", p.Justify.v),
		modifier(clusterBlock, modAlign, p.Align.v))
}

// Cluster lays its children out in a row that wraps onto further rows when it runs out of space.
func Cluster(p ClusterProps) templ.Component {
	return asParent(layoutInfo{block: clusterBlock, list: p.As.list()},
		layoutBox("Cluster", p.class(), p.As, p.Attrs))
}
