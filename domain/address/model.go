package address

import "time"

type Address struct {
	ID          int64
	Name        string
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
	IsDeleted   bool      `json:"is_deleted"`
	Street      string    `json:"street"`
	DoorNumber  int       `json:"door_number"`
	Alias       string    `json:"alias"`
}
