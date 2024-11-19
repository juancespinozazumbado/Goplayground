package animal

import (
	"animals/animal/bird"
	"animals/animal/cat"
	"animals/animal/dog"
	"fmt"
)

type Animal interface {
	Speak() string
	Category() string
}

func GetAnimalInfo(animal Animal) string {
	info := ""
	switch value := animal.(type) {
	case dog.Dog:
		info += "{\n" + fmt.Sprintf(`  "name": "%s",`, value.Name) + "\n" + fmt.Sprintf(`  "Breed": "%s",`, value.Breed) + "\n"
		info += `  "Methods": {` + "\n" + fmt.Sprintf(`   "Speak": "%s", `, value.Speak()) + "\n" + fmt.Sprintf(`   "Category": "%s", `, value.Category())
		info += "\n   }" + "\n" + " }"
	case cat.Cat:
		info += "{\n" + fmt.Sprintf(`  "name": "%s",`, value.Name) + "\n" + fmt.Sprintf(`  "Color": "%s",`, value.Color) + "\n"
		info += `  "Methods": {` + "\n" + fmt.Sprintf(`   "Speak": "%s", `, value.Speak()) + "\n" + fmt.Sprintf(`   "Category": "%s", `, value.Category())
		info += "\n   }" + "\n" + " }"
	case bird.Bird:
		info += "{\n" + fmt.Sprintf(`  "name": "%s",`, value.Name) + "\n" + fmt.Sprintf(`  "WingSpan": "%s",`, value.WingSpan) + "\n"
		info += `  "Methods": {` + "\n" + fmt.Sprintf(`   "Speak": "%s", `, value.Speak()) + "\n" + fmt.Sprintf(`   "Category": "%s", `, value.Category())
		info += "\n   }" + "\n" + " }"
	default:
		info += `{"Error":"Animal type not found"}`

	}

	return info
}

func ProcessAnimal(input interface{}) string {

	info := ""
	_, ok := input.(Animal)
	if ok {
		info += "{\n" + `  "isAnimalType" : true` + "\n}"

	} else {
		info += "{\n" + `  "isAnimalType" : false` + "\n}"

	}
	return info
}

func DescribeAnimal(animal interface{}) string {
	info := ""
	//v, ok := animal.(Animal)
	// if ok {
	switch t := animal.(type) {
	case dog.Dog:
		info += "{\n" + `  "Type" : "Dog"` + "\n}"
		t.Speak()
	case bird.Bird:
		info += "{\n" + `  "Type" : "Bird"` + "\n}"
		t.Speak()
	case cat.Cat:
		info += "{\n" + `  "Type" : "Cat"` + "\n}"
		t.Speak()
	default:
		info += "{\n" + `  "Error" : "input doesn't implement Animal interface."` + "\n}"

	}

	//}

	return info
}
