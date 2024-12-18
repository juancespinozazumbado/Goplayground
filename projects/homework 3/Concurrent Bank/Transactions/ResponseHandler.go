package transactions

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

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
