package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	L "./code"

	"github.com/gorilla/mux"
)

// RsData for result.html
type RsData struct {
	Name         string
	FrontDefault string
	BackDefault  string
	FrontShiny   string
	BackShiny    string
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
	data := L.GetPokeData(r.FormValue("pokemon"))
	dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny}
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
	r.HandleFunc("/result2", homeResult)
	/*
		http.HandleFunc("/", homePage)
		http.HandleFunc("/result/pokemon={id}", homePage)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
