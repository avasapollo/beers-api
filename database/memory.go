package database

import (
	"fmt"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/sirupsen/logrus"
)

type memoryDB struct {
	logger  *logrus.Entry
	reviews []*ReviewDto
	beers   []*BeerDto
}

func NewMemoryDB(le *logrus.Entry) Database {
	return &memoryDB{
		logger:  le,
		reviews: []*ReviewDto{},
		beers:   []*BeerDto{},
	}
}

func (m *memoryDB) AddBeer(beer *beers.Beer) error {
	m.beers = append(
		m.beers,
		NewBeerDto(beer.ID, beer.Name, beer.Brand))
	return nil
}

func (m *memoryDB) AddBeers(beers []*beers.Beer) error {
	for _, b := range beers {
		m.beers = append(
			m.beers,
			NewBeerDto(b.ID, b.Name, b.Brand))
	}
	return nil
}

func (m memoryDB) GetBeer(beerID string) (*beers.Beer, error) {
	for _, b := range m.beers {
		if b.ID != beerID {
			continue
		}
		return &beers.Beer{
			ID:    b.ID,
			Name:  b.Name,
			Brand: b.Brand,
		}, nil
	}
	return nil, fmt.Errorf("not found")
}
func (m memoryDB) GetBeers(ids ...string) ([]*beers.Beer, error) {
	var result []*beers.Beer
	for _, id := range ids {
		if b, err := m.GetBeer(id); err == nil {
			result = append(result, &beers.Beer{
				ID:    b.ID,
				Name:  b.Name,
				Brand: b.Brand,
			})
		}
	}
	return result, nil
}
func (m memoryDB) GetAllBeers() ([]*beers.Beer, error) {
	var result []*beers.Beer
	for _, b := range m.beers {
		result = append(result, &beers.Beer{
			ID:    b.ID,
			Name:  b.Name,
			Brand: b.Brand,
		})
	}
	return result, nil
}

func (m *memoryDB) AddReview(review *reviews.Review) error {
	m.reviews = append(m.reviews, NewReviewDto(
		review.ID,
		review.BeerID,
		review.Author,
		review.Description,
		review.CreatedAt))
	return nil
}

func (m memoryDB) GetReview(id string) (*reviews.Review, error) {
	for _, r := range m.reviews {
		if r.ID != id {
			continue
		}
		return &reviews.Review{
			ID:          r.ID,
			BeerID:      r.BeerID,
			Author:      r.Author,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
		}, nil
	}
	return nil, fmt.Errorf("not found")
}

func (m memoryDB) GetAllReviewsByBeerID(beerID string) ([]*reviews.Review, error) {
	var result []*reviews.Review
	for _, r := range m.reviews {
		if r.BeerID != beerID {
			continue
		}
		result = append(result, &reviews.Review{
			ID:          r.ID,
			BeerID:      r.BeerID,
			Author:      r.Author,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
		})
	}
	return result, nil
}
