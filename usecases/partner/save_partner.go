package partner

import (
	"errors"
	"ndv/domain/partner"
	partnerInfra "ndv/infrastructure/partner"
)

type SavePartner interface {
	Execute(partner *SavePartnerModel) (int64, error)
}

type savePartner struct {
	partnerRepo partnerInfra.PartnerDBRepository
}

func NewSavePartner(partnerRepo partnerInfra.PartnerDBRepository) *savePartner {
	return &savePartner{
		partnerRepo: partnerRepo,
	}
}

func (s savePartner) Execute(model *SavePartnerModel) (int64, error) {
	if err := model.Validate(); err != nil {
		return 0, err
	}

	entity := &partner.Partner{
		Name:        model.Name,
		AddressID:   model.AddressID,
		Image:       model.Image,
		Phone:       model.Phone,
		Mail:        model.Mail,
		Rating:      model.Rating,
		PartnerType: model.PartnerType,
	}

	lastInsertID, err := s.partnerRepo.Save(entity)

	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

type SavePartnerModel struct {
	Name        string `json:"name"govalid:"req|min:1"`
	AddressID   int64  `json:"address_id"govalid:"req|min:1"`
	Image       string `json:"image"govalid:"req|min:1"`
	Phone       string `json:"phone"govalid:"req|min:1"`
	Mail        string `json:"mail"govalid:"req|min:1"`
	Rating      int64  `json:"rating"govalid:"req|min:1"`
	PartnerType int64  `json:"partner_type"govalid:"req|min:1"`
}

func (p *SavePartnerModel) Validate() error {
	vio, _ := validator.Violations(p)
	if len(vio) > 0 {
		return errors.New(vio[0])
	}
	return nil
}
