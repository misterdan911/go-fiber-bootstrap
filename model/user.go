package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" validate:"required" example:"danu"`
	Email    string `json:"email" validate:"required,email,email_unique" example:"dciptadi@gmail.com"`
	Password string `json:"password" validate:"required" example:"12345678"`
}
