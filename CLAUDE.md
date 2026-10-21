# yardrail/ui

Typed templ component library for Yardrail's web UI. Go package `github.com/yardrail/ui/pkg`.

## Project structure

```
yardrail/ui/
├── build.mjs                    # esbuild config (entry points: styles.css, index.ts)
├── tsconfig.json                # include: ["pkg/**/*.ts"]
├── package.json
├── Taskfile.yml
├── lefthook.yml
├── styles.css                   # CSS entry point — @import "./tokens.css", "./pkg/pill.css", ...
├── index.ts                     # TS entry point — import "./pkg/pill", ...
├── tokens.css                   # shared design tokens
├── base.css                     # reset / base styles
├── fonts.css                    # @font-face declarations
│
├── pkg/                         # Go package: github.com/yardrail/ui/pkg
│   ├── doc.go
│   ├── assets.go                # go:embed dist
│   ├── icon.go
│   ├── attrs.go
│   │
│   │── # ── Components ──────────────────────────
│   │
│   ├── pill.css                 # yr-pill styling
│   ├── pill.ts                  # customElements.define('yr-pill', ...)
│   ├── pill.go                  # PillProps, Pill() constructor
│   ├── pill.templ               # templ pill(...)
│   ├── pill_examples.go
│   │
│   ├── stat.css
│   ├── stat.ts
│   ├── stat.go
│   ├── stat.templ
│   ├── stat_examples.go
│   │
│   ├── workflow_card.css
│   ├── workflow_card.ts
│   ├── workflow_card.go
│   ├── workflow_card.templ
│   ├── workflow_card_examples.go
│   │
│   ├── shell.css
│   ├── shell.go
│   ├── shell.templ
│   ├── shell_examples.go
│   │
│   ├── sidebar.css
│   ├── sidebar.ts
│   ├── sidebar.go
│   ├── sidebar.templ
│   ├── sidebar_examples.go
│   │
│   │── # ── Mocks (full-page demos) ─────────────
│   │
│   ├── portal_mocks.templ
│   ├── portal_mock_helpers.go
│   ├── portal_home_examples.go
│   ├── ...
│   ├── msp_mocks.templ
│   ├── msp_mock_helpers.go
│   ├── msp_dashboard_examples.go
│   ├── ...
│   │
│   │── # ── Build output (committed, embedded) ──
│   │
│   └── dist/                    # build output (committed, embedded)
│       ├── ui.css
│       ├── ui.js
│       ├── fonts/
│       └── icons/
│
└── cmd/
    └── gallery/
        ├── main.go              # imports github.com/yardrail/ui/pkg
        ├── server.go
        ├── examples.go          # Example struct, ExampleGroups() registry
        ├── snapshot_test.go     # golden-file tests against testdata/
        ├── gallery.css          # gallery chrome styles
        ├── mocks.css            # mock-only styles
        └── testdata/            # golden HTML snapshots
            ├── pill/default.html
            ├── workflow-card/default.html
            └── ...
```

## Component design

Each component is four co-located files in `pkg/`:

**Go Props + constructor** (`workflow_card.go`):

```go
package ui

import "github.com/a-h/templ"

type WorkflowCardProps struct {
    Status      templ.Component
    Runs        string
    SuccessRate string
    Action      templ.Component
    Attrs       templ.Attributes
}

func WorkflowCard(name, footer string, p WorkflowCardProps) templ.Component {
    return workflowCard(name, footer, &p)
}
```

**Templ render function** (`workflow_card.templ`):

```templ
package ui

templ workflowCard(name, footer string, p *WorkflowCardProps) {
    <yr-workflow-card { filterAttrs("WorkflowCard", p.Attrs)... }>
        <div class="yr-workflow-card-header">
            <span class="yr-workflow-card-name">{ name }</span>
            if p.Status != nil {
                @p.Status
            }
        </div>
        <div class="yr-workflow-card-stats">
            @Stat(p.Runs, StatProps{Label: "Runs"})
            @Stat(p.SuccessRate, StatProps{Label: "Success"})
        </div>
        <div class="yr-workflow-card-footer">
            <span>{ footer }</span>
            if p.Action != nil {
                @p.Action
            }
        </div>
    </yr-workflow-card>
}
```

