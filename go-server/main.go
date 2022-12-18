package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// localhost:8080/form.html to get the form
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[*] coming in -> /form")

		if err := r.ParseForm(); err != nil{
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		fmt.Fprintf(w, "POST request ok\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[*] coming in -> /hello")

		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "hello")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}