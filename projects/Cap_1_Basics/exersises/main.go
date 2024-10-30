package main

import (
	"Basics/Arrays"
	"Basics/structs"
	"fmt"
)

//
book := structs.Book{"Title", "Author", 20, false}

books := []structs.Book{ {"Title", "Author", 20, true}, }

func main() {
	
	AddBook(book, books)
}

func showArrayExersises() {

	sales := [7]int{100, 200, 50, 400, 900, 600, 700}

	fmt.Println("Show Sales flow")

	Arrays.DisplaySales(sales)
	fmt.Println("")
	Arrays.CalculateTotalSales(sales)

	higestVale, index := Arrays.LookHigestValue(sales)
	fmt.Printf("Max  Sale are: %d and is for the diay %d\n", higestVale, index+1)

	lower, index := Arrays.LookLowerValue(sales)
	fmt.Printf("Min Sale are: %d and is for the diay %d\n", lower, index+1)

	fmt.Println("Show Expenses flow")
	Arrays.Show()
}
