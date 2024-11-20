package main

import (
	"fmt"
	"time"
)

func main() {

	car := &Vehicule{}
	// car.setStructure()
	// car.setDoors()
	// car.setEngine()
	// car.setElectronic()

	go car.setStructure()
	go car.setDoors()
	go car.setEngine()
	go car.setElectronic()
	time.Sleep(2 * time.Second)
}

type Vehicule struct {
	structue    bool
	doors       bool
	engine      bool
	electronics bool
}

func (v *Vehicule) setStructure() {
	time.Sleep(1 * time.Second)
	v.structue = true
	fmt.Println("Structure set")
}
func (v *Vehicule) setDoors() {
	time.Sleep(1 * time.Second)
	v.doors = true
	fmt.Println("Doors set")
}
func (v *Vehicule) setEngine() {
	time.Sleep(1 * time.Second)
	v.engine = true
	fmt.Println("Engine set")
}
func (v *Vehicule) setElectronic() {
	time.Sleep(1 * time.Second)
	v.electronics = true
	fmt.Println("Electronic set")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		// awiat 1 seconds to print the value
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
