package main

import (
    "net/http"
    "encoding/json"
    "strings"
    chat "github.com/Iffigues/website/proto/linuxCommand"
   "fmt"
)

func getFileFortuneArray(st string) (a []string){
	dd := strings.Trim(string(st), "100.00% /usr/share/games/fortunes")
	d := strings.Split(dd,"\n")
	d = d[1:len(d)-1]
	for k,v := range d {
		d[k] = strings.TrimSpace(v)
		e := strings.Split(d[k], " ")
		d[k] = e[1]
	}
	return d
}

func handleFortuneFile(w http.ResponseWriter, r *http.Request) {
	res, err := fortuneFileGrpc()
	fmt.Println(res, err)
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
	tab := getFileFortuneArray(res.StderrResponse)
	json.NewEncoder(w).Encode(tab)
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
	e := chat.Fortune{}
        parseJsonBodyRequest(r, &e)
        res, err :=  fortuneGrpc(e)
	fmt.Println("haha=",res, err, "rr")
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
        json.NewEncoder(w).Encode(res)
}
