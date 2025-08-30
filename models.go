package main

import (
	"fmt"
	"time"
)

type CacheEntry struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
	CachedAt   time.Time
	TTL        time.Duration
}

// The global cache (map of key → cached response)
var cache = make(map[string]CacheEntry)

// Clear the whole cache
func clearCache() {
	cache = make(map[string]CacheEntry)
	fmt.Println("cleared....")
}
