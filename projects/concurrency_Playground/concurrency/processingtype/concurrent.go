package processingtype

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	filesPath = "./files"
)

var mu sync.Mutex

type fileResult struct {
	fileName string
	result   int
}
type ConcurrentProcessor struct {
	results map[string]int
}

func NewConcurrentProcessor() *ConcurrentProcessor {
	return &ConcurrentProcessor{
		results: make(map[string]int),
	}
}

func (c *ConcurrentProcessor) Process() {
	var result int

	startTime := time.Now()

	files, err := os.ReadDir(filesPath)
	if err != nil {
		fmt.Println("there was an error reading the directory: ", err)
	}

	lineByLineResultChannel := make(chan fileResult)
	closeLineByLineResultChannel := make(chan bool)

	//************************************************************************//
	//Step 2: Start the listener: c.listenToEachLine. Remember to do this concurrently
	//************************************************************************//
	// code here

	go c.listenToEachLine(lineByLineResultChannel, closeLineByLineResultChannel)

	var filesWG sync.WaitGroup
	for _, file := range files {
		filesWG.Add(1)
		c.results[file.Name()] = 0

		//************************************************************************//
		//Step 3: call the c.processFile method concurrently
		//************************************************************************//
		// code here
		go c.processFile(file.Name(), lineByLineResultChannel, &filesWG)
	}

	//************************************************************************//
	//Step 4: wait for filesWG to be completely done
	//************************************************************************//
	// code here

	closeLineByLineResultChannel <- true
	close(lineByLineResultChannel)

	mu.Lock()
	for fileName, value := range c.results {
		result += value
		fmt.Printf("File Name: %s, Result: %d\n", fileName, value)
	}
	mu.Unlock()

	elapsed := time.Since(startTime)
	fmt.Printf("The total result for all files is %d. Time taken: %f\n", result, elapsed.Seconds())
}

// This function is the receiver of each line multiplication.
// Step 5: Important: Remember to run tests to see possible racing conditions here and fix appropriately. (use "mu")
// hint: go run -race main.
func (c *ConcurrentProcessor) listenToEachLine(lineByLineResultChannel <-chan fileResult, closeLineByLineResultChannel <-chan bool) {
	for {
		select {
		case result := <-lineByLineResultChannel:
			//************************************************************************//
			//code here
			//************************************************************************//
			c.results[result.fileName] += result.result
			//************************************************************************//
			//code here
			//************************************************************************//
		case <-closeLineByLineResultChannel:
			break
		}
	}
}

func (c *ConcurrentProcessor) processFile(fileName string, lineByLineResultChannel chan fileResult, filesWG *sync.WaitGroup) {
	file, err := os.Open(filesPath + "/" + fileName)
	if err != nil {
		fmt.Println("there was an error opening the file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//************************************************************************//
	//Step 6: create a waitGroup called linesWG to:
	// - Add(1) within the for loop
	// - Pass linesWG along to c.processLine
	//************************************************************************//
	var linesWG sync.WaitGroup
	for scanner.Scan() {
		//************************************************************************//
		//code here
		//************************************************************************//
		line := scanner.Text()

		go c.processLine(fileName, line, lineByLineResultChannel, &linesWG)
	}

	linesWG.Wait()
	filesWG.Done()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (c *ConcurrentProcessor) processLine(fileName, lineData string, lineByLineChannel chan<- fileResult, linesWG *sync.WaitGroup) {
	defer linesWG.Done()

	splitData := strings.Split(lineData, ",")

	number1String := splitData[0]
	number2String := splitData[1]

	number1String = strings.TrimSpace(number1String)
	number2String = strings.TrimSpace(number2String)

	number1, err := strconv.Atoi(number1String)
	if err != nil {
		log.Fatal(err)
	}

	number2, err := strconv.Atoi(number2String)
	if err != nil {
		log.Fatal(err)
	}

	var result int
	result = number1 * number2

	//************************************************************************//
	//Step 7: [Sender]: Proceed sending the result to the lineByLineChannel channel
	//Take into account the channel's type is fileResult
	// type fileResult struct {
	// 	fileName string -> use the fileName parameter
	// 	result   int -> use the result
	// }
	//************************************************************************//
	// code here

	lineByLineChannel <- fileResult{fileName: fileName, result: result}

}
