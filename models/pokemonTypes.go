package models

type PokemonTypes struct {
	PokemonId int `gorm:"primaryKey"`
	TypeId    int `gorm:"primaryKey"`
}
