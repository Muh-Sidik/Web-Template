/* بِسْمِ اللَّهِ الرَّحْمَنِ الرَّحِيم */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Golang Web Server</h1>")
}

func index(w http.ResponseWriter, r *http.Request)  {
	var data = map[string]string{
		"Name": "Tri Apria Ningsih",
		"Message": "Halo Kamu hehehe",
	}

	var template, err = template.ParseFiles("tempalate.html")
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	template.Execute(w, data)
}

func webRoutes()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/index", index)
	fmt.Println("Starting Web Server at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main()  {
	webRoutes()
}