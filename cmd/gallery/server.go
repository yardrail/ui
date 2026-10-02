package main

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"

	"github.com/yardrail/ui"
)

// categories defines how components are grouped on the index page.
var categories = []category{
	{Slug: "primitives", Name: "Primitives", Components: []string{
		"button", "pill", "avatar", "icon-badge", "stat", "brand", "rating",
	}},
	{Slug: "forms", Name: "Forms", Components: []string{
		"field", "form", "search", "copy-field",
	}},
	{Slug: "navigation", Name: "Navigation", Components: []string{
		"tabs", "breadcrumb", "nav-shell", "account-menu", "dropdown-menu",
	}},
	{Slug: "data", Name: "Data", Components: []string{
		"table", "row-edit", "list-row", "activity-row", "meta-row",
	}},
	{Slug: "cards", Name: "Cards", Components: []string{
		"card", "action-card", "auth-card", "review-card", "template-card", "create-panel", "settings-section",
	}},
	{Slug: "feedback", Name: "Feedback", Components: []string{
		"alert", "empty-state",
	}},
	{Slug: "misc", Name: "Misc", Components: []string{
		"author-link", "connector-pill", "live-duration",
	}},
}

type category struct {
	Slug       string
	Name       string
	Components []string
}

// indexCategory is the template data for one category section.
type indexCategory struct {
	Slug       string
	Name       string
	Components []indexComponent
}

// indexComponent is the template data for one component within a category.
type indexComponent struct {
	Name     string
	Examples []indexExample
}

// indexExample is one named example of a component.
type indexExample struct {
	Name string
}

// indexPage is the full template data for the gallery index.
type indexPage struct {
	Categories []indexCategory
	Active     indexCategory
	ActiveSlug string
}

// galleryCSS holds gallery.css, the styles of the index page's own chrome (sidebar, chips,
// browser frame). The examples themselves are styled by the library's ui.css only.
//
//go:embed gallery.css
var galleryCSS embed.FS

// indexTemplate renders the sidebar + main content layout with browser-frame preview.
var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Active.Name}} · ui gallery</title>
<link rel="stylesheet" href="/ui.css">
<link rel="stylesheet" href="/gallery.css">
</head>
<body>
<aside class="sidebar">
  <h1 class="sidebar-title">ui gallery</h1>
  <ul class="sidebar-nav">
  {{range .Categories}}<li><a href="/?cat={{.Slug}}"
    {{- if eq $.ActiveSlug .Slug}} class="active"{{end}}>{{.Name}}</a></li>
  {{end}}</ul>
</aside>
<main class="main">
  <h1>{{.Active.Name}}</h1>
  {{range .Active.Components}}<div class="component">
  <h3>{{.Name}}</h3>
  {{if .Examples}}{{if gt (len .Examples) 1}}<div class="example-chips">
    {{$comp := .Name}}{{range $i, $ex := .Examples}}<a class="example-chip{{if eq $i 0}} active{{end}}"
      data-component="{{$comp}}" data-example="{{$ex.Name}}"
      href="/{{$comp}}/{{$ex.Name}}">{{$ex.Name}}</a>
    {{end}}
  </div>
  {{end}}<div class="browser-frame">
    <div class="browser-chrome">
      <div class="browser-dots">
        <span class="browser-dot red"></span>
        <span class="browser-dot yellow"></span>
        <span class="browser-dot green"></span>
      </div>
      <span class="browser-url">/{{.Name}}/{{(index .Examples 0).Name}}</span>
    </div>
    <div class="browser-body">
      <iframe src="/{{.Name}}/{{(index .Examples 0).Name}}"></iframe>
    </div>
  </div>
  {{else}}<div class="no-examples">no examples yet</div>
  {{end}}</div>
  {{end}}
</main>
<script>
document.addEventListener('click', function(e) {
  var chip = e.target.closest('.example-chip');
  if (!chip) return;
  e.preventDefault();
  var component = chip.dataset.component;
  var example = chip.dataset.example;
  var frame = chip.closest('.component').querySelector('.browser-frame');
  var url = '/' + component + '/' + example;
  chip.closest('.example-chips').querySelectorAll('.example-chip').forEach(function(c) {
    c.classList.remove('active');
  });
  chip.classList.add('active');
  frame.querySelector('.browser-url').textContent = url;
  frame.querySelector('iframe').src = url;
});
</script>
</body>
</html>
`))

// exampleTemplate wraps one rendered example in a page that loads the library CSS.
var exampleTemplate = template.Must(template.New("example").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Component}}/{{.Name}} · ui gallery</title>
<link rel="stylesheet" href="/ui.css">
<script type="module" src="/ui.js"></script>
</head>
<body>
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
	examples   map[string]templ.Component
	categories []indexCategory
	assets     fs.FS
}

// newHandler returns the gallery handler, serving ui.css from assets as /ui.css and the
// embedded gallery.css as /gallery.css.
func newHandler(assets fs.FS) http.Handler {
	groups := ui.ExampleGroups()

	// Index examples by component name for fast lookup.
	byComponent := make(map[string][]ui.Example, len(groups))
	for _, g := range groups {
		byComponent[g.Component] = g.Examples
	}

	g := &gallery{
		examples:   make(map[string]templ.Component),
		categories: make([]indexCategory, len(categories)),
		assets:     assets,
	}

	for i, cat := range categories {
		ic := indexCategory{Slug: cat.Slug, Name: cat.Name, Components: make([]indexComponent, len(cat.Components))}

		for j, comp := range cat.Components {
			ic.Components[j] = indexComponent{Name: comp, Examples: nil}

			if exs, ok := byComponent[comp]; ok {
				for _, ex := range exs {
					g.examples[comp+"/"+ex.Name] = ex.Component
					ic.Components[j].Examples = append(ic.Components[j].Examples, indexExample{Name: ex.Name})
				}
			}
		}

		g.categories[i] = ic
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", g.index)
	mux.HandleFunc("GET /ui.css", g.css)
	mux.HandleFunc("GET /gallery.css", g.galleryStyles)
	mux.HandleFunc("GET /ui.js", g.js)
	mux.HandleFunc("GET /{component}/{name}", g.example)

	return mux
}

// index renders the sidebar and the selected category's component list.
func (g *gallery) index(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("cat")

	var active *indexCategory

	for i := range g.categories {
		if g.categories[i].Slug == slug {
			active = &g.categories[i]

			break
		}
	}

	if active == nil {
		active = &g.categories[0]
		slug = active.Slug
	}

	g.write(w, indexTemplate, indexPage{
		Categories: g.categories,
		Active:     *active,
		ActiveSlug: slug,
	})
}

// css serves the library CSS.
func (g *gallery) css(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, g.assets, "ui.css")
}

// galleryStyles serves the index page's own stylesheet.
func (*gallery) galleryStyles(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, galleryCSS, "gallery.css")
}

// js serves the library JS bundle.
func (g *gallery) js(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, g.assets, "ui.js")
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
