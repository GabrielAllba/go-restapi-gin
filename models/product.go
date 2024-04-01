package models

type Product struct {
	Id          int      `gorm:"primaryKey" json:"id"`
	NamaProduct string   `gorm:"type:varchar(255)" json:"nama_product"`
	Deskripsi   string   `gorm:"type:varchar(255)" json:"deskripsi"`
	Image       string   `gorm:"type:text" json:"image"`
	Reviews     []Review `gorm:"foreignKey:ProductId" json:"reviews"`
}