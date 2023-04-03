package partner

import "time"

type Partner struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	AddressID   int64     `json:"address_id"`
	Image       string    `json:"image"`
	Phone       string    `json:"phone"`
	Mail        string    `json:"mail"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerType int64     `json:"partner_type"`
	IsDeleted   bool      `json:"is_deleted"`
}

//A ESTE MODEL AGREGARLE TODO LO QUE TIENE QUE VER CON ADDRESS Y AGREGADOS
type PartnerDetail struct {
	ID   int64
	Name string
	//Address pointer to Address type. TBD
	AddressID   int64     `json:"address_id"`
	Image       string    `json:"image"`
	Phone       string    `json:"phone"`
	Mail        string    `json:"mail"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerType int64     `json:"partner_type"`
	IsDeleted   bool      `json:"is_deleted"`
}
