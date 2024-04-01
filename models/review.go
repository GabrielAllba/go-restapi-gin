package models

type Review struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	ProductId int    `gorm:"index" json:"product_id"`
	Content   string `gorm:"type:text" json:"content"`
}
