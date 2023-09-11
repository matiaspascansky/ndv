package address

import (
	"errors"
	"fmt"
	addressModel "ndv/domain/address"
	"ndv/infrastructure/address"
	"time"
)

type SaveAddress interface {
	Execute(address *SaveAddressModel) (int64, error)
}

type saveAddress struct {
	repo address.AddressDBRepository
}

type SaveAddressModel struct {
	Name        string    `json:"name" govalid:"req|min:1"`
	Street      string    `json:"street" govalid:"req|min:1"`
	DoorNumber  int       `json:"door_number" govalid:"req|min:1"`
	UserID      int64     `json:"user_id" govalid:"req|min:1"`
	Alias       string    `json:"alias" govalid:"req|min:1"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
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
	fmt.Println("USER ID", address.UserID)
	entity := &addressModel.Address{
		Name:        address.Name,
		Street:      address.Street,
		DoorNumber:  address.DoorNumber,
		Alias:       address.Alias,
		UserID:      address.UserID,
		CreatedAt:   time.Now().UTC(),
		LastUpdated: time.Now().UTC(),
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
