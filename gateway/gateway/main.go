package main

import (
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func jsonHeader(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
	    next.ServeHTTP(w, r)
    })
}

func main() {
    router := mux.NewRouter()
	router.HandleFunc("/rig", handleRig).Methods("POST")
	corsObj:=handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router.Use(jsonHeader)

    log.Fatal(http.ListenAndServe(":9090", handlers.CORS(corsObj, headersOk, methodsOk)(router)))
}