package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	L3 "./code/database"
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

//ReturnData for validation data
func ReturnData(validation bool, input1 string) RsData {
	if validation == true {
		data := L.GetPokeData(input1)
		intData, _ := strconv.Atoi(input1)
		prevIn := intData - 1
		nextIn := intData + 1
		dateNew := RsData{data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn}
		L3.InsData(data.Name, data.Sprites.FrontDefault, data.Sprites.BackDefault, data.Sprites.FrontShiny, data.Sprites.BackShiny, prevIn, nextIn)
		return dateNew
	}
	_, result2, result3, result4, result5, result6, result7, result8 := L3.GetData(input1)
	dateNew := RsData{result2, result3, result4, result5, result6, result7, result8}
	return dateNew

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
	var validation1 bool
	var input1 = r.FormValue("pokemon")
	validation1 = L3.ValidationData(input1)
	fmt.Printf("The result validation is: %v\n", validation1)
	dateNew := ReturnData(validation1, input1)
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
	L3.GetSysDate()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/result", homeResult)
	/*
		http.HandleFunc("/", homePage)
		http.HandleFunc("/result/pokemon={id}", homePage)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	*/
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("hdmonochrome"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
