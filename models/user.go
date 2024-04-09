package models

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" gorm:"type:varchar(255)"  json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}