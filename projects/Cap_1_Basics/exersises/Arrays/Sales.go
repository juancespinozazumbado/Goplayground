package Arrays

import "fmt"

func DisplaySales(sales [7]int) {

	for i := 0; i < len(sales); i++ {
		fmt.Printf("Day %d sales: %d\n", i+1, sales[i])
	}
}

func CalculateTotalSales(sales [7]int) {

	var total = 0

	for i := 0; i < len(sales); i++ {
		total += sales[i]

	}

	fmt.Printf("Total sales: %d\n", total)

}

func LookHigestValue(sales [7]int) (int, int) {

	higest := sales[0]
	index := 0

	for i := 0; i < len(sales); i++ {

		if sales[i] > higest { // if the current value is greater than the higest
			higest = sales[i] // set the higest to current value
			index = i
		}

	}

	return higest, index
}

func LookLowerValue(sales [7]int) (int, int) {

	lower := sales[0]
	index := 0

	for i := 0; i < len(sales); i++ {

		if sales[i] < lower { // if the current value is lower than the higest
			lower = sales[i] // set the higest to current value
			index = i
		}
	}

	return lower, index
}
