package main

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
    chat "github.com/Iffigues/website/proto/linuxCommand"
)

type Toilet struct {
	F []string
	FF []string
	E []string
}


func getFileFToiletArray(st string) (a []string){
	ll := strings.Split(string(st), "Figlet control files in this directory:")
	a = strings.Split(ll[0], "\n")[3:]
	a[len(a) - 1] = "mono12"
	return
}

func getFileEToiletArray(st string) (a []string){
	ff := strings.Split(string(st), "\n")
	for _, val :=  range ff[1:len(ff) - 1] {
		r := strings.Split(val, "\"")
		a = append(a, r[1])
	}
	return
}

func handleToiletFFile() (tab []string) {
	res, err := toiletFFileGrpc()
	if err != nil {
		return
	}
	tab = getFileEToiletArray(res.StdoutResponse)
	return
}

func handleToiletEFile() (tab []string) {
	res, err := toiletEFileGrpc()
	if err != nil {
		return
	}
	tab = getFileEToiletArray(res.StdoutResponse)
	return
}

func handleToiletFFFile() (tab []string) {
	res, err := toiletFFFileGrpc()
	if err != nil {
		return
	}
	return getFileFToiletArray(res.StdoutResponse)
}


func handleToiletFile(w http.ResponseWriter, r *http.Request) {
	ee := Toilet{}
	ee.F = handleToiletFFile()
	ee.FF = handleToiletFFFile()
	ee.E = handleToiletEFile()
	json.NewEncoder(w).Encode(ee)
}

func handleToilet(w http.ResponseWriter, r *http.Request) {
	e := chat.Toilet{}
        parseJsonBodyRequest(r, &e)
        fmt.Println(e,e.F3)
	res, err :=  toiletGrpc(e)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
        json.NewEncoder(w).Encode(res)
}
