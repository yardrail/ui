package ui

import "testing"

func TestRenderChatInput(t *testing.T) {
	t.Parallel()

	html := renderString(t, ChatInput(ChatInputProps{}))
	writeClassCoverage(t, "chat_input", "default", html)
}
