package database

import (
	"fmt"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/sirupsen/logrus"
)

type memoryDB struct {
	logger  *logrus.Entry
	reviews []*reviews.Review
	beers   []*beers.Beer
}

func NewMemoryDB(le *logrus.Entry, beers []*beers.Beer,
	reviews []*reviews.Review) Database {
	return &memoryDB{
		logger:  le,
		reviews: reviews,
		beers:   beers,
	}
}

func (m *memoryDB) AddBeer(beer *beers.Beer) error {
	m.beers = append(m.beers, beer)
	return nil
}
func (m *memoryDB) AddBeers(beers []*beers.Beer) error {
	m.beers = append(m.beers, beers...)
	return nil
}
func (m memoryDB) GetBeer(beerID string) (*beers.Beer, error) {
	for _, b := range m.beers {
		if b.ID != beerID {
			continue
		}
		return b, nil
	}
	return nil, fmt.Errorf("not found")
}
func (m memoryDB) GetBeers(ids ...string) ([]*beers.Beer, error) {
	var result []*beers.Beer

	for _, id := range ids {
		if b, err := m.GetBeer(id); err == nil {
			result = append(result, b)
		}
	}
	return result, nil
}
func (m memoryDB) GetAllBeers() ([]*beers.Beer, error) {
	return m.beers, nil
}

func (m *memoryDB) AddReview(review *reviews.Review) error {
	m.reviews = append(m.reviews, review)
	return nil
}

func (m memoryDB) GetReview(id string) (*reviews.Review, error) {
	for _, r := range m.reviews {
		if r.ID != id {
			continue
		}
		return r, nil
	}
	return nil, fmt.Errorf("not found")
}

func (m memoryDB) GetAllReviewsByBeerID(beerID string) ([]*reviews.Review, error) {
	var result []*reviews.Review

	for _, r := range m.reviews {
		if r.BeerID != beerID {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}
