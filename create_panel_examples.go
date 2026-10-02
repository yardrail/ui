package ui

import "github.com/a-h/templ"

func createPanelExamples() []Example {
	return []Example{
		{Name: "expanded", Component: templ.Raw(
			`<details class="yr-create-panel" open>` +
				`<summary>Add a trigger</summary>` +
				`<form method="post" action="#" class="yr-create-panel-form">` +
				`<label>Name <input type="text" name="name" placeholder="e.g. Nightly sync" required></label>` +
				`<label>Description <input type="text" name="description" placeholder="Optional"></label>` +
				`<label>Owning team <select name="team">` +
				`<option value="">No owning team (org-wide)</option>` +
				`<option value="7">Platform</option>` +
				`</select></label>` +
				`<label class="yr-create-panel-check">` +
				`<input type="checkbox" name="isEnabled" checked> Turn on right away</label>` +
				`<yr-button variant="primary" type="submit">Add trigger</yr-button>` +
				`</form>` +
				`</details>`,
		)},
		{Name: "collapsed", Component: templ.Raw(
			`<details class="yr-create-panel">` +
				`<summary>Add a space</summary>` +
				`<form method="post" action="#" class="yr-create-panel-form">` +
				`<label>Name <input type="text" name="name" placeholder="e.g. Platform Tooling" required></label>` +
				`<yr-button variant="primary" type="submit">Add space</yr-button>` +
				`</form>` +
				`</details>`,
		)},
	}
}
