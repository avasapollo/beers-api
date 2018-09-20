package database

import (
	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
)

type Database interface {
	AddBeer(beer *beers.Beer) error
	AddBeers(beers []*beers.Beer) error
	GetBeer(beerID string) (*beers.Beer, error)
	GetBeers(ids ...string) ([]*beers.Beer, error)
	GetAllBeers() ([]*beers.Beer, error)

	AddReview(review *reviews.Review) error

	GetReview(id string) (*reviews.Review, error)
	GetAllReviewsByBeerID(beerID string) ([]*reviews.Review, error)
}
