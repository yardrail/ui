package ui

import (
	"fmt"
	"testing"
)

func TestRenderChatFab(t *testing.T) {
	t.Parallel()

	i := 0

	for _, v := range (CappyVariant{}).All() {
		for _, badge := range (ChatFabBadge{}).All() {
			html := renderString(t, ChatFab(ChatFabProps{
				Variant: v,
				Badge:   badge,
			}))
			writeClassCoverage(t, "chat_fab", fmt.Sprintf("%03d", i), html)
			i++
		}
	}
}
