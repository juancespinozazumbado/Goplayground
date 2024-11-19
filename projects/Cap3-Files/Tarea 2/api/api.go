package api

import (
	"animals/animal"
	"animals/animal/bird"
	"animals/animal/cat"
	"animals/animal/dog"
	"animals/data"
	"fmt"
	"net/http"
)

var DB *data.MemoryDb

func formatAnimalinfo(animals []animal.Animal) string {
	info := "{\n"
	if len(animals) == 0 {
		info += `"data": []`

	} else {

		info += `"data": [` + "\n"
		for i, current := range animals {
			info += "  " + animal.GetAnimalInfo(current)
			if i < (len(animals) - 1) {
				info += ",\n"
			} else {
				info += "\n"
			}

		}
		info += "]\n"

	}
	info += "}"
	return info
}

func GetAnimalsHandler(writer http.ResponseWriter, request *http.Request) {

	// set animals data
	animals := DB.GetAnimals()
	info := ""

	animalType := request.URL.Query().Get("type")
	if animalType != "" {
		if animalType != "cat" && animalType != "dog" && animalType != "bird" {
			info += `{"Error": "Unsuported Animal type !"}`
			fmt.Fprint(writer, info)
			return
		}

		animals := DB.GetAnimalsByType(animalType)
		info += formatAnimalinfo(animals)

		fmt.Fprint(writer, info)
		return
	} else {
		info += formatAnimalinfo(animals)

		fmt.Fprint(writer, info)
	}

}

func ProcessAnimalHandler(writer http.ResponseWriter, request *http.Request) {

	info := ""
	typeAnimal := request.URL.Query().Get("type")
	name := request.URL.Query().Get("name")
	attr := request.URL.Query().Get("attr")

	switch typeAnimal {
	case "dog":
		dog := dog.Dog{Name: name, Breed: attr}
		DB.AddAnimal(name, dog)
		info = animal.ProcessAnimal(dog)
	case "cat":
		cat := cat.Cat{Name: name, Color: attr}
		DB.AddAnimal(name, cat)
	case "bird":
		bird := bird.Bird{Name: name, WingSpan: attr}
		DB.AddAnimal(name, bird)
		info = animal.ProcessAnimal(bird)
	default:
		info = animal.ProcessAnimal(typeAnimal)

	}
	fmt.Fprint(writer, info)

}

func DescribeAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	info := ""
	typeAnimal := request.URL.Query().Get("type")
	name := request.URL.Query().Get("name")
	attr := request.URL.Query().Get("attr")

	switch typeAnimal {
	case "dog":
		dog := dog.Dog{Name: name, Breed: attr}
		DB.AddAnimal(name, dog)
		info = animal.DescribeAnimal(dog)
	case "cat":
		cat := cat.Cat{Name: name, Color: attr}
		DB.AddAnimal(name, cat)
		info = animal.DescribeAnimal(cat)
	case "bird":
		bird := bird.Bird{Name: name, WingSpan: attr}
		DB.AddAnimal(name, bird)
		info = animal.DescribeAnimal(bird)
	default:
		info = animal.DescribeAnimal(typeAnimal)

	}
	fmt.Fprint(writer, info)

}

/// others methods

func GetAnimalByIdHandler(writer http.ResponseWriter, request *http.Request) {
	info := ""
	id := request.URL.Query().Get("id")
	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}
	a, exist := DB.GetAnimal(id)
	if exist {

		info += animal.GetAnimalInfo(a)
	} else {
		info += `{"Error": "Animal with given id does not exist!"}`
	}

	fmt.Fprint(writer, info)

}

func DeleteAnimalByIdHandler(writer http.ResponseWriter, request *http.Request) {
	info := ""
	id := request.URL.Query().Get("id")
	if id == "" {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}

	_, exist := DB.GetAnimal(id)
	if exist {

		if DB.DeleteAnimal(id) {
			info += `{"Sucess": "Animal with given id"` + " " + id + "  deleted" + "}"
		}

	} else {
		info += `{"Error": "Animal with given id does not exist!"}`
	}

	fmt.Fprint(writer, info)

}
