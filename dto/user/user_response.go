package userdto

type UserResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Admin      bool   `json:"admin"`
	Address    string `json:"address"`
	PostalCode int    `json:"postal_code"`
}

type UserResponseID struct {
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	Admin       bool                `json:"admin"`
	Address     string              `json:"address"`
	PostalCode  int                 `json:"postal_code"`
	Transaction TransactionRelation `json:"transaction"`
}

type TransactionRelation struct {
	Status     string         `json:"status"`
	Cart       []CartRelation `json:"cart"`
	TotalPrice int            `json:"total_price"`
}

type CartRelation struct {
	ID      int             `json:"id"`
	Product ProductRelation `json:"product"`
	Qty     int             `json:"qty"`
	Amount  int             `json:"amount"`
}

type ProductRelation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}
