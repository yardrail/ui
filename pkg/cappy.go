package ui

// CappyVariant identifies a visual variant of the Cappy mascot. The consuming app sets this
// based on route or state; the library does not own routing logic.
type CappyVariant struct{ v string }

// Cappy visual variants.
var (
	CappyWave               = CappyVariant{}
	CappyExcited            = CappyVariant{v: "excited"}
	CappyThinking           = CappyVariant{v: "thinking"}
	CappyPointing           = CappyVariant{v: "pointing"}
	CappySeeNoEvil          = CappyVariant{v: "see-no-evil"}
	CappyTrainEngineer      = CappyVariant{v: "train-engineer"}
	CappyConstructionWorker = CappyVariant{v: "construction-worker"}
	CappyLearner            = CappyVariant{v: "learner"}
)

// ImageFile returns the filename of the variant's PNG in dist/cappy/.
func (v CappyVariant) ImageFile() string {
	if v.v == "" {
		return "wave.png"
	}

	return v.v + ".png"
}

// ChatSender identifies who sent a chat message.
type ChatSender struct{ v string }

// Chat senders.
var (
	ChatSenderUser  = ChatSender{v: "user"}
	ChatSenderCappy = ChatSender{v: "cappy"}
	ChatSenderHuman = ChatSender{v: "human"}
)

// String returns the sender key used in CSS class modifiers and data attributes.
func (s ChatSender) String() string {
	return s.v
}

// ToolCallStatus is the execution state of a tool call.
type ToolCallStatus struct{ v string }

// Tool call statuses.
var (
	ToolCallPending = ToolCallStatus{v: "pending"}
	ToolCallRunning = ToolCallStatus{v: "running"}
	ToolCallDone    = ToolCallStatus{v: "done"}
	ToolCallError   = ToolCallStatus{v: "error"}
)

// String returns the status key used in CSS class modifiers.
func (s ToolCallStatus) String() string {
	return s.v
}
