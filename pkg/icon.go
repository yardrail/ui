package ui

import (
	"io/fs"

	"github.com/a-h/templ"
)

// Icon returns the named Lucide icon as a templ component. The name is the
// icon's file stem (e.g. "pencil", "mail", "check-circle"). Panics if the
// icon does not exist in dist/icons/.
func Icon(name string) templ.Component {
	svg, err := fs.ReadFile(Assets, "icons/"+name+".svg")
	if err != nil {
		panic("ui.Icon: unknown icon " + name + ": " + err.Error())
	}

	return templ.Raw(string(svg))
}
