package ui

import "github.com/a-h/templ"

const pageAuthBase = `<style>html,body{margin:0;height:100%}` +
	`.demo-region{display:flex;align-items:center;justify-content:center;` +
	`font-family:var(--yr-font-mono);font-size:var(--yr-font-size-xs);` +
	`color:var(--yr-text-subtle)}</style>`

func pageAuthExamples() []Example {
	return []Example{
		{Name: exampleDefault, Component: templ.Raw(
			pageAuthBase +
				`<yr-page-auth>` +
				`<div class="demo-region" style="width:320px;height:200px;` +
				`border:1px dashed var(--yr-border);` +
				`border-radius:var(--yr-radius-md)">` +
				`content` +
				`</div>` +
				`</yr-page-auth>`,
		), DisplayURL: ""},
		{Name: "with-auth-card", Component: templ.Raw(
			pageAuthBase +
				`<yr-page-auth>` +
				`<div class="yr-auth-card">` +
				`<p class="yr-auth-card-brand">Yardrail</p>` +
				`<form class="yr-auth-card-form" method="post" action="#">` +
				`<div class="yr-field"><label for="pa-user">Username</label>` +
				` <input type="text" id="pa-user" name="username"` +
				` autocomplete="username"></div>` +
				`<div class="yr-field"><label for="pa-pass">Password</label>` +
				` <input type="password" id="pa-pass" name="password"` +
				` autocomplete="current-password"></div>` +
				`<yr-button variant="primary" type="submit">Log in</yr-button>` +
				`</form>` +
				`<p class="yr-auth-card-footer">` +
				`No account? <a href="#">Sign up</a></p>` +
				`</div>` +
				`</yr-page-auth>`,
		), DisplayURL: ""},
	}
}
