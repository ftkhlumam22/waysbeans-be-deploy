package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	Admin bool   `gorm:"type: bool" json:"admin"`
}

type RegisterResponse struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Admin    bool   `gorm:"type: bool" json:"admin"`
}

type CheckAuthResponse struct {
	ID    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Admin bool   `gorm:"type: varchar(255)" json:"admin"`
}
