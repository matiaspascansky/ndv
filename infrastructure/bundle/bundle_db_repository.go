package bundle

import "ndv/domain/bundle"

type BundleDBRepository interface {
	GetBundlesByPartner(partnerId int64) ([]*bundle.Bundle, error)
	GetAllBundles() []*bundle.Bundle
	SaveBundle(bundle *bundle.Bundle)
}
