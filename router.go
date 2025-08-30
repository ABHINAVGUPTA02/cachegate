package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/*", GetHandler)

	return r
}
