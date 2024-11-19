package data

import (
	"animals/animal"

	log "github.com/sirupsen/logrus"
)

type MemoryDb struct {
	data map[string]animal.Animal
}

func NewDB() *MemoryDb {

	return &MemoryDb{data: make(map[string]animal.Animal)}
}

func (db *MemoryDb) AddAnimal(id string, animal animal.Animal) {

	db.data[id] = animal
}

func (db *MemoryDb) GetAnimals() []animal.Animal {

	animals := make([]animal.Animal, 0, len(db.data))

	for _, animal := range db.data {
		animals = append(animals, animal)
	}
	return animals
}

func (db *MemoryDb) GetAnimal(id string) (animal.Animal, bool) {

	animal, exist := db.data[id]
	log.Info("animal", animal)
	if exist {
		return animal, true

	}
	return nil, false
}

func (db *MemoryDb) UpdateAnimal(id string, newAnimalData animal.Animal) bool {

	_, exist := db.data[id]
	if exist {
		db.data[id] = newAnimalData
		return true

	}
	return false
}

func (db *MemoryDb) DeleteAnimal(id string) bool {

	_, exist := db.data[id]
	if exist {
		delete(db.data, id)
		return true

	}
	return false
}
