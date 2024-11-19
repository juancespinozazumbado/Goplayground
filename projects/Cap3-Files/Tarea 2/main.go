package main

import (
	"animals/api"
	"animals/data"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	//add memory DB
	db := data.NewDB()
	api.DB = db

	//add main handlers
	http.HandleFunc("/animals", api.GetAnimalsHandler)      //<host>:8080/animals | <host>:8080/animals?type=dog
	http.HandleFunc("/process", api.ProcessAnimalHandler)   //<host>:8080/process?type=dog&name=name&attr=atribute
	http.HandleFunc("/describe", api.DescribeAnimalHandler) //<host>:8080/describe?type=dog&name=name&attr=atribute

	//other metods

	http.HandleFunc("/get", api.GetAnimalByIdHandler)       //<host>:8080/get?id={id}
	http.HandleFunc("/delete", api.DeleteAnimalByIdHandler) //<host>:8080/delete?id={id}

	log.Info("Server listening on...:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
