package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	handler := func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "Hellow world")
	}

	handlerParam := func(w http.ResponseWriter, request *http.Request) {

		name := request.URL.Query().Get("Name")
		fmt.Fprintf(w, "Hellow %s", name)
	}

	http.HandleFunc("/hellow", handler)
	http.HandleFunc("/param", handlerParam)

	fmt.Print("Server listening on...:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
