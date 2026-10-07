package main

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path"

	"github.com/a-h/templ"

	"github.com/yardrail/ui"
)

// Region component slugs that double as the name of the shell slot each fills.
const (
	sidebar = "sidebar"
	drawer  = "drawer"
)

// sections defines the two-tier grouping shown in the gallery sidebar.
var sections = []section{
	{Name: "UI Mocks", Categories: []category{
		{Slug: "msp-portal", Name: "MSP Portal", ContextSlot: nil, Components: []string{
			"msp-dashboard",
			"msp-clients",
			"msp-client-overview",
			"msp-client-users",
			"msp-client-connectors",
			"msp-client-settings",
		}},
	}},
	{Name: "Structure", Categories: []category{
		{Slug: "page-shell", Name: "Page Shell", ContextSlot: nil, Components: []string{
			"page-workspace",
			"page-auth",
			"page-setup",
		}},
		{Slug: "layout", Name: "Layout", ContextSlot: nil, Components: []string{
			"stack", "cluster", "grid", "split", "center",
		}},
		{Slug: "regions", Name: "Regions", ContextSlot: map[string]string{
			"navbar": "nav", sidebar: sidebar, drawer: drawer,
		}, Components: []string{
			"navbar", sidebar, drawer,
		}},
	}},
	{Name: "Components", Categories: []category{
		{Slug: "primitives", Name: "Primitives", ContextSlot: nil, Components: []string{
			"button", "pill", "avatar", "icon-badge", "stat", "brand", "rating",
		}},
		{Slug: "forms", Name: "Forms", ContextSlot: nil, Components: []string{
			"field", "form", "search", "copy-field",
		}},
		{Slug: "navigation", Name: "Navigation", ContextSlot: nil, Components: []string{
			"tabs", "breadcrumb", "nav-shell", "account-menu", "dropdown-menu",
		}},
		{Slug: "data", Name: "Data", ContextSlot: nil, Components: []string{
			"table", "row-edit", "list-row", "activity-row", "meta-row",
		}},
		{Slug: "cards", Name: "Cards", ContextSlot: nil, Components: []string{
			"card",
			"action-card",
			"auth-card",
			"setup-card",
			"review-card",
			"template-card",
			"create-panel",
			"settings-section",
		}},
		{Slug: "feedback", Name: "Feedback", ContextSlot: nil, Components: []string{
			"alert", "empty-state",
		}},
		{Slug: "misc", Name: "Misc", ContextSlot: nil, Components: []string{
			"author-link", "connector-pill", "live-duration",
		}},
	}},
}

type section struct {
	Name       string
	Categories []category
}

type category struct {
	Slug        string
	Name        string
	ContextSlot map[string]string // component name → shell slot ("nav","sidebar","drawer")
	Components  []string
}

// indexSection is the template data for one top-level sidebar group.
type indexSection struct {
	Name       string
	Categories []indexCategory
}

// indexCategory is the template data for one category section.
type indexCategory struct {
	Slug       string
	Name       string
	Components []indexComponent
}

// indexComponent is the template data for one component within a category.
type indexComponent struct {
	Name       string
	HasContext bool
	Examples   []indexExample
}

// indexExample is one named example of a component.
type indexExample struct {
	Name       string
	DisplayURL string
}

