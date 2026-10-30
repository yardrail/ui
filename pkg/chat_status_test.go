package ui

import "testing"

func TestRenderChatTyping(t *testing.T) {
	t.Parallel()

	for _, s := range (ChatSender{}).All() {
		html := renderString(t, ChatTyping(ChatTypingProps{Sender: s, Name: "Cappy"}))
		writeClassCoverage(t, "chat_typing", s.String(), html)
	}

	html := renderString(t, ChatTyping(ChatTypingProps{
		Sender: ChatSenderCappy, Name: "Cappy", Working: true,
	}))
	writeClassCoverage(t, "chat_typing", "working", html)
}

func TestRenderChatJoin(t *testing.T) {
	t.Parallel()

	html := renderString(t, ChatJoin(ChatJoinProps{Name: "Sarah"}))
	writeClassCoverage(t, "chat_join", "default", html)
}
