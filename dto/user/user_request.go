package userdto

type CreateUserRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Admin      bool   `json:"admin" form:"admin"`
	Address    string `json:"address" form:"address"`
	PostalCode int    `json:"postal_code" form:"postal_code"`
}

type UpdateUserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Admin      bool   `json:"admin"`
	Address    string `json:"address"`
	PostalCode int    `json:"postal_code"`
}
