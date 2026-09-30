package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/a-h/templ/parser/v2"
	"github.com/a-h/templ/parser/v2/visitor"
)

// Exit codes of run.
const (
	exitClean      = 0
	exitViolations = 1
	exitError      = 2
)

// errParse reports a .templ file that cannot be read or parsed.
var errParse = errors.New("cannot parse templ file")

// violation is one class= attribute in a .templ file.
type violation struct {
	path string
	line uint32
}

// String formats v as <path>:<line>: <message>.
func (v violation) String() string {
	return fmt.Sprintf("%s:%d: class= is not allowed in app templ", v.path, v.line)
}

// run lints every path, printing violations to stdout and errors to stderr, and returns the exit
// code.
func run(paths []string, stdout, stderr io.Writer) int {
	code := exitClean

	for _, path := range paths {
		found, err := lintFile(path)
		if err != nil {
			_, _ = fmt.Fprintln(stderr, err)
			code = exitError

			continue
		}

		for _, v := range found {
			_, _ = fmt.Fprintln(stdout, v)
		}

		if len(found) > 0 && code == exitClean {
			code = exitViolations
		}
	}

	return code
}

// lintFile parses the .templ file at path and returns its violations.
func lintFile(path string) ([]violation, error) {
	tf, err := parser.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %w", errParse, path, err)
	}

	return classAttributes(path, tf)
}

// classAttributes returns every class attribute in tf, in document order. templ's visitor walks
// every node kind, including element children, if/else, for, switch and templ element blocks, and
// the attributes inside conditional attributes.
func classAttributes(path string, tf *parser.TemplateFile) ([]violation, error) {
	var found []violation

	check := func(key parser.AttributeKey, r parser.Range) {
		if key.String() == "class" {
			// Range lines are zero-based.
			found = append(found, violation{path: path, line: r.From.Line + 1})
		}
	}

	v := visitor.New()
	v.ConstantAttribute = func(n *parser.ConstantAttribute) error {
		check(n.Key, n.Range)

		return nil
	}
	v.BoolConstantAttribute = func(n *parser.BoolConstantAttribute) error {
		check(n.Key, n.Range)

		return nil
	}
	v.ExpressionAttribute = func(n *parser.ExpressionAttribute) error {
		check(n.Key, n.Range)

		return nil
	}
	v.BoolExpressionAttribute = func(n *parser.BoolExpressionAttribute) error {
		check(n.Key, n.Range)

		return nil
	}

	err := tf.Visit(v)
	if err != nil {
		return nil, fmt.Errorf("walk %s: %w", path, err)
	}

	return found, nil
}
