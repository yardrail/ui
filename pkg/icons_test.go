package ui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestRenderIcons(t *testing.T) {
	t.Parallel()

	icons := map[string]templ.Component{
		"alert-triangle":    IconAlertTriangle(),
		"chevron-down":      IconChevronDown(),
		"building-2":        IconBuilding2(),
		"check-circle":      IconCheckCircle(),
		"circle-x":          IconCircleX(),
		"clock":             IconClock(),
		"credit-card":       IconCreditCard(),
		"database":          IconDatabase(),
		"file":              IconFile(),
		"file-audio":        IconFileAudio(),
		"file-image":        IconFileImage(),
		"file-spreadsheet":  IconFileSpreadsheet(),
		"file-text":         IconFileText(),
		"folder":            IconFolder(),
		"git-branch":        IconGitBranch(),
		"home":              IconHome(),
		"layout-dashboard":  IconLayoutDashboard(),
		"layout-grid":       IconLayoutGrid(),
		"list":              IconList(),
		"mail":              IconMail(),
		"message-square":    IconMessageSquare(),
		"pencil":            IconPencil(),
		"play":              IconPlay(),
		"plug":              IconPlug(),
		"refresh-cw":        IconRefreshCw(),
		"search":            IconSearch(),
		"settings":          IconSettings(),
		"shield":            IconShield(),
		"star":              IconStar(),
		"train-front":       IconTrainFront(),
		"user":              IconUser(),
		"user-plus":         IconUserPlus(),
		"users":             IconUsers(),
		"x":                 IconX(),
		"panel-left-close":  IconPanelLeftClose(),
		"panel-left-open":   IconPanelLeftOpen(),
		"panel-right-close": IconPanelRightClose(),
		"panel-right-open":  IconPanelRightOpen(),
		"zap":               IconZap(),
	}

	for name, icon := range icons {
		html := renderString(t, icon)
		writeClassCoverage(t, "icons", name, html)
	}
}
