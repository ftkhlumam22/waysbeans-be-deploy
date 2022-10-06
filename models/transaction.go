package models

type Transaction struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	Status     string       `json:"status" gorm:"type: varchar(255)"`
	UserID     int          `json:"user_id"`
	User       UserRelation `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart       []Cart       `json:"cart" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice int          `json:"total_price" gorm:"type: int"`
}
