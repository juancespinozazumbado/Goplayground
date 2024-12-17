package transactions

import "sync"

type TransactionReceiver struct {
	InputChannel  chan Transaction
	OutputChannel chan Transaction
}

type TransactionProcessor struct {
	InputChannel   chan Transaction
	OutputChannel  chan Transaction
	AccountBalance int
	Mutex          sync.Mutex
}

func (t *TransactionProcessor) Start() {

	/// method to be implemented

}

type ResponseHandler struct {
	InputChannel chan TransactionResponse
}

func (t *ResponseHandler) Start() {

	//method to be implemented

}
