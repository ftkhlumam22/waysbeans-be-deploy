package authdto

type RegisterRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Admin      bool   `json:"admin" form:"admin"`
	Address    string `json:"address"`
	PostalCode int    `json:"postal_code"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
