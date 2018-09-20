package eventhub

type BeerCreatedV1 struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func NewBeerCreatedV1(id, name, brand string) *BeerCreatedV1 {
	return &BeerCreatedV1{
		ID:    id,
		Name:  name,
		Brand: brand,
	}
}

type ReviewCreatedV1 struct {
	ID          string `json:"id"`
	BeerID      string `json:"beer_id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func NewReviewCreatedV1(id, beerID, author, description string) *ReviewCreatedV1 {
	return &ReviewCreatedV1{
		ID:          id,
		BeerID:      beerID,
		Author:      author,
		Description: description,
	}
}
