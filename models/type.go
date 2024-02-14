package models

type Type struct {
	Id    int    `gorm:"primarykey" json:"id"`
	Name  string `json:"name" binding:"required" gorm:"type:varchar(255);NOT NULL"`
	Color string `json:"color" binding:"required" gorm:"type:varchar(255);NOT NULL"`
}
