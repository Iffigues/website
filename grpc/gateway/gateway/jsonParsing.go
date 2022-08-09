package main

import (
    "encoding/json"
    "net/http"
)


func parseJsonBodyREquest(r *http.Request, data interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		println(err)
	}
	defer r.Body.Close()
	return
}
