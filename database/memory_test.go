package database

import (
	"testing"
	"time"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMemoryDB_AddBeer(t *testing.T) {
	err := memDB.AddBeer(&beers.Beer{
		ID:    beerIDTest,
		Name:  "poretti",
		Brand: "poretti group",
	})
	assert.Nil(t, err)
}

func TestMemoryDB_AddBeers(t *testing.T) {
	err := memDB.AddBeers([]*beers.Beer{
		{
			ID:    uuid.New().String(),
			Name:  "estrella",
			Brand: "estrella group",
		},
		{
			ID:    uuid.New().String(),
			Name:  "estrella galicia",
			Brand: "estrella galicia group",
		},
	})
	assert.Nil(t, err)
}

func TestMemoryDB_GetAllBeers(t *testing.T) {
	beers, err := memDB.GetAllBeers()
	assert.Nil(t, err)
	assert.NotNil(t, beers)

	// print all beers
	for _, b := range beers {
		t.Log(b)
	}
}

func TestMemoryDB_GetBeer(t *testing.T) {
	beer, err := memDB.GetBeer(beerIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, beer)
}

func TestMemoryDB_GetBeers(t *testing.T) {
	beers, err := memDB.GetBeers(beerIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, beers)
	assert.Equal(t, beerIDTest, beers[0].ID)
}

func TestMemoryDB_AddReview(t *testing.T) {
	err := memDB.AddReview(&reviews.Review{
		ID:          reviewIDTest,
		BeerID:      "0cc9b2a0-6ced-44dd-ac08-6b6a1d5646fd",
		Author:      "bomber",
		Description: "very good",
		CreatedAt:   time.Now(),
	})
	assert.Nil(t, err)
}

func TestMemoryDB_GetReview(t *testing.T) {
	review, err := memDB.GetReview(reviewIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, review)
}

func TestMemoryDB_GetAllReviewsByBeerID(t *testing.T) {
	err := memDB.AddReview(&reviews.Review{
		BeerID:      beerIDTest,
		Author:      "bomber",
		Description: "very good",
		CreatedAt:   time.Now(),
	})
	assert.Nil(t, err)
}
