package dog

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return "Wooof!"
}

func (d Dog) Category() string {
	return "Mammal"
}
