package main

import (
    "net/http"
    "encoding/json"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

func getCowFileArray(st string) (a []string){
	return
}

func handleCowFile(w http.ResponseWriter, r *http.Request) {
	res, err := cowFileGrpc()
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
	tab := getCowFileArray(res.StderrResponse)
	json.NewEncoder(w).Encode(tab)
}

func handleCow(w http.ResponseWriter, r *http.Request) {
	e := chat.Cow{}
        parseJsonBodyRequest(r, &e)
        res, err :=  CowGrpc(e)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
        json.NewEncoder(w).Encode(res)
}
