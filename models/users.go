package models

type User struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"id"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}
