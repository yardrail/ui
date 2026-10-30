package ui

import "testing"

func TestRenderSettingsSection(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(SettingsSection(&SettingsSectionProps{}), text("settings")))
	writeClassCoverage(t, "settings_section", "default", html)
}
