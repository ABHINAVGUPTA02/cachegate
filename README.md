# 🌀 CacheGate

**CacheGate** is a lightweight command-line tool (CLI) built with Go that lets you start a caching proxy server for any origin server.  
It helps in testing API responses with caching enabled, and displays cache results (`HIT` / `MISS`) directly in the terminal.

---

## ✨ Features
- 🔄 Reverse proxy for forwarding requests
- ⚡ In-memory caching for **GET** requests
- ⏱ Configurable cache TTL (Time To Live)
- 📝 Custom response headers:
  - `X-Cache: HIT` → served from cache
  - `X-Cache: MISS` → fetched from origin server
- 🧹 Ability to clear the cache
---

## 🚀 Commands

### Start Proxy Server
```bash
caching-proxy --port <NUMBER> --origin <URL>
```

Starts the proxy server on the given port, forwarding requests to the specified origin server.

### Clear Cache
```bash
caching-proxy --clear-cache
```

Clears the cache of the proxy server

### Exit
```bash
exit
```

Exits the program

Project URL: https://github.com/ABHINAVGUPTA02/cachegate/tree/master
