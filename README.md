# ui

`github.com/yardrail/ui` is Yardrail's shared UI library: typed [templ](https://templ.guide)
components in Go, plus the CSS and web components they render against.

The Yardrail name and logo are not licensed for reuse.

## Development

Prerequisites: Go (version in `go.mod`), Node.js with `corepack enable` (pnpm), [Task](https://taskfile.dev),
[templ](https://templ.guide) (`go install github.com/a-h/templ/cmd/templ@<version in go.mod>`),
[golangci-lint](https://golangci-lint.run) v2, [lefthook](https://github.com/evilmartians/lefthook) and
[commitlint](https://commitlint.js.org) (`npm install -g @commitlint/cli @commitlint/config-conventional`).

Run `lefthook install` once after cloning. Pre-commit runs `task gen:check`, `task lint` and
`task test`; commit-msg runs `commitlint` (conventional commits: `feat(scope): ...`,
`fix(scope): ...`, `chore(scope): ...`).

| Task | What it does |
| --- | --- |
| `task gen` | Compile `.templ` files to `*_templ.go` |
| `task build` | Build `src/` to `dist/` |
| `task gen:check` | Regenerate `*_templ.go` and `dist/`, fail if the committed output differs |
| `task lint` | Go lints (`lint:go`) and stylelint (`lint:css`) |
| `task test` | Go tests |
| `task gallery` | Serve every component example on :6007 with live reload |
