package ui

import (
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

var (
	// errAttrNotAllowed reports an Attrs key outside the allow-list.
	errAttrNotAllowed = errors.New("attribute not allowed in Attrs")
	// errRequiredArgEmpty reports an empty required positional argument.
	errRequiredArgEmpty = errors.New("required argument is empty")
)

// allowedAttrNames are the exact Attrs keys accepted besides the hx-* and data-* families.
var allowedAttrNames = []string{"id", "aria-label", "aria-labelledby", "aria-describedby"}

// report handles a call-site mistake: it panics under go test, so the mistake fails the test
// that rendered it, and otherwise logs a warning and lets the render continue.
func report(component string, err error) {
	err = fmt.Errorf("ui.%s: %w", component, err)
	if testing.Testing() {
		panic(err)
	}

	slog.Default().Warn("ui component misuse", "component", component, "err", err)
}

// attrAllowed reports whether key may pass through Attrs: hx-* except hx-on* and hx-confirm,
// data-*, id, and the aria naming and description attributes. A data-hx-* key follows the hx-*
// rules, since htmx reads it as the same attribute.
func attrAllowed(key string) bool {
	key = htmxName(strings.ToLower(key))

	switch {
	case strings.HasPrefix(key, "hx-on"), key == "hx-confirm":
		return false
	case strings.HasPrefix(key, "hx-"), strings.HasPrefix(key, "data-"):
		return true
	default:
		return slices.Contains(allowedAttrNames, key)
	}
}

// htmxName strips one leading data- from a data-hx-* key, returning the attribute name htmx
// reads; any other key is returned unchanged.
func htmxName(key string) string {
	if rest, ok := strings.CutPrefix(key, "data-"); ok && strings.HasPrefix(rest, "hx-") {
		return rest
	}

	return key
}

// filterAttrs returns the allowed subset of attrs, reporting every dropped key in sorted order.
func filterAttrs(component string, attrs templ.Attributes) templ.Attributes {
	kept := make(templ.Attributes, len(attrs))

	for _, key := range slices.Sorted(maps.Keys(attrs)) {
		if !attrAllowed(key) {
			report(component, fmt.Errorf("%w: %q", errAttrNotAllowed, key))

			continue
		}

		kept[key] = attrs[key]
	}

	return kept
}

// requireArg reports an empty required positional argument named arg of component.
func requireArg(component, arg, value string) {
	if value == "" {
		report(component, fmt.Errorf("%w: %s", errRequiredArgEmpty, arg))
	}
}
