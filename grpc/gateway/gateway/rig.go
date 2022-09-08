package main

import (
    "net/http"
    "encoding/json"
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

func handleRig(w http.ResponseWriter, r *http.Request) {
        rigs := rig{}
        parseJsonBodyREquest(r, &rigs)
        res, _ :=  rigGrpc(rigs)
        json.NewEncoder(w).Encode(res)
}
