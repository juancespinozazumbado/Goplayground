package models

import "fmt"

type Student struct {
	ID     int
	Age    int
	Grades []float64
}

func (student *Student) AddGrade(grade float64) {

	student.Grades = append(student.Grades, grade)
}

func CalculateAverage(student Student) (float64, bool) {

	average := 0.0
	grades := student.Grades
	length := len(grades)
	sum := 0.0

	if length == 0 {
		return 0.0, false
	}
	for _, grade := range grades {

		sum += grade
	}

	average = sum / float64(length)

	return average, true
}

func GetMinMaxGrade(student Student) (float64, float64, bool) {

	grades := student.Grades
	max := grades[0]
	min := grades[0]
	length := len(grades)

	if length == 0 {
		return 0.0, 0.0, false
	}

	for _, grade := range grades {

		if max < grade {
			max = grade
		}
		if min > grade {
			min = grade
		}
	}

	return max, min, true
}

func DisplayStudentInfo(student Student) {

	fmt.Printf("Student data: %+v\n", student)
}
