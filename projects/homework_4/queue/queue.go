package queue

import (
	"sortingPackages/models"
	"sync"
)

var (
	UnsortedPackages []models.Package
	SortedPackages   = make(map[string][]models.Package)
	Mutex            sync.Mutex
)
