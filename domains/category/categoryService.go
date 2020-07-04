package category

import (
	"database/sql"
)

type CategoryService struct {
	CategoryRepo CategoryRepository
}

type CategoryServiceInterface interface {
	GetCategorys() (*[]Category, error)
	GetCategoryByID(id string) (*Category, error)
	HandlePOSTCategory(d Category) (*Category, error)
	HandleUPDATECategory(id string, data Category) (*Category, error)
	HandleDELETECategory(id string) (*Category, error)
}

func NewCategoryService(db *sql.DB) CategoryServiceInterface {
	return CategoryService{NewCategoryRepo(db)}
}

func (s CategoryService) GetCategorys() (*[]Category, error) {
	Category, err := s.CategoryRepo.HandleGETAllCategory()
	if err != nil {
		return nil, err
	}

	return Category, nil
}

func (s CategoryService) GetCategoryByID(id string) (*Category, error) {
	Category, err := s.CategoryRepo.HandleGETCategory(id, "A")
	if err != nil {
		return nil, err
	}
	return Category, nil
}

func (s CategoryService) HandlePOSTCategory(d Category) (*Category, error) {
	result, err := s.CategoryRepo.HandlePOSTCategory(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s CategoryService) HandleUPDATECategory(id string, data Category) (*Category, error) {
	result, err := s.CategoryRepo.HandleUPDATECategory(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s CategoryService) HandleDELETECategory(id string) (*Category, error) {
	result, err := s.CategoryRepo.HandleDELETECategory(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
