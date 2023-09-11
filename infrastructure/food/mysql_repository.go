package food

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"ndv/domain/food"
	"time"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) (*productRepository, error) {
	return &productRepository{db: db}, nil
}

func (p productRepository) GetProductsByPartner(id int64) ([]*food.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) SaveProduct(model *food.Product) (int64, error) {
	now := time.Now().UTC()
	pDto := FromProductTODTO(model)

	sqlStatement := `insert into ndv.product (name, description, max_quantity, age_check,price,image, rating, created_at, last_updated, partner_id) values (?,?,?,?,?,?,?,?,?,?)`
	res, err := p.db.Exec(sqlStatement, pDto.Name, pDto.Description, pDto.MaxQuantity, pDto.AgeCheck, pDto.Price, pDto.Image, pDto.Rating, now, now, pDto.PartnerID)

	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, errors.New("error getting last insert id")
	}
	return lastId, nil
}
