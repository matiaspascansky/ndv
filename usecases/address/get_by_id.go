package address

import (
	"errors"
	"ndv/domain/address"
	address2 "ndv/infrastructure/address"
)

var ErrNotFound = errors.New("usecases address: address not found")

type GetAddressByID interface {
	Execute(id int64) (*address.Address, error)
}

type getAddressByID struct {
	addressRepository address2.AddressDBRepository
}

func NewGetAddressByID(repository address2.AddressDBRepository) *getAddressByID {
	return &getAddressByID{addressRepository: repository}
}

func (g getAddressByID) Execute(id int64) (*address.Address, error) {
	if id == 0 {
		return nil, ErrNotFound
	}

	entity, err := g.addressRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, ErrNotFound
	}

	return entity, nil

}
