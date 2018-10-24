package main

import (
	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/database"
	"github.com/avasapollo/beers-api/eventhub"
	"github.com/avasapollo/beers-api/global"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/avasapollo/beers-api/web"
	"github.com/sirupsen/logrus"
)

// Env `MONGO_URL`: allow you to use mongo database in the application,
// the mongo url should be something like this mongodb://localhost:27017
// if it is not configured the application will configured the memory database
func main() {
	// log
	le := logrus.New().WithField("app", "beers-api")

	// get app config
	config, err := global.NewAppConfig()
	if err != nil {
		le.Fatal(err)
	}

	// datbase
	var db database.Database
	// check if mongo url is configured will be set mongo database, if it is not set will be set memory database
	if !config.MongoUrlIsSet() {
		// memory database
		db = database.NewMemoryDB(le)
		le.Info("memory database is configured")
	} else {
		// mongo database
		if db, err = database.NewMongoDB(le, config.MongoConfig); err != nil {
			le.WithField("error", err.Error()).Fatal("mongo is down")
		}
		le.Info("mongo database is configured")
	}

	le.WithFields(logrus.Fields{
		"app_name": config.AppName,
		"env":      config.Env,
		"api_port": config.Port,
	}).Info("application starting...")
	// broker
	brokerSvc := eventhub.NewService(le)

	// services
	beersSvc := beers.NewService(le, db, brokerSvc)
	reviewsSvc := reviews.NewService(le, db, brokerSvc)

	web.NewRestApi(le, beersSvc, reviewsSvc).ListenServe(config.Port)
}
