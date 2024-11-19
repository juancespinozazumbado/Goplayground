package main

import (
	"net/http"
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

	validateShapeSwitch(circle)
	validateShapeSwitch(rectangle)
	validateShapeSwitch(triangle)
	validateShapeSwitch("Any")

	http.Handle("/shapes", nil)

	http.ListenAndServe(":8080", nil)

}

func showShapeArea(shapes []shapes.Shape) {

	log.Info(shapes)

	for _, shape := range shapes {
		if validateShape(shape) {
			log.Info(shape.Area())
		} else {
			log.Error("input doesn't implement the shape interface", shape)
		}

	}

	// fmt.Printf("Shape type: %d\n", reflect.TypeOf(shape).Kind())
	// fmt.Printf("Shape Area: %f\n", shape.Area())
}

func validateShape(input interface{}) bool {

	_, ok := input.(shapes.Shape)
	return ok
}

func validateShapeSwitch(input interface{}) {

	switch i := input.(type) {
	case *shapes.Rectangle:
		log.WithFields(log.Fields{
			"Type":  "Rectangule",
			"Long":  i.Long,
			"Width": i.Width,
			"Area":  i.Area(),
		}).Info("info of the shape ", input)

	case *shapes.Circle:
		log.WithFields(log.Fields{
			"Type":   "Circle",
			"Radius": i.Radius,
			"Area":   i.Area(),
		}).Info("info of the shape ", input)

	case *shapes.Triangle:
		log.WithFields(log.Fields{
			"Type":  "Triangle",
			"Base":  i.Base,
			"Higth": i.Higt,
			"Area":  i.Area(),
		}).Info("info of the shape ", input)

	default:
		log.Error("input doesnt imnplement interface shape ", i)

	}
}
