package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	L "./code/getdata"
	L2 "./code/write"

	"github.com/gorilla/mux"
)

// RsData for result.html
type RsData struct {
	Name         string
	FrontDefault string
	BackDefault  string
	FrontShiny   string
	BackShiny    string
	DataBefore   int
	DataAfter    int
}

//var dataPokemon *L.Response

func homePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	//data := function.getData()
	data := L.GetValue()
	//fmt.Println(data)
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeResult(w http.ResponseWriter, r *http.Request) {
	var input1 = r.FormValue("pokemon")
	data := L.GetPokeData(input1)
	intData, _ := strconv.Atoi(input1)
	prevIn := intData - 1
	nextIn := intData + 1
	dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn}
	b, _ := json.Marshal(dateNew)
	L2.LoggingWrite(string(b))
	var filepath = path.Join("views", "result.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, dateNew)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//fmt.Fprint(data)

}

func main() {
	//var responseObject Response = getDataPok()
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/result", homeResult)
	/*
		http.HandleFunc("/", homePage)
		http.HandleFunc("/result/pokemon={id}", homePage)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
