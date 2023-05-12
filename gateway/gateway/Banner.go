package main

import (
    "net/http"
    "encoding/json"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

func handleBanner(w http.ResponseWriter, r *http.Request) {
	e := chat.Banner{}
        parseJsonBodyRequest(r, &e)
        res, err :=  bannerGrpc(e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(res)
}
