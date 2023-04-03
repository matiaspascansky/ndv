package partner

import (
	"errors"
	"fmt"
	"ndv/domain/partner"
	partner2 "ndv/infrastructure/partner"
)

var ErrNotFound = errors.New("usecase partner: cannot find any partner")

type GetAllPartners interface {
	Execute() ([]*partner.Partner, error)
}

type getAllPartners struct {
	partnerRepo partner2.PartnerDBRepository
}

func NewGetAllPartners(partnerRepository partner2.PartnerDBRepository) *getAllPartners {
	return &getAllPartners{
		partnerRepo: partnerRepository,
	}
}

func (g getAllPartners) Execute() ([]*partner.Partner, error) {

	partners, err := g.partnerRepo.GetAll()
	if err != nil {
		fmt.Println("usecase partners: Cannot get all partners")
		return nil, err
	}

	if len(partners) == 0 {
		return nil, ErrNotFound
	}

	return partners, nil
}
