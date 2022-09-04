package main

import (
	"fmt"
    "encoding/json"
    "net/http"
)


func parseJsonBodyREquest(r *http.Request, data interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	defer r.Body.Close()
	return
}
