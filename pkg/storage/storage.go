// Copyright (c) 2026 heathcetide. All rights reserved.

// Package storage wraps ling-base/stores object storage as a package-level singleton.
package storage

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/heathcetide/infectious-disease/pkg/stores"
	"github.com/heathcetide/infectious-disease/pkg/stores/kodo"
	"github.com/heathcetide/infectious-disease/pkg/stores/local"
)

// Package-level singleton.
var (
	// Store is the active object storage instance. nil until Init() is called.
	Store stores.Store

	// MaxFileSize is the maximum file size in bytes (0 = unlimited).
	MaxFileSize int64

	// PublicURLBase is the public access base URL (empty = use Store.PublicURL() default).
	PublicURLBase string

	// Private is true when the active backend requires signed read URLs.
	Private bool
)

// InitOptions carries cloud-backend options beyond the basic Init args.
type InitOptions struct {
	Private bool // private bucket (kodo signed URLs)
}

// Init initializes the object storage singleton from configuration.
func Init(driver, root, region, endpoint, accessKey, secretKey, publicURLBase string, maxFileSize int64) error {
	return InitWithOptions(driver, root, region, endpoint, accessKey, secretKey, publicURLBase, maxFileSize, InitOptions{})
}

// InitWithOptions initializes storage with extra options (e.g. private bucket).
func InitWithOptions(driver, root, region, endpoint, accessKey, secretKey, publicURLBase string, maxFileSize int64, opts InitOptions) error {
	switch strings.ToLower(strings.TrimSpace(driver)) {
	case "local", "":
		Store = local.New(local.Config{
			Root:       root,
			NewDirPerm: 0755,
		})
	case "kodo", "qiniu":
		domain := publicURLBase
		if domain == "" {
			domain = endpoint
		}
		Store = kodo.New(kodo.Config{
			AccessKey:  accessKey,
			SecretKey:  secretKey,
			BucketName: root,
			Domain:     domain,
			Private:    opts.Private,
			Region:     region,
		})
	default:
		return fmt.Errorf("unsupported storage driver: %s (go get the corresponding module first)", driver)
	}

	MaxFileSize = maxFileSize
	PublicURLBase = strings.TrimRight(publicURLBase, "/")
	Private = opts.Private
	return nil
}

// Write stores the content from reader to the given key.
func Write(key string, r io.Reader) error {
	if Store == nil {
		return fmt.Errorf("storage not initialized, call Init() first")
	}
	return Store.Write(key, r)
}

// Read returns a reader and size for the object at the given key.
func Read(key string) (io.ReadCloser, int64, error) {
	if Store == nil {
		return nil, 0, fmt.Errorf("storage not initialized, call Init() first")
	}
	return Store.Read(key)
}

// Delete removes the object at the given key (no error if it does not exist).
func Delete(key string) error {
	if Store == nil {
		return fmt.Errorf("storage not initialized, call Init() first")
	}
	return Store.Delete(key)
}

// Exists checks whether the object at the given key exists.
func Exists(key string) (bool, error) {
	if Store == nil {
		return false, fmt.Errorf("storage not initialized, call Init() first")
	}
	return Store.Exists(key)
}

// PublicURL returns the public access URL for the given key.
// For private kodo buckets Prefer SignedURL / Store.PublicURL (already signed).
// When PublicURLBase is set and store is local, returns PublicURLBase/key.
func PublicURL(key string) string {
	if Store == nil {
		return ""
	}
	key = strings.TrimLeft(key, "/")
	// Prefer backend PublicURL (kodo signs private objects).
	if u := Store.PublicURL(key); u != "" {
		return u
	}
	if PublicURLBase != "" {
		return PublicURLBase + "/" + key
	}
	return ""
}

// StableURL returns an unsigned CDN URL (domain/key) suitable for DB persistence.
// Private buckets still need SignedURL at read time.
func StableURL(key string) string {
	key = strings.TrimLeft(key, "/")
	if PublicURLBase == "" {
		return PublicURL(key)
	}
	return PublicURLBase + "/" + key
}

// SignedURL returns a time-limited URL for private objects.
func SignedURL(key string, expires time.Duration) (string, error) {
	if Store == nil {
		return "", fmt.Errorf("storage not initialized, call Init() first")
	}
	return stores.SignedURL(Store, key, expires)
}

// CheckFileSize validates that the file size is within the allowed limit.
func CheckFileSize(size int64) error {
	if MaxFileSize > 0 && size > MaxFileSize {
		return fmt.Errorf("file size %d exceeds limit %d", size, MaxFileSize)
	}
	return nil
}

// Close releases resources.
func Close() error { return nil }
