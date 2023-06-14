package bundle

import (
	"ndv/domain/food"
	"time"
)

type Bundle struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerID   int       `json:"partner_id"`
	//Partner     partner.Partner `json:"partner"`
	IsDeleted bool `json:"is_deleted"`
	//TODO popular productos en bundles
	Products []food.Product `json:"products"`
}
