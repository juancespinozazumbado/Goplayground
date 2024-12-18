package transactions

import log "github.com/sirupsen/logrus"

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
