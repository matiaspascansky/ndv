package category

import "time"

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Rating      int64     `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	PartnerType int64     `json:"partner_type"`
	IsDeleted   bool      `json:"is_deleted"`
}
