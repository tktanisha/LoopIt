package category_service

import (
	"loopit/internal/models"
	"loopit/internal/repository/category_repo"
)

type CategoryService struct {
	categoryRepo category_repo.CategoryRepo
}

func NewCategoryService(repo category_repo.CategoryRepo) *CategoryService {
	return &CategoryService{categoryRepo: repo}
}

func (c *CategoryService) GetAllCategories() ([]models.Category, error) {
	return c.categoryRepo.FindAll()
}
