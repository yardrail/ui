package ui

import "testing"

func TestRenderDownloadList(t *testing.T) {
	t.Parallel()

	html := renderString(t, withChildren(
		DownloadList(&DownloadListProps{}),
		text("<a>file.zip</a>"),
	))
	writeClassCoverage(t, "download_list", "default", html)
}
