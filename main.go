package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

func homePage(writer http.ResponseWriter, r *http.Request) {
	writeFileToHtml(writer, "html/index.html")
	////fmt.Fprintf(writer, "Hello home page")
	//
	//html, err := ioutil.ReadFile("html/index.html")
	//if err != nil {
	//	panic(err)
	//}
	//
	//t := template.New("table")
	//
	////t, err := t.Parse(AddForm)
	//t, err2 := t.Parse(string(html))
	//if err2 != nil {
	//	panic(err)
	//}
	//
	//writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	////fmt.Fprint(w, AddForm)
	//
	//t.Execute(os.Stdout,nil)
	//t.Execute(writer, nil)
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