package main

import (
	"fmt"
	"time"
)

func main() {

	result := make(chan int)
	go sum(2, 5, result)

	sumResult := <-result

	fmt.Printf("REsult recived from chanel %+v \n", sumResult)

	// go func() {
	// 	listener(ch)

	// }()

	// fmt.Println("Queue value")
	// ch <- 3
	// fmt.Println("Queue value")
	// ch <- 11.5
	// fmt.Println("Queue value")
	// ch <- "Hellow"
	// fmt.Println("Queue value")
	// ch <- true
	// fmt.Println("Queue value")
	// ch <- 'A'

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

	time.Sleep(5 * time.Second)
}

func sum(a, b int, result chan int) {

	time.Sleep(3 * time.Second)
	result <- a + b
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
