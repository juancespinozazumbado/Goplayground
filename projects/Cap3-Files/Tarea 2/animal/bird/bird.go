package bird

type Bird struct {
	Name     string
	WingSpan string
}

func (b Bird) Speak() string {
	return "Chirp!"
}

func (b Bird) Category() string {
	return "Aves"
}
