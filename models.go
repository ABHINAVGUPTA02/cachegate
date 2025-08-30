package main

import (
	"time"
)

type CacheEntry struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
	CachedAt   time.Time
	TTL        time.Duration
}

var cache = make(map[string]CacheEntry)

func clearCache() {
	cache = make(map[string]CacheEntry)
}
