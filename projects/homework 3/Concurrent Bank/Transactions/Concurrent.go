package transactions

import (
	"fmt"
	"sync"
)

type TransactionReceiver struct {
	InputChannel  chan *Transaction
	OutputChannel chan *Transaction
}

func NewTransactionReceiver(ch chan *Transaction, cr chan *Transaction) *TransactionReceiver {
	reciever := &TransactionReceiver{
		InputChannel:  ch,
		OutputChannel: cr,
	}
	go reciever.start()
	return reciever
}

func (t *TransactionReceiver) start() {

	/// method to be implemented

	for {
		select {
		case trans := <-t.InputChannel:

			fmt.Printf("Recieved Transaction: %+v \n", trans)

			t.OutputChannel <- trans
			//close(t.OutputChannel)
		}
	}

}

////////// Processor  ////////////////////

var Mutex sync.Mutex

type TransactionProcessor struct {
	InputChannel   chan *Transaction
	OutputChannel  chan *TransactionResponse
	AccountBalance int
}

func NewTransactionProcessor(ch chan *Transaction, cr chan *TransactionResponse) *TransactionProcessor {
	processor := &TransactionProcessor{
		InputChannel:  ch,
		OutputChannel: cr,
	}
	go processor.start()
	return processor
}

func (t *TransactionProcessor) start() {

	/// method to be implemented
	for trans := range t.InputChannel {
		Mutex.Lock()

		if trans.Type == "whithdraw" {
			if trans.Amount > t.AccountBalance {
				t.OutputChannel <- &TransactionResponse{TransactionId: trans.ID, Succes: false, NewBalance: 0, Message: "insuficient fonds !"}
				Mutex.Unlock()
				continue
			}

			t.AccountBalance -= trans.Amount
			t.OutputChannel <- &TransactionResponse{TransactionId: trans.ID, Succes: true, NewBalance: t.AccountBalance, Message: "withdraw succesfully!"}

		} else {
			t.AccountBalance += trans.Amount
			t.OutputChannel <- &TransactionResponse{TransactionId: trans.ID, Succes: true, NewBalance: t.AccountBalance, Message: "deposit succesfully!"}

		}
		Mutex.Unlock()

		fmt.Printf("Processed transaction %+v \n", trans)
	}

	//}

}

////////// Reposnse Handler  ////////////////////

type ResponseHandler struct {
	InputChannel chan *TransactionResponse
}

func NewResponseHandler(ch chan *TransactionResponse, wg *sync.WaitGroup) *ResponseHandler {
	handler := &ResponseHandler{
		InputChannel: ch,
	}

	go handler.start(wg)
	return handler
}

func (t *ResponseHandler) start(wg *sync.WaitGroup) {
	for resp := range t.InputChannel {
		fmt.Printf("Transaccion result %+v \n", resp)

		wg.Done()
	}
}
