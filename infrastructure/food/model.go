package food

import (
	"ndv/domain/food"
	"time"
)

type ProductDTO struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MaxQuantity int64   `json:"max_quantity"`
	AgeCheck    bool    `json:"age_check"`
	Price       float64 `json:"price"`
	//Category    category.Category `json:"category"`
	Image       string    `json:"image"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerID   int64     `json:"partner_id"`
}

func FromProductTODTO(product *food.Product) *ProductDTO {

	return &ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		MaxQuantity: product.MaxQuantity,
		AgeCheck:    product.AgeCheck,
		Price:       product.Price,
		Image:       product.Image,
		Rating:      product.Rating,
		PartnerID:   product.PartnerID,
	}
}
