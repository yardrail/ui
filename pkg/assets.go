package ui

import (
	"embed"
	"io/fs"
)

// dist is the committed build output of src/ (pnpm build) plus the vendored htmx.min.js.
//
//go:embed dist
var dist embed.FS

// Assets holds the library's browser files at its root: ui.js, ui.css, their source maps, and
// htmx.min.js. Apps serve it with http.FileServerFS, mounted at /static/ui/.
var Assets = mustSub(dist, "dist")

// mustSub returns the subtree of fsys rooted at dir, panicking if dir is not a valid path.
func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}

	return sub
}
