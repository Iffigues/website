package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

type FileFiglet struct {
	E []string
	F []string
}

func getFileFigletArray(st string) (a []string, b []string){
	ll := strings.Split(string(st), "Figlet control files in this directory:")
	a = strings.Split(ll[0], "\n")[3:]
	a[len(a) - 1] = "mono12"
	b = strings.Split(ll[1], "\n")[1:]
	return
}

func handleFigletFile(w http.ResponseWriter, r *http.Request) {
	res, err := figletFileGrpc()
	if err != nil {
		 http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	ee :=  FileFiglet{}
	tab, tabf := getFileFigletArray(res.StdoutResponse)
	ee.F, ee.E = tab, tabf
	json.NewEncoder(w).Encode(ee)
}

func handleFiglet(w http.ResponseWriter, r *http.Request) {
	e := chat.Figlet{}
        parseJsonBodyRequest(r, &e)
        res, err :=  figletGrpc(e)
	fmt.Println(e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(res)
}
