package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/yardrail/ui"
)

// testCSS is the stylesheet the test gallery serves.
const testCSS = ".yr-button { color: red; }\n"

func TestGalleryIndex(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	status, _ := get(t, handler, "/")
	if status != http.StatusOK {
		t.Fatalf("status = %d, want %d", status, http.StatusOK)
	}

	// Every example must be reachable as a standalone page.
	for _, group := range ui.ExampleGroups() {
		for _, ex := range group.Examples {
			path := "/" + group.Component + "/" + ex.Name

			status, _ := get(t, handler, path)
			if status != http.StatusOK {
				t.Errorf("GET %s = %d, want %d", path, status, http.StatusOK)
			}
		}
	}
}

func TestGalleryExample(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	status, body := get(t, handler, "/button/primary")
	if status != http.StatusOK {
		t.Fatalf("status = %d, want %d", status, http.StatusOK)
	}

	for _, want := range []string{"<button", `href="/ui.css"`} {
		if !strings.Contains(body, want) {
			t.Errorf("page is missing %s:\n%s", want, body)
		}
	}
}

func TestGalleryUnknownExample(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	status, _ := get(t, handler, "/button/missing")
	if status != http.StatusNotFound {
		t.Errorf("status = %d, want %d", status, http.StatusNotFound)
	}
}

func TestGalleryCSS(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	status, body := get(t, handler, "/ui.css")
	if status != http.StatusOK {
		t.Fatalf("status = %d, want %d", status, http.StatusOK)
	}

	if body != testCSS {
		t.Errorf("body = %q, want %q", body, testCSS)
	}
}

// newTestHandler returns a gallery handler whose assets hold a ui.css of testCSS.
func newTestHandler(t *testing.T) http.Handler {
	t.Helper()

	return newHandler(fstest.MapFS{"ui.css": {Data: []byte(testCSS)}})
}

// get sends a GET for path to handler and returns the status and body.
func get(t *testing.T, handler http.Handler, path string) (int, string) {
	t.Helper()

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, path, nil))

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	return rec.Code, string(body)
}
