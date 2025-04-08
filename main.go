package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
)

//go:embed static
var staticFiles embed.FS

func main() {
	// Initialize Echo
	e := echo.New()
	e.HideBanner = true

	// Initialize minifier
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Read and minify index.html content
	indexHTMLBytes, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		log.Fatalf("Failed to read index.html: %v", err)
	}

	// Minify HTML
	indexHTML, err := m.String("text/html", string(indexHTMLBytes))
	if err != nil {
		log.Printf("Warning: failed to minify HTML: %v", err)
		indexHTML = string(indexHTMLBytes)
	}

	// Handler for index.html
	e.GET("/", func(c echo.Context) error {
		// Generate ETag based on content
		etag := "\"" + strconv.FormatInt(int64(len(indexHTML)), 10) + "-" + 
			strconv.FormatInt(time.Now().Unix()/86400, 10) + "\""
		
		// Handle conditional requests (If-None-Match)
		ifNoneMatch := c.Request().Header.Get("If-None-Match")
		if ifNoneMatch == etag {
			c.Response().Header().Set("ETag", etag)
			return c.NoContent(http.StatusNotModified)
		}

		// Set caching headers
		c.Response().Header().Set("Content-Type", echo.MIMETextHTMLCharsetUTF8)
		c.Response().Header().Set("ETag", etag)
		c.Response().Header().Set("Cache-Control", "public, max-age=86400, s-maxage=31536000")
		c.Response().Header().Set("Vary", "Accept-Encoding")
		
		expires := time.Now().Add(time.Hour * 24).Format(time.RFC1123)
		c.Response().Header().Set("Expires", expires)

		return c.HTML(http.StatusOK, indexHTML)
	})

	// Middleware for JS file minification
	jsMinMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Is this a JavaScript file?
			if strings.HasSuffix(c.Request().URL.Path, ".js") {
				// Get file by path
				path := strings.TrimPrefix(c.Request().URL.Path, "/")
				content, err := staticFiles.ReadFile("static/" + path)
				if err != nil {
					return next(c)
				}

				// Minify JavaScript
				minified, err := m.Bytes("application/javascript", content)
				if err != nil {
					log.Printf("Warning: failed to minify JS %s: %v", path, err)
					return next(c)
				}

				// Return minified JavaScript with caching headers
				c.Response().Header().Set("Content-Type", "application/javascript")
				c.Response().Header().Set("Cache-Control", "public, max-age=604800") // Cache for 7 days
				return c.Blob(http.StatusOK, "application/javascript", minified)
			}
			return next(c)
		}
	}

	// Setup static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}

	// Handler for i18n static files with minification
	e.GET("/i18n/*", jsMinMiddleware(echo.WrapHandler(http.StripPrefix("/i18n/", http.FileServer(http.FS(staticFS))))))

	// Start server
	log.Println("Starting server on :8080...")
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal("Server error: ", err)
	}
}
