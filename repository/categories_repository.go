package repository

import (
	"finalproject4/model"

	"gorm.io/gorm"
)

type CategoriesRepository interface {
	CreateCategory(categories model.Category) (model.Category, error)
	GetAllCategory() ([]model.Category, error)
	GetCategoryByID(id int) (model.Category, error)
	UpdateCategory(id int, categories model.Category) (model.Category, error)
	DeleteCategory(id int) error
}

type categoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *categoriesRepository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) CreateCategory(category model.Category) (model.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoriesRepository) GetAllCategory() ([]model.Category, error) {
	var category []model.Category
	err := r.db.Preload("Products").Find(&category).Error
	return category, err
}

func (r *categoriesRepository) GetCategoryByID(id int) (model.Category, error) {
	var category model.Category
	err := r.db.Find(&category, id).Error
	return category, err
}

func (r *categoriesRepository) UpdateCategory(id int, category model.Category) (model.Category, error) {
	err := r.db.Where("id = ?", id).Updates(&category).Error
	return category, err
}

func (r *categoriesRepository) DeleteCategory(id int) error {
	err := r.db.Where("id = ?", id).Delete(&model.Category{}).Error
	return err
}
