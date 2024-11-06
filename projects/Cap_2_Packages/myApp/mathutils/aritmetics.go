package mathutils

func Add(a int, b int) int {

	sum := a + b
	return sum
}

// func Add(a float64, b float64) float64 {

// 	return a + b
// }

func Substract(a int, b int) int {
	less := a - b
	return less
}

func Mult(a int, b int) int {
	mult := a * b
	return mult
}

func Div(a int, b int) (bool, float64) {

	if a == 0 {
		return false, 0
	}
	return true, float64(a / b)
}
