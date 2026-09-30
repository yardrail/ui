// Package ui is Yardrail's typed templ component library. Each component owns
// the look and accessibility of the elements it renders: required positional
// arguments carry what the component cannot be correct without (usually the
// accessible name), a <Name>Props struct carries optional settings whose zero
// values are the defaults, and Props.Attrs carries htmx and script hooks the
// library copies through without interpreting.
//
// Assets embeds the built CSS and JavaScript the components depend on, so an
// app serves the same library version it compiles against.
package ui
