package categories

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Delete(id uint) (*Model, error)
	GetAll() ([]Model, error)
	// GetById(id uint) (*Model, error)
}

type repository struct {
	db *gorm.DB
}

var _ Repository = repository{}

func NewReporsitory(db *gorm.DB) Repository {
	return repository{db: db}
}

func (repo repository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.Table("Category").First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repository) Create(model Model) (uint, error) {
	err := repo.db.Table("Category").Create(&model).Error
	if err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (repo repository) Delete(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.Table("Category").Delete(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repository) GetAll() ([]Model, error) {
	//newsList := make([]*Model, 0)
	var results []Model
	err := repo.db.Table("Category").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
