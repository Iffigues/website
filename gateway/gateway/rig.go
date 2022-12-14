package main

import (
    "net/http"
    "encoding/json"
    "strings"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)


type rig struct {
        Man   bool `json:"man"`
        Woman bool `json:"woman"`
        Nbr   string `json:"nbr"`
}

type RigStruct struct {
	Name string
	Addr	string
	Postal string
	Tel string
}

func rigToStringArray(e string) (r []RigStruct) {
	for _, val := range  strings.Split(strings.TrimSpace(e), "\n\n") {
		var ee RigStruct
		vv :=  strings.Split(strings.TrimSpace(val), "\n")
		ee.Name, ee.Addr, ee.Postal, ee.Tel = vv[0], vv[1], vv[2], vv[3]
		r = append(r, ee)
	}
	return
}

func handleRig(w http.ResponseWriter, r *http.Request) {
        rigs := chat.Rig{}
        parseJsonBodyRequest(r, &rigs)
        res, err :=  rigGrpc(rigs)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	data := rigToStringArray(res.StdoutResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(data)
}
