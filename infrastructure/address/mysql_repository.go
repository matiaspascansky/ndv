package address

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"ndv/domain/address"
	"time"
)

type addressRepository struct {
	db *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) (*addressRepository, error) {
	return &addressRepository{db: db}, nil
}

func (a addressRepository) GetByUserId(id int64) (*address.Address, error) {
	fmt.Println("repository: getting address by id:  ", id)
	sqlStatement := `select id,
		name,
		created_at,
		last_updated,
		is_deleted,
		street,
		door_number
		from ndv.address
	    where user_id = ?;`
	row := a.db.QueryRow(sqlStatement, id)

	var aDTO AddressDTO
	err := row.Scan(&aDTO.ID, &aDTO.Name, &aDTO.CreatedAt, &aDTO.LastUpdated, &aDTO.IsDeleted, &aDTO.Street, &aDTO.DoorNumber)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	address := FromDTOtoAddress(&aDTO)
	return address, nil
}

func (a addressRepository) Save(model *address.Address) (int64, error) {
	now := time.Now().UTC()
	aDto := FromAddressToDTO(model)
	sqlStatement := `insert into ndv.address (name, created_at, last_updated, is_deleted,street,door_number, user_id, alias) values (?,?,?,?,?,?,?,?)`
	res, err := a.db.Exec(sqlStatement, aDto.Name, now, now, false, aDto.Street, aDto.DoorNumber, aDto.UserID, aDto.Alias)

	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New("error getting last insert id")
	}
	return lastId, nil
}
