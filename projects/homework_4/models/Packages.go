package models

type Package struct {
	PackageId   int
	SenderName  string
	Destination string
	Status      bool
}

func NewPackage(id int, sourceName string, Destination string) *Package {
	return &Package{PackageId: id,
		SenderName:  sourceName,
		Destination: Destination,
		Status:      false}

}
