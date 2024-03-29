package address

import (
	"ndv/domain/address"
	"time"
)

type AddressDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	IsDeleted   bool      `json:"is_deleted"`
	Street      string    `json:"street"`
	DoorNumber  int64     `json:"door_number"`
	Alias       string    `json:"alias"`
	UserID      int64     `json:"user_id"`
}

func FromDTOtoAddress(dto *AddressDTO) *address.Address {
	return &address.Address{
		ID:          dto.ID,
		Name:        dto.Name,
		CreatedAt:   dto.CreatedAt,
		LastUpdated: dto.LastUpdated,
		IsDeleted:   dto.IsDeleted,
		Street:      dto.Street,
		DoorNumber:  int(dto.DoorNumber),
		Alias:       dto.Alias,
	}
}

func FromAddressToDTO(address *address.Address) *AddressDTO {
	return &AddressDTO{
		ID:          address.ID,
		Name:        address.Name,
		CreatedAt:   address.CreatedAt,
		LastUpdated: address.LastUpdated,
		IsDeleted:   address.IsDeleted,
		Street:      address.Street,
		DoorNumber:  int64(address.DoorNumber),
		Alias:       address.Alias,
		UserID:      address.UserID,
	}
}
