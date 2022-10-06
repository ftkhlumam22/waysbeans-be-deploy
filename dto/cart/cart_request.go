package cartdto

type CreateCartRequest struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty"`
	Amount        int `json:"amount"`
}

type UpdateCartRequest struct {
	ProductID     int `json:"product_id"`
	TransactionID int `json:"transaction_id"`
	Qty           int `json:"qty"`
	Amount        int `json:"amount"`
}
