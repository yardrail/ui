package ui

import "testing"

func TestAllMethodsReturnValues(t *testing.T) {
	t.Parallel()

	check := func(name string, n int) {
		t.Helper()

		if n == 0 {
			t.Errorf("%s.All() returned 0 values", name)
		}
	}

	check("Tone", len(Tone{}.All()))
	check("ButtonType", len(ButtonType{}.All()))
	check("ButtonVariant", len(ButtonVariant{}.All()))
	check("CappyVariant", len(CappyVariant{}.All()))
	check("ChatSender", len(ChatSender{}.All()))
	check("ToolCallStatus", len(ToolCallStatus{}.All()))
	check("ChatFabBadge", len(ChatFabBadge{}.All()))
	check("SearchVariant", len(SearchVariant{}.All()))
	check("FieldType", len(FieldType{}.All()))
	check("Autocomplete", len(Autocomplete{}.All()))
	check("Space", len(Space{}.All()))
	check("Measure", len(Measure{}.All()))
	check("SideWidth", len(SideWidth{}.All()))
	check("ItemWidth", len(ItemWidth{}.All()))
	check("Align", len(Align{}.All()))
	check("Justify", len(Justify{}.All()))
	check("As", len(As{}.All()))
	check("ShellWidth", len(ShellWidth{}.All()))
}
