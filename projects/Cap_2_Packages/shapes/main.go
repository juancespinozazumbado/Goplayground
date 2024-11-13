package main

import (
	"fmt"
	"shapes/shapes"
)

func main() {

	circle := shapes.Circle{Radius: 12.5}
	triangle := shapes.Triangle{Base: 124.33, Higt: 1254.44}
	rectangle := shapes.Rectangle{Width: 125.22, Long: 12154.022}

	showShapeArea(&circle)
	showShapeArea(&triangle)
	showShapeArea(&rectangle)

}

func showShapeArea(shape shapes.Shape) {
	//fmt.Printf("Shape type: %d\n", shape.(types.Type))
	fmt.Printf("Shape Area: %f\n", shape.Area())
}
