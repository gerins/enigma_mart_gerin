package product

import "database/sql"

type ProductService struct {
	productRepo ProductRepository
}

type ProductServiceInterface interface {
	GetProducts() (*[]Product, error)
	GetProductByID(id string) (*Product, error)
	HandlePOSTProduct(d Product) (*Product, error)
	HandleUPDATEProduct(id string, data Product) (*Product, error)
	HandleDELETEProduct(id string) (*Product, error)
}

func NewProductService(db *sql.DB) ProductServiceInterface {
	return ProductService{NewProductRepo(db)}
}

func (s ProductService) GetProducts() (*[]Product, error) {
	Product, err := s.productRepo.HandleGETAllProduct()
	if err != nil {
		return nil, err
	}

	return Product, nil
}

func (s ProductService) GetProductByID(id string) (*Product, error) {
	Product, err := s.productRepo.HandleGETProduct(id, "A")
	if err != nil {
		return nil, err
	}
	return Product, nil
}

func (s ProductService) HandlePOSTProduct(d Product) (*Product, error) {
	result, err := s.productRepo.HandlePOSTProduct(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s ProductService) HandleUPDATEProduct(id string, data Product) (*Product, error) {
	result, err := s.productRepo.HandleUPDATEProduct(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s ProductService) HandleDELETEProduct(id string) (*Product, error) {
	result, err := s.productRepo.HandleDELETEProduct(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
