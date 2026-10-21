package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	ui "github.com/yardrail/ui/pkg"
)

// testCSS is the stylesheet the test gallery serves.
const testCSS = ".yr-button { color: red; }\n"

// Paths the tests request.
const (
	pathButtonPrimary = "/button/primary"
	pathShellAppPage  = "/shell/app-page"
	pathStaticUICSS   = "/static/ui/ui.css"
)

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

	status, body := get(t, handler, pathButtonPrimary)
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

func TestGalleryStylesheet(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	_, index := get(t, handler, "/")
	if !strings.Contains(index, `href="/gallery.css"`) {
		t.Errorf("index does not link /gallery.css:\n%s", index)
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/gallery.css", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if got := rec.Header().Get("Content-Type"); !strings.HasPrefix(got, "text/css") {
		t.Errorf("Content-Type = %q, want text/css", got)
	}

	if !strings.Contains(rec.Body.String(), ".browser-frame") {
		t.Errorf("stylesheet is missing .browser-frame:\n%s", rec.Body.String())
	}
}

func TestGalleryMocksStylesheet(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	_, mock := get(t, handler, "/msp-dashboard/default")
	if !strings.Contains(mock, `href="/mocks.css"`) {
		t.Errorf("mock does not link /mocks.css:\n%s", mock)
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/mocks.css", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if got := rec.Header().Get("Content-Type"); !strings.HasPrefix(got, "text/css") {
		t.Errorf("Content-Type = %q, want text/css", got)
	}

	if !strings.Contains(rec.Body.String(), ".mock-stat-card") {
		t.Errorf("stylesheet is missing .mock-stat-card:\n%s", rec.Body.String())
	}
}

func TestGalleryShellExample(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	status, body := get(t, handler, pathShellAppPage)
	if status != http.StatusOK {
		t.Fatalf("GET %s = %d, want %d", pathShellAppPage, status, http.StatusOK)
	}

	if !strings.HasPrefix(body, "<!doctype html>") {
		t.Errorf("body does not start with <!doctype html>:\n%s", body)
	}

	if n := strings.Count(body, "<html"); n != 1 {
		t.Errorf("body holds %d <html elements, want 1:\n%s", n, body)
	}

	_, wrapped := get(t, handler, pathButtonPrimary)
	if !strings.Contains(wrapped, `href="/ui.css"`) {
		t.Errorf("GET %s is not wrapped by the example page:\n%s", pathButtonPrimary, wrapped)
	}
}

func TestGalleryStaticAssets(t *testing.T) {
	t.Parallel()

	handler := newTestHandler(t)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequestWithContext(t.Context(), http.MethodGet, pathStaticUICSS, nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("GET %s = %d, want %d", pathStaticUICSS, rec.Code, http.StatusOK)
	}

	if got := rec.Header().Get("Content-Type"); !strings.HasPrefix(got, "text/css") {
		t.Errorf("Content-Type = %q, want text/css", got)
	}

	if rec.Body.String() != testCSS {
		t.Errorf("body = %q, want %q", rec.Body.String(), testCSS)
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
