package main

import (
	"fmt"
	"math/rand/v2"
	"students/models"
)

func main() {

	//students := []models.Student{}
	//students = append(students, {})

	student := models.Student{ID: 1212, Age: 23}

	for i := 0; i < 10; i++ {
		student.AddGrade((rand.Float64()*100 + float64(i)))

	}

	// student.AddGrade(50.5)
	// student.AddGrade(100)
	// student.AddGrade(80.5)

	models.DisplayStudentInfo(student)

	average, canCalculate := models.CalculateAverage(student)
	if canCalculate {
		fmt.Printf("The average of the notes are: %f\n", average)
	} else {
		fmt.Printf("Canot calculate average: %f", average)
	}

	max, min, canCalculate := models.GetMinMaxGrade(student)

	if canCalculate {

		fmt.Printf("Min note : %f Max note: %f\n", min, max)
	} else {

		fmt.Printf("Canot calculate average: %f", min)
	}

	fmt.Print("")
}
