package ui

import "testing"

func TestRenderChatMessage(t *testing.T) {
	t.Parallel()

	for _, sender := range (ChatSender{}).All() {
		html := renderString(t, withChildren(
			ChatMessage(ChatMessageProps{
				Sender:    sender,
				Name:      "Sender",
				AvatarURL: "/avatar.png",
				Timestamp: "2:34 PM",
			}),
			text("Hello"),
		))
		writeClassCoverage(t, "chat_message", sender.String(), html)
	}
}
