package models

type Package struct {
	PackageId   int    `json:"Id"`
	SenderName  string `json:"SenderName"`
	Destination string `json:"Destination"`
	Status      bool   `json:"status"`
}

func NewPackage(id int, sourceName string, Destination string) *Package {
	return &Package{PackageId: id,
		SenderName:  sourceName,
		Destination: Destination,
		Status:      false}
}
