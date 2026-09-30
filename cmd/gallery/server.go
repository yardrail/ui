package main

import (
	"bytes"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"

	"github.com/yardrail/ui"
)

// indexTemplate lists every example, grouped by component.
var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>ui gallery</title>
<link rel="stylesheet" href="/ui.css">
</head>
<body>
<h1>ui gallery</h1>
{{range .}}<section>
<h2>{{.Component}}</h2>
<ul>
{{$c := .Component}}{{range .Examples}}<li><a href="/{{$c}}/{{.Name}}">{{.Name}}</a></li>
{{end}}</ul>
</section>
{{end}}</body>
</html>
`))

// exampleTemplate wraps one rendered example in a page that loads the library CSS.
var exampleTemplate = template.Must(template.New("example").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Component}}/{{.Name}} · ui gallery</title>
<link rel="stylesheet" href="/ui.css">
</head>
<body>
<p><a href="/">ui gallery</a> / {{.Component}} / {{.Name}}</p>
<main>{{.HTML}}</main>
</body>
</html>
`))

// examplePage is the data exampleTemplate renders.
type examplePage struct {
	Component string
	Name      string
	HTML      template.HTML
}

// gallery serves the ui examples and the library CSS.
type gallery struct {
	examples map[string]templ.Component
	groups   []ui.ExampleGroup
	assets   fs.FS
}

// newHandler returns the gallery handler, serving ui.css from assets as /ui.css.
func newHandler(assets fs.FS) http.Handler {
	g := &gallery{
		examples: make(map[string]templ.Component),
		groups:   ui.ExampleGroups(),
		assets:   assets,
	}

	for _, group := range g.groups {
		for _, ex := range group.Examples {
			g.examples[group.Component+"/"+ex.Name] = ex.Component
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", g.index)
	mux.HandleFunc("GET /ui.css", g.css)
	mux.HandleFunc("GET /{component}/{name}", g.example)

	return mux
}

// index lists every example.
func (g *gallery) index(w http.ResponseWriter, _ *http.Request) {
	g.write(w, indexTemplate, g.groups)
}

// css serves the library CSS.
func (g *gallery) css(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, g.assets, "ui.css")
}

// example renders one example inside a page, or 404s for an unknown component or name.
func (g *gallery) example(w http.ResponseWriter, r *http.Request) {
	component, name := r.PathValue("component"), r.PathValue("name")

	c, ok := g.examples[component+"/"+name]
	if !ok {
		http.NotFound(w, r)

		return
	}

	var buf bytes.Buffer

	err := c.Render(r.Context(), &buf)
	if err != nil {
		slog.Error("render example", "component", component, "name", name, "err", err)
		http.Error(w, "render failed: "+err.Error(), http.StatusInternalServerError)

		return
	}

	//nolint:gosec // The markup is the ui package's own escaped render output.
	g.write(w, exampleTemplate, examplePage{Component: component, Name: name, HTML: template.HTML(buf.String())})
}

// write executes t with data into a buffer, so a template error becomes a 500 instead of a
// truncated page.
func (*gallery) write(w http.ResponseWriter, t *template.Template, data any) {
	var buf bytes.Buffer

	err := t.Execute(&buf, data)
	if err != nil {
		slog.Error("execute gallery template", "template", t.Name(), "err", err)
		http.Error(w, "template failed", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = w.Write(buf.Bytes())
	if err != nil {
		slog.Warn("write gallery response", "err", err)
	}
}
