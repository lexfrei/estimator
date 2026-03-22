package main

import (
	"compress/gzip"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

//go:embed static
var staticFiles embed.FS

func main() {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	indexHTMLBytes, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		log.Fatalf("Failed to read index.html: %v", err)
	}

	indexHTML, err := m.String("text/html", string(indexHTMLBytes))
	if err != nil {
		log.Printf("Warning: failed to minify HTML: %v", err)
		indexHTML = string(indexHTMLBytes)
	}

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		etag := "\"" + strconv.FormatInt(int64(len(indexHTML)), 10) + "-" +
			strconv.FormatInt(time.Now().Unix()/86400, 10) + "\""

		if r.Header.Get("If-None-Match") == etag {
			w.Header().Set("ETag", etag)
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("ETag", etag)
		w.Header().Set("Cache-Control", "public, max-age=86400, s-maxage=31536000")
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Expires", time.Now().Add(24*time.Hour).Format(time.RFC1123))
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, indexHTML)
	})

	mux.Handle("GET /i18n/", http.StripPrefix("/i18n/", jsMinHandler(m, staticFS)))

	handler := withGzip(withRecover(withLogger(mux)))

	log.Println("Starting server on :8080...")
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server error: ", err)
	}
}

func jsMinHandler(m *minify.M, staticFS fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(staticFS))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, ".js") {
			fileServer.ServeHTTP(w, r)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/")
		content, err := fs.ReadFile(staticFS, path)
		if err != nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		minified, err := m.Bytes("application/javascript", content)
		if err != nil {
			log.Printf("Warning: failed to minify JS %s: %v", path, err)
			fileServer.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "public, max-age=604800")
		w.Write(minified)
	})
}

func withLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func withRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func withGzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, 5)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")
		next.ServeHTTP(gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}
