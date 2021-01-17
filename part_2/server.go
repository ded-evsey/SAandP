package main

import (
	"./finder"
	"fmt"
	"log"
	"net/http"
	"html/template"
)

type ViewData struct{
	MapKeys []string
	ErrorFlag bool
}
func main() {
	f := finder.Finder{Homedir: "/home/user/Desktop/Ed/"}  // тестовая директория
	f = f.MakeHash()
	var mapKeys []string
	for key, _ := range f.Hash{
		mapKeys = append(mapKeys, key)
	}
	 var errFlag bool
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			errFlag = f.GetFile(writer,request)
		}
		data := ViewData{
			MapKeys: mapKeys,
			ErrorFlag:errFlag,
		}
		t, _ := template.ParseFiles("/home/user/Desktop/Ed/SAandP/part_2/static/templates/index.html")

		t.Execute(writer, data)
	})
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}