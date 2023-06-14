package bundle

import (
	"ndv/domain/bundle"
	"time"
)

type BundleDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerID   int64     `json:"partner_id"`
	IsDeleted   bool      `json:"is_deleted"`
	Products    []int64   `json:"products"`
}

func FromDtoToBundle(dto *BundleDTO) *bundle.Bundle {
	return &bundle.Bundle{
		ID:          dto.ID,
		Name:        dto.Name,
		CreatedAt:   dto.CreatedAt,
		LastUpdated: dto.LastUpdated,
		IsDeleted:   dto.IsDeleted,
		Rating:      dto.Rating,
		PartnerID:   int(dto.PartnerID),
	}
}
