package partner

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"ndv/domain/partner"
	"time"
)

type partnerRepository struct {
	db *sqlx.DB
}

func NewPartnerRepository(db *sqlx.DB) (*partnerRepository, error) {
	return &partnerRepository{db: db}, nil
}

func (p partnerRepository) GetAll() ([]*partner.Partner, error) {
	sqlStatement := `select id, name, address_id, image, phone, mail, rating,created_at,last_updated,partner_type from ndv.partners`
	results := &sql.Rows{}
	results, err := p.db.Query(sqlStatement)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	out := make([]*partner.Partner, 0, 2000)
	var i int

	for results.Next() {
		var partnerDTO PartnerDTO
		err = results.Scan(
			&partnerDTO.ID,
			&partnerDTO.Name,
			&partnerDTO.AddressID,
			&partnerDTO.Image,
			&partnerDTO.Phone,
			&partnerDTO.Mail,
			&partnerDTO.Rating,
			&partnerDTO.CreatedAt,
			&partnerDTO.LastUpdated,
			&partnerDTO.PartnerType,
		)
		if err != nil {
			return nil, err
		}
		partner := FromDtoToPartner(&partnerDTO)
		out = append(out, partner)
		i++
	}
	return out, nil
}

func (p partnerRepository) GetByType(partnerType int) ([]*partner.Partner, error) {
	sqlStatement := `select id, name, address_id, image, phone, mail, rating,created_at,last_updated,partner_type from ndv.partners where partner_type = ?`
	results := &sql.Rows{}
	results, err := p.db.Query(sqlStatement, partnerType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	out := make([]*partner.Partner, 0, 2000)
	var i int

	for results.Next() {
		var partnerDTO PartnerDTO
		err = results.Scan(
			&partnerDTO.ID,
			&partnerDTO.Name,
			&partnerDTO.AddressID,
			&partnerDTO.Image,
			&partnerDTO.Phone,
			&partnerDTO.Mail,
			&partnerDTO.Rating,
			&partnerDTO.CreatedAt,
			&partnerDTO.LastUpdated,
			&partnerDTO.PartnerType,
		)
		if err != nil {
			return nil, err
		}
		partner := FromDtoToPartner(&partnerDTO)
		out = append(out, partner)
		i++
	}
	return out, nil
}

func (p partnerRepository) GetByID(id int64) (*partner.Partner, error) {

	sqlStatement := `select id, name, address_id, image, phone, mail, rating,created_at,last_updated,partner_type from ndv.partners where id = ? and is_deleted = 0`
	row := p.db.QueryRow(sqlStatement, id)
	var pDTO PartnerDTO
	err := row.Scan(&pDTO.ID, &pDTO.Name, &pDTO.AddressID, &pDTO.Image, &pDTO.Phone, &pDTO.Mail, &pDTO.Rating, &pDTO.CreatedAt, &pDTO.LastUpdated, &pDTO.PartnerType)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	partner := FromDtoToPartner(&pDTO)
	return partner, nil

}

func (p partnerRepository) Save(model *partner.Partner) (int64, error) {
	now := time.Now().UTC()

	pDto := FromPartnerToDTO(model)
	sqlStatement := `insert into ndv.partners (name, address_id, image, phone,mail,rating, created_at,last_updated,is_deleted,partner_type) values (?, ?, ?, ?, ?, ?, ?,?,?,?)
`
	res, err := p.db.Exec(sqlStatement, pDto.Name, pDto.AddressID, pDto.Image, pDto.Phone, pDto.Mail, pDto.Rating, now, now, false, pDto.PartnerType)

	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New("partner repository: error getting last insert id")
	}
	return lastId, nil
}
