package user

import "gorm.io/gorm"

type Repository interface {
	Get(id uint) (*Model, error)
	Delete(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Login(data LoginDTO) (*Model, error)
	Migration() error
}

type repository struct {
	db *gorm.DB
}

var _ Repository = repository{}

func NewReporsitory(db *gorm.DB) Repository {
	return repository{db: db}
}

func (repo repository) Delete(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.Delete(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repository) Create(model Model) (uint, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (repo repository) Login(data LoginDTO) (*Model, error) {
	model := &Model{Email: data.Email}
	err := repo.db.First(&model, "email = ?", data.Email).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (repo repository) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}
