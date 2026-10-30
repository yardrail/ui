package ui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestRenderWithAttrs(t *testing.T) {
	t.Parallel()

	attrs := templ.Attributes{"data-testid": "x"}

	cases := []struct {
		name string
		c    templ.Component
	}{
		{"alert", withChildren(Alert(&AlertProps{Tone: ToneSuccess, Attrs: attrs}), text("msg"))},
		{"pill", Pill("tag", &PillProps{Tone: ToneAccent, Attrs: attrs})},
		{"avatar", Avatar("BL", &AvatarProps{Attrs: attrs})},
		{"card", withChildren(Card(&CardProps{Attrs: attrs}), text("body"))},
		{"card-interactive", withChildren(Card(&CardProps{Interactive: true, Attrs: attrs}), text("body"))},
		{"icon-badge", withChildren(IconBadge(ToneDanger, &IconBadgeProps{Attrs: attrs}), text("icon"))},
		{"breadcrumb", withChildren(Breadcrumb(&BreadcrumbProps{Attrs: attrs}), text("crumbs"))},
		{"copy-field", CopyField("secret", &CopyFieldProps{Attrs: attrs})},
		{"connector-pill", withChildren(ConnectorPill(&ConnectorPillProps{Title: "Slack", Attrs: attrs}), text("SL"))},
		{"content", withChildren(Content(&ContentProps{Attrs: attrs}), text("body"))},
		{"content-main", withChildren(ContentMain(&ContentMainProps{Attrs: attrs}), text("body"))},
		{"create-panel", withChildren(CreatePanel(&CreatePanelProps{Attrs: attrs}), text("form"))},
		{"download-list", withChildren(DownloadList(&DownloadListProps{Attrs: attrs}), text("items"))},
		{"drawer", withChildren(Drawer(&DrawerProps{Attrs: attrs}), text("body"))},
		{"dropdown-menu", withChildren(DropdownMenu(&DropdownMenuProps{Attrs: attrs}), text("items"))},
		{"empty-state", withChildren(EmptyState(&EmptyStateProps{Attrs: attrs}), text("empty"))},
		{"hitl-prompt", withChildren(HitlPrompt(&HitlPromptProps{Tone: ToneWarning, Attrs: attrs}), text("prompt"))},
		{"journal-log", withChildren(JournalLog(&JournalLogProps{Attrs: attrs}), text("log"))},
		{"list-row", withChildren(ListRow(&ListRowProps{Attrs: attrs}), text("row"))},
		{"live-duration", LiveDuration("2024-01-01T00:00:00Z", &LiveDurationProps{Attrs: attrs})},
		{"meta-row", withChildren(MetaRow(&MetaRowProps{Attrs: attrs}), text("meta"))},
		{"review-card", withChildren(ReviewCard(&ReviewCardProps{Attrs: attrs}), text("review"))},
		{"section-header", withChildren(SectionHeader(&SectionHeaderProps{Attrs: attrs}), text("header"))},
		{"settings-section", withChildren(SettingsSection(&SettingsSectionProps{Attrs: attrs}), text("settings"))},
		{
			"status-banner",
			withChildren(StatusBanner(&StatusBannerProps{Tone: ToneSuccess, Attrs: attrs}), text("banner")),
		},
		{"action-card", withChildren(ActionCard(&ActionCardProps{Attrs: attrs}), text("card"))},
		{"auth-card", withChildren(AuthCard(&AuthCardProps{Attrs: attrs}), text("form"))},
		{"account-menu", withChildren(AccountMenu(&AccountMenuProps{Attrs: attrs}), text("menu"))},
		{"brand", withChildren(Brand(&BrandProps{Attrs: attrs}), text("Yardrail"))},
		{"brand-with-href", withChildren(Brand(&BrandProps{Href: "/", Attrs: attrs}), text("Yardrail"))},
		{"stat", Stat("42", &StatProps{Label: "Users", Attrs: attrs})},
		{"rating", Rating("4.5", &RatingProps{Reviews: "12", Attrs: attrs})},
		{"resize-handle", ResizeHandle(&ResizeHandleProps{Attrs: attrs})},
		{"page-auth", withChildren(PageAuth(&PageAuthProps{Attrs: attrs}), text("auth"))},
		{"page-workspace", withChildren(PageWorkspace(&PageWorkspaceProps{Attrs: attrs}), text("workspace"))},
		{"tabs", withChildren(Tabs(&TabsProps{Attrs: attrs}), text("tab content"))},
	}

	for _, tc := range cases {
		html := renderString(t, tc.c)
		writeClassCoverage(t, "attrs", tc.name, html)
	}
}
