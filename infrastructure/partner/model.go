package partner

import (
	"ndv/domain/partner"
	"time"
)

type PartnerDTO struct {
	ID   int64
	Name string
	//Address pointer to Address type. TBD
	AddressID   int64     `json:"address_id"`
	Image       string    `json:"image"`
	Phone       string    `json:"phone"`
	Mail        string    `json:"mail"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerType int64     `json:"partner_type"`
	IsDeleted   bool      `json:"is_deleted"`
}

func FromDtoToPartner(dto *PartnerDTO) *partner.Partner {
	return &partner.Partner{
		ID:          dto.ID,
		Name:        dto.Name,
		AddressID:   dto.AddressID,
		Image:       dto.Image,
		Phone:       dto.Phone,
		Mail:        dto.Mail,
		Rating:      dto.Rating,
		CreatedAt:   dto.CreatedAt,
		LastUpdated: dto.LastUpdated,
		PartnerType: dto.PartnerType,
		IsDeleted:   dto.IsDeleted,
	}
}

func FromPartnerToDTO(partner *partner.Partner) *PartnerDTO {
	return &PartnerDTO{
		ID:          partner.ID,
		Name:        partner.Name,
		AddressID:   partner.AddressID,
		Image:       partner.Image,
		Phone:       partner.Phone,
		Mail:        partner.Mail,
		Rating:      partner.Rating,
		CreatedAt:   partner.CreatedAt,
		LastUpdated: partner.LastUpdated,
		PartnerType: partner.PartnerType,
		IsDeleted:   partner.IsDeleted,
	}
}
