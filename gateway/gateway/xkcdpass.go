package main

import (
	"fmt"
    "net/http"
    "encoding/json"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

func handleXkcdpass(w http.ResponseWriter, r *http.Request) {
	println("ezezze")
	e := chat.Xkcdpass{}
        parseJsonBodyRequest(r, &e)
        res, err :=  xkcdpassGrpc(e)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(res)
}
