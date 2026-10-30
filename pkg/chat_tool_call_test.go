package ui

import (
	"fmt"
	"testing"
)

func TestRenderChatToolCall(t *testing.T) {
	t.Parallel()

	i := 0

	for _, status := range (ToolCallStatus{}).All() {
		for _, mutation := range []bool{false, true} {
			html := renderString(t, ChatToolCall(ChatToolCallProps{
				Name:     "Read file",
				Status:   status,
				Mutation: mutation,
			}))
			writeClassCoverage(t, "chat_tool_call", fmt.Sprintf("%03d", i), html)
			i++
		}
	}
}

func TestRenderChatToolCalls(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		props ChatToolCallsProps
	}{
		{"running", ChatToolCallsProps{Count: 3, RunningCount: 1}},
		{"done", ChatToolCallsProps{Count: 3, DoneCount: 3}},
		{"error", ChatToolCallsProps{Count: 3, ErrorCount: 1}},
		{"no_errors", ChatToolCallsProps{Count: 2, DoneCount: 2}},
	}

	for _, tt := range tests {
		html := renderString(t, withChildren(
			ChatToolCalls(tt.props),
			text("tool calls"),
		))
		writeClassCoverage(t, "chat_tool_calls", tt.name, html)
	}
}

func TestChatToolCallsSummary(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    ChatToolCallsProps
		want string
	}{
		{"running", ChatToolCallsProps{RunningCount: 1}, "Working…"},
		{"error", ChatToolCallsProps{ErrorCount: 1}, "Done with errors"},
		{"done", ChatToolCallsProps{}, "Done"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.p.summary(); got != tt.want {
				t.Errorf("summary() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestChatToolCallsStatusClass(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    ChatToolCallsProps
		want string
	}{
		{"running", ChatToolCallsProps{RunningCount: 1}, "running"},
		{"error", ChatToolCallsProps{ErrorCount: 1}, "error"},
		{"done", ChatToolCallsProps{}, "done"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.p.statusClass(); got != tt.want {
				t.Errorf("statusClass() = %q, want %q", got, tt.want)
			}
		})
	}
}
