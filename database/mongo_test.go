package database

import (
	"testing"
	"time"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMongoDB_AddBeer(t *testing.T) {
	checkMongoDBTest(t)

	err := mongoDB.AddBeer(&beers.Beer{
		ID:    beerIDTest,
		Name:  "poretti",
		Brand: "poretti group",
	})
	assert.Nil(t, err)

}

func TestMongoDB_AddBeers(t *testing.T) {
	checkMongoDBTest(t)

	err := mongoDB.AddBeers([]*beers.Beer{
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

func TestMongoDB_GetAllBeers(t *testing.T) {
	checkMongoDBTest(t)

	beers, err := mongoDB.GetAllBeers()
	assert.Nil(t, err)
	assert.NotNil(t, beers)

	// print all beers
	for _, b := range beers {
		t.Log(b)
	}
}

func TestMongoDB_GetBeer(t *testing.T) {
	checkMongoDBTest(t)

	beer, err := mongoDB.GetBeer(beerIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, beer)
}

func TestMongoDB_GetBeers(t *testing.T) {
	checkMongoDBTest(t)

	beers, err := mongoDB.GetBeers(beerIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, beers)
	assert.Equal(t, beerIDTest, beers[0].ID)
}

func TestMongoDB_AddReview(t *testing.T) {
	checkMongoDBTest(t)

	err := mongoDB.AddReview(&reviews.Review{
		ID:          reviewIDTest,
		BeerID:      "0cc9b2a0-6ced-44dd-ac08-6b6a1d5646fd",
		Author:      "bomber",
		Description: "very good",
		CreatedAt:   time.Now(),
	})
	assert.Nil(t, err)
}

func TestMongoDB_GetReview(t *testing.T) {
	checkMongoDBTest(t)

	review, err := mongoDB.GetReview(reviewIDTest)
	assert.Nil(t, err)
	assert.NotNil(t, review)
}

func TestMongoDB_GetAllReviewsByBeerID(t *testing.T) {
	checkMongoDBTest(t)

	err := mongoDB.AddReview(&reviews.Review{
		ID:          uuid.New().String(),
		BeerID:      beerIDTest,
		Author:      "bomber",
		Description: "very good",
		CreatedAt:   time.Now(),
	})
	assert.Nil(t, err)
}

func checkMongoDBTest(t *testing.T) {
	if !mongoConfigured {
		t.Skip("mongo url is not configured")
	}
}
