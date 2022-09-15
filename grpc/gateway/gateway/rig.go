package main

import (
    "net/http"
    "encoding/json"
	"strings"
	"fmt"
)


type rig struct {
        Man   bool `json:"man"`
        Woman bool `json:"woman"`
        Nbr   string `json:"nbr"`
}

type rigs struct {
        Man   bool `json:"man"`
        Woman bool `json:"woman"`
        Nbr   int `json:"nbr"`
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
        rigs := rig{}
        parseJsonBodyREquest(r, &rigs)
        res, err :=  rigGrpc(rigs)
	data := rigToStringArray(res.StdoutResponse)
	fmt.Println(data, res)
	if err != nil {
		return
	}
        json.NewEncoder(w).Encode(res)
}
