// Command uilint checks app .templ files against the ui library rules. It implements one rule:
// no class= attribute, since styling belongs to ui components. Usage:
//
//	go run github.com/yardrail/ui/cmd/uilint FILE.templ...
//
// It exits 0 with no violations, 1 with any, and 2 when a file cannot be read or parsed.
package main

import "os"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
