package service

import (
	"errors"
	"fmt"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	fmt.Println(id)
	fmt.Println("=====================")
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}