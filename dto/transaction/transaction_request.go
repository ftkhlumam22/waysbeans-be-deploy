package transactiondto

type CreateTransactionRequest struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	UserID     int    `json:"user_id"`
	TotalPrice int    `json:"total_price"`
}

type UpdateTransactionRequest struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	UserID     int    `json:"user_id"`
	TotalPrice int    `json:"total_price"`
}
