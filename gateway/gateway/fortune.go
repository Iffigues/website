package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)


type fortune struct {
        A   bool `json:"A"`
        C bool `json:"C"`
	E   bool `json:"E"`
	F   bool `json:"F"`
	L   bool `json:"L"`
	O   bool `json:"O"`
	S   bool `json:"S"`
	I   bool `json:"I"`
	U   bool `json:"U"`
        Nbr   string `json:"nbr"`
}

type FortuneStructs struct {
	Name string
	Addr	string
	Postal string
	Tel string
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
        fortunes := fortune{}
        parseJsonBodyREquest(r, &fortunes)
        res, err :=  fortuneGrpc(fortunes)
	fmt.Println(res)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
        json.NewEncoder(w).Encode(res)
}
