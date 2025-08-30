package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	GetRequest(w, r)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	url := originServer + r.URL.Path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}
	println(url)
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	hash := sha256.Sum256(bodyBytes)
	key := r.Method + "#" + r.URL.Path + "#" + r.URL.RawQuery + "#" + hex.EncodeToString(hash[:])

	// if the entry is present in the cache
	if entry, ok := cache[key]; ok {
		if time.Since(entry.CachedAt) < entry.TTL {
			fmt.Println("Cache HIT:", key)

			// restoring headers
			for k, v := range entry.Headers {
				for _, val := range v {
					w.Header().Add(k, val)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(entry.StatusCode)
			w.Write(entry.Body)
			return
		} else {
			fmt.Println("Cache expired:", key)
			delete(cache, key) // remove expired entry
		}
	}

	// forwarding the request to origin server
	req, _ := http.NewRequest(r.Method, url, bytes.NewReader(bodyBytes))
	req.Header = r.Header.Clone()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "error contacting origin", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	// Save to cache
	cache[key] = CacheEntry{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header.Clone(),
		Body:       respBytes,
		CachedAt:   time.Now(),
		TTL:        60 * time.Second, // you can tune per route/request
	}

	// Write response back
	for k, v := range resp.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}
	w.Header().Set("X-Cache", "MISS")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBytes)
	return
}
