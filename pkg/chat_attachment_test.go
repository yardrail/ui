package ui

import "testing"

func TestRenderChatAttachment(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		typ  ChatAttachmentType
	}{
		{"image", AttachmentImage},
		{"audio", AttachmentAudio},
		{"file", AttachmentFile},
	} {
		html := renderString(t, ChatAttachment(ChatAttachmentProps{
			Type: tt.typ,
			Name: "test.txt",
			URL:  "/files/test.txt",
		}))
		writeClassCoverage(t, "chat_attachment", tt.name, html)
	}
}

func TestRenderChatAttachments(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(ChatAttachments(), text("attachments")))
	writeClassCoverage(t, "chat_attachments", "default", html)
}
