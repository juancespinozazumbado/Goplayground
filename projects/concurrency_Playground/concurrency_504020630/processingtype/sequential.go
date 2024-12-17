package processingtype

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SequentialProcessor struct {
	results map[string]int
}

func NewSequentialProcessor() *SequentialProcessor {
	return &SequentialProcessor{
		results: make(map[string]int),
	}
}

func (c *SequentialProcessor) Process() {
	var result int
	filesPath := "./files"

	startTime := time.Now()

	files, err := os.ReadDir(filesPath)
	if err != nil {
		fmt.Println("there was an error reading the directory: ", err)
	}

	//************************************************************************//
	//Step 2:
	// - For each loop, call the processFile method and print the result - use this format "File Name: %s, Result: %d\n"
	// - For each loop sum up the result into the existent variable var result int in line 23
	//************************************************************************//
	for _, file := range files {
		c.results[file.Name()] = 0

		//code here

		sum := c.processFile(file.Name())
		result += sum
	}

	elapsed := time.Since(startTime)
	fmt.Printf("The total result for all files is %d. Time taken: %f\n", result, elapsed.Seconds())
}

func (c *SequentialProcessor) processFile(fileName string) int {
	var result int
	filesPath := "./files"

	file, err := os.Open(filesPath + "/" + fileName)
	if err != nil {
		fmt.Println("there was an error opening the file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += 1
		//************************************************************************//
		//Step 3: call the c.processLine method and sum up the result into "var result variable" at line 49.
		//handle any possible errors
		//************************************************************************//
		//code here

		sum, _ := c.processLine(line)
		fmt.Printf("The total result for the multiply of the two values of file line %d is %d \n", count, *sum)
		result += *sum
	}

	fmt.Printf("The total result for the file %s is %d \n", fileName, result)
	return result
}

func (c *SequentialProcessor) processLine(lineData string) (*int, error) {
	splitData := strings.Split(lineData, ",")

	number1String := splitData[0]
	number2String := splitData[1]

	number1String = strings.TrimSpace(number1String)
	number2String = strings.TrimSpace(number2String)

	number1, err := strconv.Atoi(number1String)
	if err != nil {
		return nil, err
	}

	number2, err := strconv.Atoi(number2String)
	if err != nil {
		return nil, err
	}

	//************************************************************************//
	//Step 4: Multiply the 2 numbers and return the result
	//************************************************************************//
	//code here

	result := number1 * number2
	return &result, nil
}
