package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan interface{})
	go func() {
		listener(ch)

	}()

	ch <- 3
	ch <- 11.5
	ch <- "Hellow"
	ch <- true

	//car := &Vehicule{}
	// car.setStructure()
	// car.setDoors()
	// car.setEngine()
	// car.setElectronic()

	// go car.setDoors()
	// go car.setEngine()
	// go car.setElectronic()

	// go also suport anonymous funcs
	// go func() {
	// 	fmt.Println("Somethin is hapening...")
	// 	car.setStructure()
	// 	car.setDoors()
	// 	car.setEngine()
	// 	car.setElectronic()
	// 	car.setTires()
	// }()

	time.Sleep(11 * time.Second)
}

type Vehicule struct {
	structue    bool
	doors       bool
	engine      bool
	electronics bool
	tires       bool
}

func (v *Vehicule) setStructure() {
	time.Sleep(1 * time.Second)
	v.structue = true
	fmt.Println("Structure set")
}
func (v *Vehicule) setDoors() {
	time.Sleep(4 * time.Second)
	v.doors = true
	fmt.Println("Doors set")
}
func (v *Vehicule) setEngine() {
	time.Sleep(1 * time.Second)
	v.engine = true
	fmt.Println("Engine set")
}
func (v *Vehicule) setElectronic() {
	time.Sleep(2 * time.Second)
	v.electronics = true
	fmt.Println("Electronic set")
}

func (v *Vehicule) setTires() {
	time.Sleep(4 * time.Second)
	v.tires = true
	fmt.Println("Tires set")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		// awiat 1 seconds to print the value
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

///chanels

func listener(ch chan interface{}) {

	for chanVal := range ch {

		fmt.Printf("Value in the channel is %+v \n", chanVal)
	}
}
