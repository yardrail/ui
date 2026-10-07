package ui

import "github.com/a-h/templ"

func authCardExamples() []Example {
	return []Example{
		{Name: "sign-in", Component: templ.Raw(
			`<div class="yr-auth-card">` +
				`<p class="yr-auth-card-brand">Yardrail</p>` +
				`<form class="yr-auth-card-form" method="post" action="#">` +
				`<div class="yr-field"><label for="auth-username">Username</label>` +
				` <input type="text" id="auth-username" name="username" autocomplete="username"></div>` +
				`<div class="yr-field"><label for="auth-password">Password</label>` +
				` <input type="password" id="auth-password" name="password" autocomplete="current-password"></div>` +
				`<yr-button variant="primary" type="submit">Log in</yr-button>` +
				`</form>` +
				`<p class="yr-auth-card-footer">No account? <a href="#">Sign up</a></p>` +
				`</div>`,
		), DisplayURL: ""},
		{Name: "with-error", Component: templ.Raw(
			`<div class="yr-auth-card">` +
				`<p class="yr-auth-card-brand">Yardrail</p>` +
				`<form class="yr-auth-card-form" method="post" action="#">` +
				`<yr-alert tone="danger">Wrong username or password.</yr-alert>` +
				`<div class="yr-field"><label for="auth-error-username">Username</label>` +
				` <input type="text" id="auth-error-username" name="username" value="pat"></div>` +
				`<div class="yr-field"><label for="auth-error-password">Password</label>` +
				` <input type="password" id="auth-error-password" name="password"></div>` +
				`<yr-button variant="primary" type="submit">Log in</yr-button>` +
				`</form>` +
				`<p class="yr-auth-card-footer">No account? <a href="#">Sign up</a></p>` +
				`</div>`,
		), DisplayURL: ""},
	}
}
