/*
Problem Statement
You are tasked with processing a list of text files. Each file contains several lines, with two numbers separated by a comma on each line. Your goal is to:

- Multiply the two numbers on each line.
- Calculate the sum of all the lines for each file.
- Compute the grand total of these sums across all files.

Specifications
 - Sequential Version:
	- Process files one at a time, reading and multiplying the numbers line by line sequentially.
	- Measure the time taken for processing.

 - Concurrent Version:
	- Call the processFile in a goroutine and create a goroutine for each line in the file.
	- Use the processLine method as a sender, to send the results into a one-directional channel (lineByLineResultChannel).
	- Use the listenToEachLine method as a receiver, to collect results from the channel and calculate the total for the file. This data will be stored in a map (c.results within the concurrentProcessor).


	IMPORTANT: There are different places in the code where you need to focus. Please find these spots with the comments "//Step X" and follow appropriately
	Do only what it's required in the steps.
*/

// ************************************************************************//
// Remember to run go mod init and go mod tidy
// ************************************************************************//
package main

import (
	"filesWork/processingtype"
)

func main() {
	useConcurrency := false

	if useConcurrency {
		println("Using concurrency")

		//************************************************************************//
		//Step 1: Create a new ConcurrentProcessor and hit Process()
		//************************************************************************//
		//code here

	} else {
		println("Using Sequential")

		//************************************************************************//
		//Step 2: Create a new SequentialProcessor and hit Process()
		//************************************************************************//
		//code here

		seqProcessor := processingtype.NewSequentialProcessor()
		seqProcessor.Process()

	}
}
