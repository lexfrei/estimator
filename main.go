package main

import (
	_ "embed"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/index.html
var indexHTML string

func main() {
	// Initialize Echo
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Handler для index.html
	e.GET("/", func(c echo.Context) error {
		// Генерируем ETag на основе содержимого
		etag := "\"" + strconv.FormatInt(int64(len(indexHTML)), 10) + "-" + 
			strconv.FormatInt(time.Now().Unix()/86400, 10) + "\""
		
		// Обрабатываем условные запросы (If-None-Match)
		ifNoneMatch := c.Request().Header.Get("If-None-Match")
		if ifNoneMatch == etag {
			c.Response().Header().Set("ETag", etag)
			return c.NoContent(http.StatusNotModified)
		}

		// Устанавливаем заголовки для кеширования
		c.Response().Header().Set("Content-Type", echo.MIMETextHTMLCharsetUTF8)
		c.Response().Header().Set("ETag", etag)
		c.Response().Header().Set("Cache-Control", "public, max-age=86400, s-maxage=31536000")
		c.Response().Header().Set("Vary", "Accept-Encoding")
		
		expires := time.Now().Add(time.Hour * 24).Format(time.RFC1123)
		c.Response().Header().Set("Expires", expires)

		return c.HTML(http.StatusOK, indexHTML)
	})

	// Запуск сервера
	log.Println("Starting server on :8080...")
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal("Server error: ", err)
	}
}
