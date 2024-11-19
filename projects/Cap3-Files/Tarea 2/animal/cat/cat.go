package cat

type Cat struct {
	Name  string
	Color string
}

func (c Cat) Speak() string {
	return "Meow!"
}
func (c Cat) Category() string {
	return "Mammal"
}
