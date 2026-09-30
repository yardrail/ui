// Command gallery serves every ui package example on a local port, styled with the built library
// CSS: an index at / and one page per example at /<component>/<name>. Run it with task gallery.
package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/yardrail/ui"
)

// readHeaderTimeout bounds how long a client may take to send request headers.
const readHeaderTimeout = 5 * time.Second

func main() {
	addr := flag.String("addr", "localhost:6007", "address to listen on")

	flag.Parse()

	srv := &http.Server{
		Addr:              *addr,
		Handler:           newHandler(ui.Assets),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	slog.Info("serving ui gallery", "url", "http://"+*addr)

	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("gallery server stopped", "err", err)
		os.Exit(1)
	}
}
