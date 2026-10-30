package ui

import (
	"fmt"
	"testing"
)

func TestRenderAvatar(t *testing.T) {
	t.Parallel()

	for i, large := range []bool{false, true} {
		html := renderString(t, Avatar("BL", &AvatarProps{Large: large}))
		writeClassCoverage(t, "avatar", fmt.Sprintf("%03d", i), html)
	}
}
