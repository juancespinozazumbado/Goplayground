package transactions

import (
	"sync"

	log "github.com/sirupsen/logrus"
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

			log.Debug("Recieved Transaction:", trans)

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

		log.Info("Processed transaction ", trans)
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
		if !resp.Succes {
			log.Error("Transaccion result ", resp)
		}
		// fmt.Printf("Transaccion result %+v \n", resp)
		log.Info("Transaccion result ", resp)

		wg.Done()
	}
}
