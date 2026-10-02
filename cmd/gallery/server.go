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

// indexTemplate renders the sidebar + main content layout with browser-frame preview.
var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Active.Name}} · ui gallery</title>
<link rel="stylesheet" href="/ui.css">
<style>
html, body { margin: 0; height: 100%; }
body { display: flex; background: var(--cream); }

.sidebar {
  width: 200px;
  min-width: 200px;
  background: var(--cream-surface);
  border-right: 1px solid var(--wood);
  padding: 1.5rem 0;
  display: flex;
  flex-direction: column;
  height: 100vh;
  position: sticky;
  top: 0;
}
.sidebar-title {
  font-family: var(--font-display);
  font-size: 1.1rem;
  color: var(--navy);
  padding: 0 1.25rem 1rem;
  margin: 0;
  border-bottom: 1px solid var(--wood);
}
.sidebar-nav { list-style: none; margin: 0; padding: 0.5rem 0; }
.sidebar-nav a {
  display: block;
  padding: 0.5rem 1.25rem;
  font-family: var(--font-ui);
  font-size: 0.85rem;
  color: var(--slate);
  text-decoration: none;
  border-left: 3px solid transparent;
}
.sidebar-nav a:hover { background: var(--cream-hover); color: var(--charcoal); text-decoration: none; }
.sidebar-nav a.active {
  color: var(--navy);
  font-weight: 600;
  border-left-color: var(--navy);
  background: var(--navy-light);
}

.main {
  flex: 1;
  min-width: 0;
  padding: 2rem 3rem;
  overflow-y: auto;
}
.main > h1 {
  color: var(--navy);
  font-size: 1.5rem;
  margin: 0 0 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid var(--wood);
}

.component { margin: 1.5rem 0; }
.component > h3 {
  font-family: var(--font-mono);
  font-size: 0.9rem;
  color: var(--slate);
  margin: 0 0 0.5rem;
}

.example-chips { display: flex; flex-wrap: wrap; gap: 0.4rem; margin-bottom: 0.75rem; }
.example-chip {
  font-family: var(--font-mono);
  font-size: 0.8rem;
  padding: 0.3rem 0.7rem;
  border-radius: 4px;
  border: 1px solid var(--wood);
  background: var(--cream-surface);
  color: var(--slate);
  cursor: pointer;
  text-decoration: none;
  transition: all 0.1s;
}
.example-chip:hover { background: var(--cream-hover); color: var(--charcoal); text-decoration: none; }
.example-chip.active {
  background: var(--navy);
  color: #fff;
  border-color: var(--navy);
}

.browser-frame {
  border: 1px solid var(--wood);
  border-radius: 8px;
  overflow: hidden;
  background: var(--cream-surface);
  box-shadow: 0 2px 8px rgba(0,0,0,0.06);
}
.browser-chrome {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  background: #e8e4dd;
  border-bottom: 1px solid var(--wood);
}
.browser-dots { display: flex; gap: 5px; }
.browser-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}
.browser-dot.red { background: #ec6a5e; }
.browser-dot.yellow { background: #f4bf4f; }
.browser-dot.green { background: #61c554; }
.browser-url {
  flex: 1;
  font-family: var(--font-mono);
  font-size: 0.75rem;
  color: var(--slate);
  background: var(--cream-surface);
  border: 1px solid var(--wood);
  border-radius: 4px;
  padding: 0.25rem 0.6rem;
  margin-left: 0.5rem;
}
.browser-body { background: #fff; }
.browser-body iframe {
  display: block;
  width: 100%;
  height: 300px;
  border: none;
}

.no-examples {
  color: var(--slate-disabled);
  font-style: italic;
  font-size: 0.85rem;
  padding: 0.75rem 1rem;
  background: var(--cream-surface);
  border: 1px dashed var(--wood);
  border-radius: 6px;
}
</style>
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

// newHandler returns the gallery handler, serving ui.css from assets as /ui.css.
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
