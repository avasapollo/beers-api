package main

import (
	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/database"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/avasapollo/beers-api/web"
	"github.com/sirupsen/logrus"
)

func main() {
	// log
	le := logrus.New().WithField("app", "beers-api")

	// database
	db := database.NewMemoryDB(le, database.GetInitBeers(), database.GetInitReviews())

	// services
	beersSvc := beers.NewService(le, db)
	reviewsSvc := reviews.NewService(le, db)

	web.NewRestApi(le, beersSvc, reviewsSvc).ListenServe()
}
