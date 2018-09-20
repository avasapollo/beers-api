package reviews

import "time"

type Review struct {
	ID          string    `json:"id"`
	BeerID      string    `json:"beer_id"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
