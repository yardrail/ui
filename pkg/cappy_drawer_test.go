package ui

import (
	"fmt"
	"testing"
)

func TestRenderCappyDrawer(t *testing.T) {
	t.Parallel()

	for i, v := range (CappyVariant{}).All() {
		html := renderString(t, withChildren(
			CappyDrawer(CappyDrawerProps{Variant: v}),
			text("chat log"),
		))
		writeClassCoverage(t, "cappy_drawer", fmt.Sprintf("variant-%03d", i), html)
	}

	html := renderString(t, withChildren(
		CappyDrawer(CappyDrawerProps{Open: true, AutoApprove: true, Placeholder: "Ask Cappy"}),
		text("chat log"),
	))
	writeClassCoverage(t, "cappy_drawer", "open-autoapprove", html)

	html = renderString(t, withChildren(
		CappyDrawer(CappyDrawerProps{Support: SupportOnline}),
		text("chat log"),
	))
	writeClassCoverage(t, "cappy_drawer", "support-online", html)

	html = renderString(t, withChildren(
		CappyDrawer(CappyDrawerProps{Support: SupportOffline}),
		text("chat log"),
	))
	writeClassCoverage(t, "cappy_drawer", "support-offline", html)

	html = renderString(t, withChildren(
		CappyDrawer(CappyDrawerProps{
			SupportName:      "Sarah",
			SupportAvatarURL: "/avatar.png",
		}),
		text("chat log"),
	))
	writeClassCoverage(t, "cappy_drawer", "with-support-person", html)

	html = renderString(t, withChildren(
		CappyDrawer(CappyDrawerProps{
			SupportName: "Sarah",
		}),
		text("chat log"),
	))
	writeClassCoverage(t, "cappy_drawer", "support-name-only", html)
}
