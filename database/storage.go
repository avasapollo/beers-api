package database

import (
	"time"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/google/uuid"
)

func GetInitBeers() []*beers.Beer {
	return []*beers.Beer{
		{
			ID:    "0cc9b2a0-6ced-44dd-ac08-6b6a1d5646fd",
			Name:  "moretti",
			Brand: "moretti group",
		},
		{
			ID:    "e2a3f7d7-fc5b-4809-a27f-146be7ef127e",
			Name:  "peroni",
			Brand: "peroni group",
		},
	}
}

func GetInitReviews() []*reviews.Review {
	return []*reviews.Review{
		{
			ID:          uuid.New().String(),
			BeerID:      "e2a3f7d7-fc5b-4809-a27f-146be7ef127e",
			Author:      "andrea",
			Description: "this beer is amazing!!",
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			BeerID:      "e2a3f7d7-fc5b-4809-a27f-146be7ef127e",
			Author:      "matteo",
			Description: "you should by it man",
			CreatedAt:   time.Now(),
		},
	}
}
