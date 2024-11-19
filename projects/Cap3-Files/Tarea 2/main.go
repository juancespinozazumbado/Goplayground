package main

import (
	"animals/animal"
	"animals/animal/bird"
	"animals/animal/cat"
	"animals/animal/dog"
	"fmt"
)

func main() {

	animals := []animal.Animal{
		cat.Cat{Name: "Tom", Color: "Gray"},
		bird.Bird{Name: "Kity", WingSpan: "FFFFD"},
	}

	tom := dog.Dog{Name: "Tom", Breed: "Chiken"}
	animals = append(animals, tom)

	info := "[\n"

	for _, current := range animals {
		info += " " + animal.GetAnimalInfo(current) + ",\n"
	}
	info += "]"
	fmt.Println(info)

	for _, a := range animals {

		fmt.Println(animal.DescribeAnimal(a))

	}

	for _, a := range animals {

		fmt.Println(animal.ProcessAnimal(a))

	}

	fmt.Println(animal.ProcessAnimal("Zorro"))
	fmt.Println(animal.DescribeAnimal("Perico"))

}
