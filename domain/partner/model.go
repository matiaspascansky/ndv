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
	FoodType    string    `json:"food_type"`
	IsDeleted   bool      `json:"is_deleted"`
}

// Partner types mapping
var partnerTypes = map[int]string{
	1: "burger",
	2: "pasta",
	3: "bakery",
	4: "sushi",
	5: "pizza",
	6: "wok",
}

func PartnerTypeFromID(ID int) string {
	switch ID {
	case 0:
		{
			return "unknown"
		}
	case 1:
		{
			return "burger"
		}
	case 2:
		{
			return "pasta"
		}
	case 3:
		{
			return "bakery"
		}
	case 4:
		{
			return "sushi"
		}
	case 5:
		{
			return "pizza"
		}
	case 6:
		{
			return "wok"
		}
	}

	return "invalid"
}

// A ESTE MODEL AGREGARLE LO QUE TIENE QUE VER CON ADDRESS Y AGREGADOS
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