// indexPage is the full template data for the gallery index.
type indexPage struct {
	Sections   []indexSection
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
  {{range .Sections}}<li class="sidebar-section-label">{{.Name}}</li>
  {{range .Categories}}<li><a href="/?cat={{.Slug}}"
    {{- if eq $.ActiveSlug .Slug}} class="active"{{end}}>{{.Name}}</a></li>
  {{end}}{{end}}</ul>
</aside>
<main class="main">
  <h1>{{.Active.Name}}</h1>
  {{range .Active.Components}}<details class="component" open{{if .HasContext}} data-has-context{{end}}>
  <summary><h3>{{.Name}}</h3></summary>
  {{if .Examples}}<div class="example-toolbar">
  {{if gt (len .Examples) 1}}<div class="example-chips">
    {{$comp := .Name}}{{range $i, $ex := .Examples}}<a class="example-chip{{if eq $i 0}} active{{end}}"
      data-component="{{$comp}}" data-example="{{$ex.Name}}"{{if $ex.DisplayURL}}
      data-display-url="{{$ex.DisplayURL}}"{{end}}
      href="/{{$comp}}/{{$ex.Name}}">{{$ex.Name}}</a>
    {{end}}
  </div>{{end}}
  {{if .HasContext}}<div class="view-toggle">
    <button class="view-toggle-btn" data-view="isolated">Isolated</button>
    <button class="view-toggle-btn active" data-view="in-shell">In Shell</button>
  </div>{{end}}
  </div>
  <div class="browser-frame">
    <div class="browser-chrome">
      <div class="browser-dots">
        <span class="browser-dot red"></span>
        <span class="browser-dot yellow"></span>
        <span class="browser-dot green"></span>
      </div>
      <span class="browser-url">{{if (index .Examples 0).DisplayURL}}{{(index .Examples 0).DisplayURL}}
        {{- else}}/{{.Name}}/{{(index .Examples 0).Name}}{{if .HasContext}}?ctx=1{{end}}{{end}}</span>
    </div>
    <div class="browser-body">
      <iframe src="/{{.Name}}/{{(index .Examples 0).Name}}{{if .HasContext}}?ctx=1{{end}}"></iframe>
    </div>
  </div>
  {{else}}<div class="no-examples">no examples yet</div>
  {{end}}</details>
  {{end}}
</main>
<script>
function updateFrame(component) {
  var chips = component.querySelector('.example-chips');
  var activeChip = chips ? chips.querySelector('.example-chip.active') : null;
  var compName = component.querySelector('h3').textContent;
  var src = component.querySelector('iframe').src;
  var exName = activeChip ? activeChip.dataset.example : src.split('/').pop().split('?')[0];
  var url = '/' + compName + '/' + exName;
  var toggle = component.querySelector('.view-toggle-btn.active');
  if (toggle && toggle.dataset.view === 'in-shell') url += '?ctx=1';
  var frame = component.querySelector('.browser-frame');
  var displayUrl = activeChip && activeChip.dataset.displayUrl ? activeChip.dataset.displayUrl : url;
  frame.querySelector('.browser-url').textContent = displayUrl;
  frame.querySelector('iframe').src = url;
}
document.addEventListener('click', function(e) {
  var chip = e.target.closest('.example-chip');
  if (chip) {
    e.preventDefault();
    chip.closest('.example-chips').querySelectorAll('.example-chip').forEach(function(c) {
      c.classList.remove('active');
    });
    chip.classList.add('active');
    updateFrame(chip.closest('.component'));
    return;
  }
  var toggle = e.target.closest('.view-toggle-btn');
  if (toggle) {
    toggle.closest('.view-toggle').querySelectorAll('.view-toggle-btn').forEach(function(b) {
      b.classList.remove('active');
    });
    toggle.classList.add('active');
    updateFrame(toggle.closest('.component'));
  }
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

// contextTemplate wraps one rendered region example inside a workspace page shell.
var contextTemplate = template.Must(template.New("context").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{.Component}}/{{.Name}} (in shell) · ui gallery</title>
<link rel="stylesheet" href="/ui.css">
<script type="module" src="/ui.js"></script>
<style>
html, body { margin: 0; height: 100%; overflow: hidden; }
.demo-region {
  display: flex; align-items: center; justify-content: center;
  font-family: var(--yr-font-mono); font-size: var(--yr-font-size-xs);
  color: var(--yr-text-disabled);
}
</style>
</head>
<body>
<yr-page-workspace>
  {{if eq .Slot "nav"}}{{.HTML}}{{else -}}
  <yr-nav-shell>
    <div class="demo-region" style="width:100%">nav</div>
  </yr-nav-shell>
  {{- end}}
  {{if eq .Slot "sidebar"}}{{.HTML}}{{else -}}
  <yr-sidebar>
    <div class="demo-region" style="height:100%">sidebar</div>
  </yr-sidebar>
  {{- end}}
  <yr-resize-handle target="sidebar"></yr-resize-handle>
  <yr-content>
    <yr-content-main>
      <div class="demo-region" style="height:100%">main</div>
    </yr-content-main>
    {{if eq .Slot "drawer" -}}
    <yr-resize-handle target="drawer"></yr-resize-handle>
    {{.HTML}}
    {{- end}}
  </yr-content>
</yr-page-workspace>
</body>
</html>
`))

// examplePage is the data exampleTemplate renders.
type examplePage struct {
	Component string
	Name      string
	HTML      template.HTML
}

// contextPage is the data contextTemplate renders.
type contextPage struct {
	Component string
	Name      string
	Slot      string
	HTML      template.HTML
}

// gallery serves the ui examples and the library CSS.
type gallery struct {
	examples    map[string]templ.Component
	contextSlot map[string]string // component name → shell slot
	sections    []indexSection
	assets      fs.FS
}

// newHandler returns the gallery handler, serving ui.css from assets as /ui.css and the
// embedded gallery.css as /gallery.css.
func newHandler(assets fs.FS) http.Handler {
	groups := ui.ExampleGroups()

	byComponent := make(map[string][]ui.Example, len(groups))
	for _, g := range groups {
		byComponent[g.Component] = g.Examples
	}

	g := &gallery{
		examples:    make(map[string]templ.Component),
		contextSlot: make(map[string]string),
		sections:    make([]indexSection, len(sections)),
		assets:      assets,
	}

	for i, sec := range sections {
		is := indexSection{Name: sec.Name, Categories: make([]indexCategory, len(sec.Categories))}

		for j, cat := range sec.Categories {
			ic := indexCategory{Slug: cat.Slug, Name: cat.Name, Components: make([]indexComponent, len(cat.Components))}

			for k, comp := range cat.Components {
				_, hasCtx := cat.ContextSlot[comp]
				ic.Components[k] = indexComponent{Name: comp, HasContext: hasCtx, Examples: nil}

				if slot, ok := cat.ContextSlot[comp]; ok {
					g.contextSlot[comp] = slot
				}

				if exs, ok := byComponent[comp]; ok {
					for _, ex := range exs {
						g.examples[comp+"/"+ex.Name] = ex.Component
						ic.Components[k].Examples = append(
							ic.Components[k].Examples,
							indexExample{Name: ex.Name, DisplayURL: ex.DisplayURL},
						)
					}
				}
			}

			is.Categories[j] = ic
		}

		g.sections[i] = is
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", g.index)
	mux.HandleFunc("GET /ui.css", g.css)
	mux.HandleFunc("GET /gallery.css", g.galleryStyles)
	mux.HandleFunc("GET /ui.js", g.js)
	mux.HandleFunc("GET /fonts/{file}", g.font)
	mux.HandleFunc("GET /logo.png", g.logo)
	mux.HandleFunc("GET /{component}/{name}", g.example)

	return mux
}

// index renders the sidebar and the selected category's component list.
func (g *gallery) index(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("cat")

	var active *indexCategory

	for i := range g.sections {
		for j := range g.sections[i].Categories {
			if g.sections[i].Categories[j].Slug == slug {
				active = &g.sections[i].Categories[j]

				break
			}
		}
	}

	if active == nil {
		active = &g.sections[0].Categories[0]
		slug = active.Slug
	}

	g.write(w, indexTemplate, indexPage{
		Sections:   g.sections,
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

// font serves vendored woff2 font files.
func (g *gallery) font(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, g.assets, path.Join("fonts", path.Base(r.PathValue("file"))))
}

// logo serves the Yardrail logo.
func (g *gallery) logo(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, g.assets, "logo.png")
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
	html := template.HTML(buf.String())

	if slot, ok := g.contextSlot[component]; ok && r.URL.Query().Get("ctx") == "1" {
		g.write(w, contextTemplate, contextPage{Component: component, Name: name, Slot: slot, HTML: html})

		return
	}

	g.write(w, exampleTemplate, examplePage{Component: component, Name: name, HTML: html})
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
