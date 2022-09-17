package main

import (
    "net/http"
    "encoding/json"
)


type fortune struct {
        Man   bool `json:"man"`
        Woman bool `json:"woman"`
        Nbr   string `json:"nbr"`
}

type FortuneStruct struct {
	Name string
	Addr	string
	Postal string
	Tel string
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
        fortunes := fortune{}
        parseJsonBodyREquest(r, &fortunes)
        res, err :=  fortuneGrpc(fortunes)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
        json.NewEncoder(w).Encode(fortunes)
}
