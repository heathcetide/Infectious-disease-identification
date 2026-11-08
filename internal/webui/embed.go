// Copyright (c) 2026 heathcetide. All rights reserved.

// Package webui embeds the built admin SPA and serves it from the API process
// so local/dev only needs the Go server (+ mini-program), not a separate Vite.
package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Dist is the Vite production build output (admin → internal/webui/dist).
// Rebuild with: ./scripts/build-admin.sh
//
//go:embed all:dist
var Dist embed.FS

// Mount registers static admin UI routes on the Gin engine.
// Assets are public (no JWT). SPA fallback serves index.html for non-API GETs.
func Mount(r *gin.Engine) error {
	sub, err := fs.Sub(Dist, "dist")
	if err != nil {
		return err
	}
	fileServer := http.FileServer(http.FS(sub))

	// Hashed assets / favicon / logo under /
	r.GET("/assets/*filepath", gin.WrapH(fileServer))
	r.GET("/logo.png", gin.WrapH(fileServer))
	r.GET("/vite.svg", gin.WrapH(fileServer))

	serveIndex := func(c *gin.Context) {
		data, err := Dist.ReadFile("dist/index.html")
		if err != nil {
			c.String(http.StatusServiceUnavailable, "admin UI not built; run scripts/build-admin.sh")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	}

	r.GET("/", serveIndex)
	r.GET("/index.html", serveIndex)

	// Hash-router SPA: only fall through for browser document navigations that
	// are not API/docs/health. (Hash routes usually hit "/" anyway.)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			c.Status(http.StatusNotFound)
			return
		}
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api/") ||
			strings.HasPrefix(path, "/docs") ||
			strings.HasPrefix(path, "/openapi") ||
			strings.HasPrefix(path, "/cdc") ||
			strings.HasPrefix(path, "/medical") ||
			strings.HasPrefix(path, "/user/") ||
			strings.HasPrefix(path, "/ai/") ||
			strings.HasPrefix(path, "/_AMapProxy/") ||
			path == "/health" || path == "/live" || path == "/ready" {
			c.Status(http.StatusNotFound)
			return
		}
		serveIndex(c)
	})
	return nil
}
