package repositories

import (
	"github.com/filipegms5/pokedex-backend/models"
	"gorm.io/gorm"
)

type TypeRepository struct {
	db *gorm.DB
}

func NewTypeRepository(db *gorm.DB) *TypeRepository {
	return &TypeRepository{db}
}

func (r *TypeRepository) FindAll() ([]models.Type, error) {
	var typeList []models.Type
	err := r.db.Find(&typeList).Error
	if err != nil {
		return nil, err
	}
	return typeList, err
}
func (r *TypeRepository) FindOneByName(name string) (models.Type, error) {
	var typePokemon models.Type
	err := r.db.Where("name = ?", name).Find(&typePokemon).Error
	if err != nil {
		return models.Type{}, err
	}
	return typePokemon, err
}
func (r *TypeRepository) FindOneById(id int) (models.Type, error) {
	var typePokemon models.Type
	err := r.db.First(&typePokemon, id).Error
	if err != nil {
		return models.Type{}, err
	}
	return typePokemon, err
}
