package partner

import "ndv/domain/partner"

type PartnerDBRepository interface {
	GetAll() ([]*partner.Partner, error)
	GetByID(int64) (*partner.Partner, error)
	GetByType(int) ([]*partner.Partner, error)
	Save(*partner.Partner) (int64, error)
}
