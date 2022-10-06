package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Price       int       `json:"price" gorm:"type: int"`
	Image       string    `json:"image" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: text"`
	Stock       int       `json:"stock" gorm:"type: int"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ProductRelation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

func (ProductRelation) TableName() string {
	return "products"
}
