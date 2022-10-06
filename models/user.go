package models

import "time"

type User struct {
	ID            int         `json:"id"`
	Name          string      `json:"name" gorm:"type: varchar(255)"`
	Email         string      `json:"email" gorm:"type: varchar(255)"`
	Password      string      `json:"password" gorm:"type: varchar(255)"`
	Admin         bool        `json:"admin" gorm:"type: bool"`
	Address       string      `json:"address" gorm:"type: varchar(255)"`
	PostalCode    int         `json:"postal_code" gorm:"type: int"`
	TransactionID int         `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type UserRelation struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}

func (UserRelation) TableName() string {
	return "users"
}
