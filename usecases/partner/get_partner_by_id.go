package partner

import (
	"fmt"
	"ndv/domain/partner"
	partner2 "ndv/infrastructure/partner"
)

type GetPartnerByID interface {
	Execute(id int64) (*partner.Partner, error)
}

type getPartnerByID struct {
	partnerRepo partner2.PartnerDBRepository
}

func NewGetPartnerByID(partnerRepository partner2.PartnerDBRepository) *getPartnerByID {
	return &getPartnerByID{
		partnerRepo: partnerRepository,
	}
}

func (g getPartnerByID) Execute(id int64) (*partner.Partner, error) {

	partner, err := g.partnerRepo.GetByID(id)
	if err != nil {
		fmt.Println("usecase partners: Cannot get partner", id)
		return nil, err
	}

	if partner == nil {
		return nil, ErrNotFound
	}

	return partner, nil
}
