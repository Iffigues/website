package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/fortune", index).Methods("GET")
    r.HandleFunc("/rig", index).Methods("GET")
    r.HandleFunc("/cow", index).Methods("GET")
    r.HandleFunc("/figlet", index).Methods("GET")
    r.HandleFunc("/toilet", index).Methods("GET")
    buildHandler := http.FileServer(http.Dir("../build"))
    r.PathPrefix("/").Handler(buildHandler)
    srv := &http.Server{
        Handler:      r,
        Addr:         ":80",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())

}

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../build/index.html")
}
