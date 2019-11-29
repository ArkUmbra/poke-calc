package main

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)


func homePage(writer http.ResponseWriter, r *http.Request) {
	writeFileToHtml(writer, "web/html/index.html")
}

func writeFileToHtml(writer http.ResponseWriter, filename string) {
	html, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	t := template.New("page")

	t, err2 := t.Parse(string(html))
	if err2 != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	t.Execute(os.Stdout,nil)
	t.Execute(writer, nil)
}

func startServer() {
	//box := rice.MustFindBox("web")
	cssFileServer := http.StripPrefix("/web/", http.FileServer(box.HTTPBox()))
	http.Handle("/web/", cssFileServer)


	http.HandleFunc("/", homePage)

	port := "8080"
	setPort := os.Getenv("PORT")
	if setPort != "" {
		port = setPort
	}

	error := http.ListenAndServe(":" + port, nil)
	fmt.Println(error)
}

func main() {
	startServer()
}