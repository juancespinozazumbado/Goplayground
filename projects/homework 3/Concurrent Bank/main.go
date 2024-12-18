package main

import (
	transactions "ConcurrentBank/Transactions"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	trascInvCHannel := make(chan *transactions.Transaction)
	trascOutCHannel := make(chan *transactions.Transaction)
	ResponseChan := make(chan *transactions.TransactionResponse)

	handler := transactions.NewResponseHandler(ResponseChan, &wg)
	processor := transactions.NewTransactionProcessor(trascOutCHannel, ResponseChan)
	reciever := transactions.NewTransactionReceiver(trascInvCHannel, trascOutCHannel)

	fmt.Printf("Reciever: %+v\n", reciever)
	fmt.Printf("Processor: %+v\n", processor)
	fmt.Printf("Hanlder: %+v\n", handler)

	fmt.Println("Process started....")

	// sequence of transactions
	wg.Add(4)

	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "deposit", Amount: 200}
	trascInvCHannel <- &transactions.Transaction{ID: "454545", Type: "deposit", Amount: 300}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 1000}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 400}

	wg.Wait() // wait

}