**CSS** (`workflow_card.css`):

```css
yr-workflow-card { display: block; background: var(--yr-surface); ... }
yr-workflow-card .yr-workflow-card-header { ... }
yr-workflow-card .yr-workflow-card-stats { ... }
yr-workflow-card .yr-workflow-card-footer { ... }
```

**TypeScript** (`workflow_card.ts`):

```ts
export class YrWorkflowCard extends HTMLElement {}
customElements.define('yr-workflow-card', YrWorkflowCard);
```

## Component conventions

- Required args are positional (usually the accessible name). Optional config goes in a `Props` struct with zero-value defaults.
- `Props.Attrs` carries passthrough HTML attributes (htmx, data-*), filtered by `filterAttrs()` before rendering.
- Compound components accept `templ.Component` fields for composable slots (e.g. `Status`, `Action`).
- Templ is the source of truth for markup. Never edit `*_templ.go` — run `templ generate`.
- Web components use the `yr-` prefix (e.g. `yr-sidebar`, `yr-pill`).
- Icons: `Icon(name)` renders a Lucide SVG from `dist/icons/`.

## Usage in a mock page

```templ
templ portalSpacePage() {
    @portalShell("Marketing Ops · Yardrail", portalNav("space-marketing")) {
        <div class="portal-grid portal-grid-3">
            @WorkflowCard("Lead Scoring", "Waiting for input", WorkflowCardProps{
                Status: Pill("Waiting", PillProps{Tone: "warning", Shape: ShapePill}),
                Runs:   "42",
                SuccessRate: "91%",
                Action: Button("View", ButtonProps{Variant: VariantSecondary}),
            })
        </div>
    }
}
```

## Usage by a consuming project

```go
import ui "github.com/yardrail/ui/pkg"

// Serve the bundled CSS/JS
http.Handle("/static/ui/", http.StripPrefix("/static/ui/", http.FileServerFS(ui.Assets)))
```

```templ
import ui "github.com/yardrail/ui/pkg"

templ myPage() {
    @ui.Shell("My Page", ui.ShellProps{}) {
        @ui.WorkflowCard("Nightly Sync", "Last run 1h ago", ui.WorkflowCardProps{
            Status: ui.Pill("Active", ui.PillProps{Tone: "success"}),
            Runs:   "89",
            SuccessRate: "97%",
        })
    }
}
```

## UI mocks

Full-page demos live alongside the components as `*_examples.go` files. Two mock families:

- **Portal mocks** (`portal_mocks.templ`): end-user portal pages.
- **MSP mocks** (`msp_mocks.templ`): managed service provider admin pages.

Both use `Shell` + `Sidebar` + `Navbar` components for layout. Mock-only CSS lives in `cmd/gallery/mocks.css`. Each mock page is a templ function (e.g. `portalHomePage()`) referenced by a slim `*_examples.go` file.

## Gallery

`cmd/gallery/` serves all examples at `:6007`. Run with `task gallery` (uses `air` for hot reload). Examples that render a `<!doctype html>` are served as-is; fragments get wrapped in a minimal page shell.

## Commands

- `task build` — build CSS/JS
- `task gen` — compile `.templ` files
- `task test` — Go tests (includes snapshot tests against `testdata/`)
- `task lint` — golangci-lint + stylelint
- `task gallery` — dev server with hot reload

## Git hooks (lefthook)

Pre-commit runs `gen:check`, `lint`, `test` in parallel. Commit messages validated by commitlint. Use conventional commits (`feat:`, `fix:`, `refactor:`, `chore:`).
