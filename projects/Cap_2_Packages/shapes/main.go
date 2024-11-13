package main

import (
	"shapes/shapes"

	log "github.com/sirupsen/logrus"
)

func main() {

	shapesArray := []shapes.Shape{}

	circle := &shapes.Circle{Radius: 12.5}
	triangle := &shapes.Triangle{Base: 124.33, Higt: 1254.44}
	rectangle := &shapes.Rectangle{Width: 125.22, Long: 12154.022}

	shapesArray = append(shapesArray, circle)
	shapesArray = append(shapesArray, rectangle)
	shapesArray = append(shapesArray, triangle)

	showShapeArea(shapesArray)

}

func showShapeArea(shapes []shapes.Shape) {

	// log.WithFields(log.Fields{
	// 	"Section": "Testing interfaces",
	// }).Info("Starting the app...")

	log.Info(shapes)

	for _, shape := range shapes {
		if validateShape(shape) {
			log.Info(shape.Area())
		}

	}

	// fmt.Printf("Shape type: %d\n", reflect.TypeOf(shape).Kind())
	// fmt.Printf("Shape Area: %f\n", shape.Area())
}

func validateShape(input interface{}) bool {

	_, ok := input.(shapes.Shape)
	return ok
}
