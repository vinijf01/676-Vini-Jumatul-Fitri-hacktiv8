package repository

import "go-unit-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
