package main

import (
	transactions "ConcurrentBank/Transactions"
	"sync"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {

	trascInvCHannel := make(chan *transactions.Transaction)
	trascOutCHannel := make(chan *transactions.Transaction)
	ResponseChan := make(chan *transactions.TransactionResponse)

	handler := transactions.NewResponseHandler(ResponseChan, &wg)
	processor := transactions.NewTransactionProcessor(trascOutCHannel, ResponseChan)
	reciever := transactions.NewTransactionReceiver(trascInvCHannel, trascOutCHannel)

	//log.SetLevel(log.TraceLevel)
	log.Info("Setting up components....")

	log.Info("Reciever:", reciever)
	log.Info("Processor:", processor)
	log.Info("Hanlder", handler)

	log.Info("Process started....")

	// sequence of transactions
	// add 4 to waitGroup case it is 4 transactions to proecess
	wg.Add(8)

	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "deposit", Amount: 200} // create a new instance of Transaction (pointer)
	trascInvCHannel <- &transactions.Transaction{ID: "454545", Type: "deposit", Amount: 300}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 1000}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 400}

	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "deposit", Amount: 500} // create a new instance of Transaction (pointer)
	trascInvCHannel <- &transactions.Transaction{ID: "454545", Type: "deposit", Amount: 500}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 30000}
	trascInvCHannel <- &transactions.Transaction{ID: "1234", Type: "whithdraw", Amount: 10}

	wg.Wait() // wait til each process finish

}
