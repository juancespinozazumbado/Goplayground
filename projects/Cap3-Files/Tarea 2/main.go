package main

import (
	"animals/api"
	"animals/data"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	db := data.NewDB()
	api.DB = db

	http.HandleFunc("/animals", api.GetAnimalsHandler)
	http.HandleFunc("/process", api.ProcessAnimalHandler)
	http.HandleFunc("/describe", api.DescribeAnimalHandler)

	//other metods

	http.HandleFunc("/get", api.GetAnimalByIdHandler)
	http.HandleFunc("/delete", api.DeleteAnimalByIdHandler)

	log.Info("Server listening on...:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
