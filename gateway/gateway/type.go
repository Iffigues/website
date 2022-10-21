package main

import "net/http"

type Url struct {
	Url string
	Handle func(w http.ResponseWriter, r *http.Request)
	Method []string
}
