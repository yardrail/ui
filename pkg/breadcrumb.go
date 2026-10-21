package ui

import "github.com/a-h/templ"

// BreadcrumbProps configures a Breadcrumb.
type BreadcrumbProps struct {
	Attrs templ.Attributes
}

// Breadcrumb renders a breadcrumb navigation. Children are the links, separators, and current item.
func Breadcrumb(p *BreadcrumbProps) templ.Component {
	return breadcrumb(p)
}
