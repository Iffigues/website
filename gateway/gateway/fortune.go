package main

import (
    "net/http"
    "encoding/json"
    "strings"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)


type Percent struct {
	Name string `json:"name"`
	Percent string `json"percent"`
}

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
	M   string `json:"M"`
        N   string `json:"N"`
	Percent []Percent `json:"Percent"`
}


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

func handleFileFortune(w http.ResponseWriter, r *http.Request) {
	res, err := fortuneFileGrpc()
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
	if err != nil {
		json.NewEncoder(w).Encode(res)
		return
	}
        json.NewEncoder(w).Encode(res)
}
