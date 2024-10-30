package main

import (
	"fmt"
	"students/models"
)

func main() {

	//students := []models.Student{}
	//students = append(students, {})

	student := models.Student{ID: 1212, Age: 23}

	student.AddGrade(50.5)
	student.AddGrade(100)
	student.AddGrade(80.5)

	models.DisplayStudentInfo(student)

	average, canCalculate := models.CalculateAverage(student)
	if canCalculate {
		fmt.Printf("The average of the notes are: %f\n", average)
	} else {
		fmt.Printf("Canot calculate average: %f", average)
	}

	min, max, canCalculate := models.GetMinMaxGrade(student)

	if canCalculate {

		fmt.Printf("Min note : %f Max note: %f\n", min, max)
	} else {

		fmt.Printf("Canot calculate average: %f", min)
	}

	fmt.Print("")
}
