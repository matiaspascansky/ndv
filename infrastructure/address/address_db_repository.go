package address

import "ndv/domain/address"

type AddressDBRepository interface {
	GetByUserId(id int64) (*address.Address, error)
	Save(model *address.Address) (int64, error)
}
