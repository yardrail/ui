package ui

import "testing"

func TestCappyVariantImageFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    CappyVariant
		want string
	}{
		{"zero", CappyVariant{}, "wave.png"},
		{"excited", CappyExcited, "excited.png"},
		{"thinking", CappyThinking, "thinking.png"},
		{"pointing", CappyPointing, "pointing.png"},
		{"see-no-evil", CappySeeNoEvil, "see-no-evil.png"},
		{"train-engineer", CappyTrainEngineer, "train-engineer.png"},
		{"construction-worker", CappyConstructionWorker, "construction-worker.png"},
		{"learner", CappyLearner, "learner.png"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.v.ImageFile(); got != tt.want {
				t.Errorf("ImageFile() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestChatSenderString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    ChatSender
		want string
	}{
		{"user", ChatSenderUser, "user"},
		{"cappy", ChatSenderCappy, "cappy"},
		{"human", ChatSenderHuman, "human"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToolCallStatusString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    ToolCallStatus
		want string
	}{
		{"pending", ToolCallPending, "pending"},
		{"running", ToolCallRunning, "running"},
		{"done", ToolCallDone, "done"},
		{"error", ToolCallError, "error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCappyVariantAll(t *testing.T) {
	t.Parallel()

	all := (CappyVariant{}).All()
	if len(all) != 8 {
		t.Errorf("All() returned %d values, want 8", len(all))
	}
}

func TestChatSenderAll(t *testing.T) {
	t.Parallel()

	all := (ChatSender{}).All()
	if len(all) != 3 {
		t.Errorf("All() returned %d values, want 3", len(all))
	}
}

func TestToolCallStatusAll(t *testing.T) {
	t.Parallel()

	all := (ToolCallStatus{}).All()
	if len(all) != 4 {
		t.Errorf("All() returned %d values, want 4", len(all))
	}
}
