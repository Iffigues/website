package main

import (
	"net/http"
)

type Handler struct {
	H      func(w http.ResponseWriter, r *http.Request)
	Method []string
	Val    []int
}

type Data struct {
	Handler  map[string]Handler
}
