package address

import (
	"errors"
	addressModel "ndv/domain/address"
	"ndv/infrastructure/address"
)

type SaveAddress interface {
	Execute(address *SaveAddressModel) (int64, error)
}

type saveAddress struct {
	repo address.AddressDBRepository
}

type SaveAddressModel struct {
	Name       string `json:"name" govalid:"req|min:1"`
	Street     string `json:"street" govalid:"req|min:1"`
	DoorNumber int    `json:"door_number" govalid:"req|min:1"`
}

func NewSaveAddress(repository address.AddressDBRepository) *saveAddress {
	return &saveAddress{
		repo: repository,
	}
}

func (s saveAddress) Execute(address *SaveAddressModel) (int64, error) {
	if err := address.Validate(); err != nil {
		return 0, err
	}
	entity := &addressModel.Address{
		Name:       address.Name,
		Street:     address.Street,
		DoorNumber: address.DoorNumber,
	}

	lastInsertId, err := s.repo.Save(entity)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (m *SaveAddressModel) Validate() error {
	vio, _ := validator.Violations(m)
	if len(vio) > 0 {
		return errors.New(vio[0])
	}
	return nil
}
