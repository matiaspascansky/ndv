package partner

import (
	"fmt"
	"ndv/domain/partner"
	partner2 "ndv/infrastructure/partner"
)

type GetPartnerByType interface {
	Execute(id int) ([]*partner.Partner, error)
}

type getPartnerByType struct {
	partnerRepo partner2.PartnerDBRepository
}

func NewGetPartnerByType(partnerRepository partner2.PartnerDBRepository) *getPartnerByType {
	return &getPartnerByType{
		partnerRepo: partnerRepository,
	}
}

func (g getPartnerByType) Execute(id int) ([]*partner.Partner, error) {

	partners, err := g.partnerRepo.GetByType(id)
	if err != nil {
		fmt.Println("usecase partners: Cannot get partner", id)
		return nil, err
	}

	if len(partners) == 0 {
		return nil, ErrNotFound
	}

	return partners, nil
}
