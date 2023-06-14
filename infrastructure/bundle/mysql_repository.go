package bundle

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"ndv/domain/bundle"
)

type bundleRepository struct {
	db *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) (*bundleRepository, error) {
	return &bundleRepository{db: db}, nil
}

func (b bundleRepository) GetBundlesByPartner(partnerId int64) ([]*bundle.Bundle, error) {
	sqlStatement := `select id,
		name,
		image,
		rating,
		created_at,
		last_updated,
		partner_id,
		is_deleted,
		from ndv.bundle
		where id = ?`
	results, err := b.db.Query(sqlStatement, partnerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	out := make([]*bundle.Bundle, 0, 2000)
	var i int

	for results.Next() {
		var bundleDTO BundleDTO
		err = results.Scan(
			&bundleDTO.ID,
			&bundleDTO.Name,
			&bundleDTO.Image,
			&bundleDTO.Rating,
			&bundleDTO.CreatedAt,
			&bundleDTO.LastUpdated,
			&bundleDTO.PartnerID,
			&bundleDTO.IsDeleted,
		)
		if err != nil {
			return nil, err
		}
		bundle := FromDtoToBundle(&bundleDTO)
		out = append(out, bundle)
		i++
	}
	return out, nil
}

func (b bundleRepository) GetAllBundles() []*bundle.Bundle {
	//TODO implement me
	panic("implement me")
}

func (b bundleRepository) SaveBundle(bundle *bundle.Bundle) {
	//TODO implement me
	panic("implement me")
}
