package models

type Pokemon struct {
	Id       int    `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Info     string `gorm:"type:varchar(255);NOT NULL" json:"info" binding:"required"`
	PhotoUrl string `gorm:"type:varchar(255);NOT NULL" json:"photoUrl" binding:"required"`
	Types    []Type `gorm:"many2many:pokemon_types;" json:"types" binding:"required"`
}
