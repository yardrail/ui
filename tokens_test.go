package ui

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// scaleTokenRE matches the declaration of a scale token in src/tokens.css. Control properties
// (--yr-control-*) and font families (--yr-font-body) are not scale tokens.
var scaleTokenRE = regexp.MustCompile(`(?m)^\s*(--yr-(?:space|font-size|radius|measure|side)-[a-z0-9-]+)\s*:`)

// TestTokensMatchCSS fails when a scale token in src/tokens.css has no Go value in tokens.go, or
// a Go value names a token the stylesheet does not define.
func TestTokensMatchCSS(t *testing.T) {
	t.Parallel()

	prefixes := map[string]string{
		"Space":     "--yr-space-",
		"FontSize":  "--yr-font-size-",
		"Radius":    "--yr-radius-",
		"Measure":   "--yr-measure-",
		"SideWidth": "--yr-side-",
	}

	css, err := os.ReadFile("src/tokens.css")
	if err != nil {
		t.Fatal(err)
	}

	inCSS := make(map[string]bool)
	for _, m := range scaleTokenRE.FindAllSubmatch(css, -1) {
		inCSS[string(m[1])] = true
	}

	inGo := make(map[string]string)

	for _, v := range declaredTokens(t, "tokens.go") {
		prefix, ok := prefixes[v.typ]
		if !ok {
			t.Errorf("%s has type %s, which is not a scale type", v.name, v.typ)

			continue
		}

		if !strings.HasPrefix(v.property, prefix) {
			t.Errorf("%s carries %s, want a %s* token", v.name, v.property, prefix)
		}

		if other, dup := inGo[v.property]; dup {
			t.Errorf("%s and %s both carry %s", other, v.name, v.property)
		}

		inGo[v.property] = v.name

		if !inCSS[v.property] {
			t.Errorf("%s carries %s, which src/tokens.css does not define", v.name, v.property)
		}
	}

	for property := range inCSS {
		if _, ok := inGo[property]; !ok {
			t.Errorf("src/tokens.css defines %s, which has no Go value in tokens.go", property)
		}
	}
}

// declaredToken is one package-level value of a scale type.
type declaredToken struct {
	name     string
	typ      string
	property string
}

// declaredTokens parses the Go file at path and returns every package-level variable declared as
// a composite literal with a v field, such as Space4 = Space{v: "--yr-space-4"}.
func declaredTokens(t *testing.T, path string) []declaredToken {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}

	var out []declaredToken

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.VAR {
			continue
		}

		for _, spec := range gen.Specs {
			value, isValue := spec.(*ast.ValueSpec)
			if !isValue {
				continue
			}

			for i, name := range value.Names {
				if i >= len(value.Values) {
					t.Fatalf("%s: %s has no value", path, name.Name)
				}

				out = append(out, parseToken(t, name.Name, value.Values[i]))
			}
		}
	}

	if len(out) == 0 {
		t.Fatalf("%s declares no token values", path)
	}

	return out
}

// parseToken reads the type and custom property out of a Type{v: "--yr-..."} literal.
func parseToken(t *testing.T, name string, expr ast.Expr) declaredToken {
	t.Helper()

	lit, ok := expr.(*ast.CompositeLit)
	if !ok {
		t.Fatalf("%s is not a composite literal", name)
	}

	typ, ok := lit.Type.(*ast.Ident)
	if !ok || len(lit.Elts) != 1 {
		t.Fatalf("%s is not of the form Type{v: \"--yr-...\"}", name)
	}

	kv, ok := lit.Elts[0].(*ast.KeyValueExpr)
	if !ok {
		t.Fatalf("%s does not set its field by name", name)
	}

	str, ok := kv.Value.(*ast.BasicLit)
	if !ok || str.Kind != token.STRING {
		t.Fatalf("%s does not carry a string literal", name)
	}

	property, err := strconv.Unquote(str.Value)
	if err != nil {
		t.Fatalf("%s: %v", name, err)
	}

	return declaredToken{name: name, typ: typ.Name, property: property}
}
