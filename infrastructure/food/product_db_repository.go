package food

import (
	"ndv/domain/food"
)

type ProductDBRepository interface {
	GetProductsByPartner(id int64) ([]*food.Product, error)
	SaveProduct(model *food.Product) (int64, error)
}
