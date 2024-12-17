package transactions

type Transaction struct {
	ID     string
	Type   string
	Amount int
}

type TransactionResponse struct {
	TransactionId string
	Succes        bool
	NewBalance    int
	Message       string
}
