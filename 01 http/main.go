package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData = IndexData{
		Title: "go study",
		Desc:  "入门",
	}

	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)

}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData = IndexData{
		Title: "go study",
		Desc:  "入门",
	}
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, err := t.ParseFiles(path + "/template/index.html")
	if err != nil {
		log.Writer()
	}
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
