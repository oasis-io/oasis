package model

type User struct {
	Model
	Username string `json:"username" gorm:"index;not null;"`
	Password string `json:"password"  gorm:"not null;"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
