// Copyright (c) 2026 heathcetide. All rights reserved.

package webui

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var amapHostOK = func(host string) bool {
	h := strings.ToLower(host)
	return strings.HasSuffix(h, ".amap.com") ||
		h == "amap.com" ||
		strings.HasSuffix(h, ".autonavi.com") ||
		h == "autonavi.com"
}

// MountAMapProxy mirrors the Vite /_AMapProxy helper so embedded admin can
// load map tiles when the browser blocks direct amap.com access.
// Path: /_AMapProxy/<http|https>/<host>/<path>?query
func MountAMapProxy(r *gin.Engine) {
	client := &http.Client{Timeout: 30 * time.Second}
	r.Any("/_AMapProxy/*proxyPath", func(c *gin.Context) {
		raw := strings.TrimPrefix(c.Param("proxyPath"), "/")
		parts := strings.SplitN(raw, "/", 3)
		if len(parts) < 2 {
			c.String(http.StatusBadRequest, "bad proxy path")
			return
		}
		proto, host := parts[0], parts[1]
		if proto != "http" && proto != "https" {
			c.String(http.StatusBadRequest, "bad proto")
			return
		}
		if !amapHostOK(host) {
			c.String(http.StatusForbidden, "host not allowed")
			return
		}
		rest := ""
		if len(parts) == 3 {
			rest = parts[2]
		}
		target := proto + "://" + host + "/" + rest
		if q := c.Request.URL.RawQuery; q != "" {
			target += "?" + q
		}
		req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, target, c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		for k, vals := range c.Request.Header {
			lk := strings.ToLower(k)
			if lk == "host" || lk == "referer" || lk == "origin" || lk == "accept-encoding" {
				continue
			}
			for _, v := range vals {
				req.Header.Add(k, v)
			}
		}
		req.Header.Set("Referer", proto+"://"+host+"/")
		resp, err := client.Do(req)
		if err != nil {
			c.String(http.StatusBadGateway, err.Error())
			return
		}
		defer resp.Body.Close()
		for k, vals := range resp.Header {
			for _, v := range vals {
				c.Writer.Header().Add(k, v)
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Status(resp.StatusCode)
		_, _ = io.Copy(c.Writer, resp.Body)
	})
}
