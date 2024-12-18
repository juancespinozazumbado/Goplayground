package transactions

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

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
