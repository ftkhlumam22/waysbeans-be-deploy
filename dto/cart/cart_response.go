package cartdto

type CartResponse struct {
	ProductID     int         `json:"product_id"`
	Product       ProductCart `json:"product"`
	TransactionID int         `json:"transaction_id"`
	Qty           int         `json:"qty"`
	Amount        int         `json:"amount"`
}

type ProductCart struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}
