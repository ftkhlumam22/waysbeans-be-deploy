package transactiondto

type TransactionResponse struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	UserID     int    `json:"user_id"`
	TotalPrice int    `json:"total_price"`
}
