package ui

import "github.com/a-h/templ"

// SettingsSectionProps configures a SettingsSection.
type SettingsSectionProps struct {
	Attrs templ.Attributes
}

// SettingsSection renders a settings card with heading and fields. Children are the section content.
func SettingsSection(p *SettingsSectionProps) templ.Component {
	return settingsSection(p)
}
