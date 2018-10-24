package database

import "time"

type BeerDto struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Brand string `bson:"brand"`
}

func NewBeerDto(id, name, brand string) *BeerDto {
	return &BeerDto{
		ID:    id,
		Name:  name,
		Brand: brand,
	}
}

type ReviewDto struct {
	ID          string    `bson:"_id"`
	BeerID      string    `bson:"beer_id"`
	Author      string    `bson:"author"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"created_at"`
}

func NewReviewDto(id, beerID, author, description string, createdAt time.Time) *ReviewDto {
	return &ReviewDto{
		ID:          id,
		BeerID:      beerID,
		Author:      author,
		Description: description,
		CreatedAt:   createdAt,
	}
}
