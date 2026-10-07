package ui

import "github.com/a-h/templ"

func settingsSectionExamples() []Example {
	return []Example{
		{Name: "avatar-row", Component: templ.Raw(
			`<section class="yr-settings-section">` +
				`<h2>Profile photo</h2>` +
				`<p class="yr-settings-section-success">Your photo has been updated.</p>` +
				`<div class="yr-settings-section-avatar-row">` +
				`<yr-avatar size="lg">BL</yr-avatar>` +
				`<form method="post" action="#" enctype="multipart/form-data">` +
				`<input type="file" name="avatar" accept="image/png,image/jpeg" required>` +
				` <yr-button variant="primary" type="submit">Upload</yr-button>` +
				`</form>` +
				`</div>` +
				`</section>`,
		), DisplayURL: ""},
		{Name: "field-row", Component: templ.Raw(
			`<section class="yr-settings-section">` +
				`<h2>Display name</h2>` +
				`<form class="yr-form-stacked" method="post" action="#">` +
				`<div class="yr-field"><label for="settings-display-name">Display name</label>` +
				` <input type="text" id="settings-display-name" name="display_name" value="Blake Lindqvist"></div>` +
				`<div><yr-button variant="primary" type="submit">Save</yr-button></div>` +
				`</form>` +
				`</section>`,
		), DisplayURL: ""},
	}
}
