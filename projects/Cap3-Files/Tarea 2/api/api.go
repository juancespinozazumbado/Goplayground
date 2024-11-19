package api

import (
	"animals/animal"
	"animals/animal/bird"
	"animals/animal/cat"
	"animals/animal/dog"
	"animals/data"
	"fmt"
	"html"
	"net/http"
)

var DB *data.MemoryDb

func HandleAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetAnimalsHandler(w http.ResponseWriter, r *http.Request) {

	// set animals data
	animals := DB.GetAnimals()
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

	fmt.Fprint(w, info)

}

func AddAnimal(writer http.ResponseWriter, request *http.Request) {

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
	case "bird":
		bird := bird.Bird{Name: name, WingSpan: attr}
		DB.AddAnimal(name, bird)
		info = animal.DescribeAnimal(bird)
	default:
		info = animal.DescribeAnimal(typeAnimal)

	}
	fmt.Fprint(writer, info)

}
