package main

import (
	"fmt"
    "strings"
    "net/http"
    "encoding/json"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

func getCowFileArray(st string) (a []string){
	ff := strings.Split(strings.TrimSpace(strings.Trim(st, "Cow files in /usr/share/cowsay/cows:")), " ")
	for _, gg := range ff {
		ffg := strings.Split(gg, "\n")
		for _, ez := range ffg {
			a = append(a, ez)
		}
	}
	return
}

func handleCowFile(w http.ResponseWriter, r *http.Request) {
	res, err := cowFileGrpc()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tab := getCowFileArray(res.StdoutResponse)
	json.NewEncoder(w).Encode(tab)
}

func handleCow(w http.ResponseWriter, r *http.Request) {
	e := chat.Cow{}
        parseJsonBodyRequest(r, &e)
        res, err :=  cowGrpc(e)
	fmt.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(res)
}
