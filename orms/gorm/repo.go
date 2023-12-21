package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/crud"
	"github.com/GoLangWebSDK/crud/database"
	"gorm.io/gorm"
)

var _ crud.Repository[any] = (*Repository[any])(nil)

type Repository[T any] struct {
	DB             *gorm.DB
	Model          T
	deletedAtQuery string
}

func NewRepository[T any](db *database.Database, model T) *Repository[T] {
	repo := &Repository[T]{
		DB:             db.Adapter.Gorm(),
		Model:          model,
		deletedAtQuery: "%s.deleted_at IS NULL",
	}
	return repo
}

func (repo *Repository[T]) Create(model T) (T, error) {
	err := repo.DB.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (repo *Repository[T]) ReadAll() ([]T, error) {
	records := []T{}

	err := repo.DB.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (repo *Repository[T]) Read(ID uint) (T, error) {
	var record T
	err := repo.DB.Where("id = ?", ID).First(&record).Error
	if err != nil {
		return record, err
	}

	return record, nil
}

func (repo *Repository[T]) Update(ID uint, model T) (T, error) {
	if ID == 0 {
		return model, fmt.Errorf("missing ID")
	}

	err := repo.DB.First(&model, ID).Error
	if err != nil {
		return model, err
	}

	err = repo.DB.Model(&model).Where("id = ?", ID).Updates(model).Error
	if err != nil {
		return model, err
	}

	return model, nil
}

func (repo *Repository[T]) Delete(ID uint) error {
	if ID == 0 {
		return fmt.Errorf("missing ID")
	}

	err := repo.DB.First(&repo.Model, ID).Error
	if err != nil {
		return err
	}

	return repo.DB.Delete(&repo.Model, ID).Error
}
