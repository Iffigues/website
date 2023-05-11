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

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func main() {
	router := mux.NewRouter()

	url := []Url{
		Url{"/rig", handleRig, []string{"POST"}},
		Url{"/fortune", handleFortune, []string{"POST"}},
		Url{"/fortunefile", handleFortuneFile, []string{"POST"}},
		Url{"/cowfile", handleCowFile, []string{"POST"}},
		Url{"/cow", handleCow, []string{"POST"}},
		Url{"/figletfile", handleFigletFile, []string{"POST"}},
		Url{"/figlet", handleFiglet, []string{"POST"}},
		Url{"/toiletfile", handleToiletFile, []string{"POST"}},
		Url{"/toilet", handleToilet, []string{"POST"}},
		Url{"/xkcdpass", handleXkcdpass, []string{"POST"}},
	}

	for _, val := range url {
		endpoint := router.HandleFunc(val.Url, val.Handle)
		for _, method := range val.Method {
			endpoint.Methods(method)
		}
	}

	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router.Use(jsonHeader)
	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(corsObj, headersOk, methodsOk)(router)))
}
