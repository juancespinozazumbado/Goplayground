package main

import (
	"fmt"

	"Basics/Arrays"
)

func main() {

	sales := [7]int{100, 200, 50, 400, 900, 600, 700}

	Arrays.DisplaySales(sales)
	fmt.Println("")
	Arrays.CalculateTotalSales(sales)

	higestVale, index := Arrays.LookHigestValue(sales)
	fmt.Printf("Max  Sale are: %d and is for the diay %d\n", higestVale, index+1)

	lower, index := Arrays.LookLowerValue(sales)
	fmt.Printf("Min Sale are: %d and is for the diay %d\n", lower, index+1)

	Arrays.Show()

}
