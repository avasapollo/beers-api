package main

import (
	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/database"
	"github.com/avasapollo/beers-api/eventhub"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/avasapollo/beers-api/web"
	"github.com/sirupsen/logrus"
)

func main() {
	// log
	le := logrus.New().WithField("app", "beers-api")

	// database
	db := database.NewMemoryDB(le, database.GetInitBeers(), database.GetInitReviews())
	brokerSvc := eventhub.NewService(le)

	// services
	beersSvc := beers.NewService(le, db, brokerSvc)
	reviewsSvc := reviews.NewService(le, db, brokerSvc)

	web.NewRestApi(le, beersSvc, reviewsSvc).ListenServe()
}
