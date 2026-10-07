package ui

import "github.com/a-h/templ"

// dropdownMenuChevronSize is the chevron size yr-account-menu draws in its own trigger.
const dropdownMenuChevronSize = 12

func dropdownMenuTrigger() string {
	return `<button slot="trigger" type="button">Actions ` +
		exampleIcon("chevron-down", `class="yr-dropdown-menu-chevron"`, dropdownMenuChevronSize) + `</button>`
}

func dropdownMenuExamples() []Example {
	return []Example{
		{Name: "closed", Component: templ.Raw(
			`<div style="display:flex;justify-content:flex-end;min-height:12rem">` +
				`<yr-dropdown-menu>` +
				dropdownMenuTrigger() +
				`<div slot="panel">` +
				`<a href="#" class="yr-dropdown-menu-item">Rename</a>` +
				`<a href="#" class="yr-dropdown-menu-item">Duplicate</a>` +
				`<div class="yr-dropdown-menu-divider"></div>` +
				`<button type="button" class="yr-dropdown-menu-item">Delete</button>` +
				`</div>` +
				`</yr-dropdown-menu>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: "open", Component: templ.Raw(
			`<div style="display:flex;justify-content:flex-end;min-height:20rem">` +
				`<yr-dropdown-menu open>` +
				dropdownMenuTrigger() +
				`<div slot="panel">` +
				`<div class="yr-dropdown-menu-header">` +
				`<yr-avatar size="lg">JE</yr-avatar>` +
				`<div>` +
				`<div class="yr-dropdown-menu-header-name">Jordan Ellery</div>` +
				`<div class="yr-dropdown-menu-header-email">jordan@ellerylogistics.com</div>` +
				`</div>` +
				`</div>` +
				`<a href="#" class="yr-dropdown-menu-item">Organization</a>` +
				`<a href="#" class="yr-dropdown-menu-item">Account settings</a>` +
				`<div class="yr-dropdown-menu-divider"></div>` +
				`<button type="button" class="yr-dropdown-menu-item">Log out</button>` +
				`</div>` +
				`</yr-dropdown-menu>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: "account-menu", Component: templ.Raw(
			`<div style="display:flex;justify-content:flex-end;min-height:20rem">` +
				`<yr-account-menu open name="Jordan Ellery" email="jordan@ellerylogistics.com">` +
				`<span slot="avatar">JE</span>` +
				`<a href="#" class="yr-dropdown-menu-item">Organization</a>` +
				`<a href="#" class="yr-dropdown-menu-item">Account settings</a>` +
				`<div class="yr-dropdown-menu-divider"></div>` +
				`<button type="button" class="yr-dropdown-menu-item">Log out</button>` +
				`</yr-account-menu>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
