package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// There are integration tests
// There are tests for Memory Database and Mongo Database
// To run the Mongo Database integration tests you need to set the env variable MONGO_URL
// To do the tests I used a docker container with Mongo on this url mongodb://localhost:27017

var (
	le              *logrus.Entry
	memDB           Database
	mongoDB         Database
	mongoConfigured bool
	beerIDTest      = uuid.New().String()
	reviewIDTest    = uuid.New().String()
)

func TestMain(m *testing.M) {
	le = logrus.WithField("app", "testing")
	memDB = NewMemoryDB(
		le)

	if os.Getenv("MONGO_URL") != "" {
		var err error
		mongoDB, err = NewMongoDB(le, &MongoConfig{
			MongoUrl: os.Getenv("MONGO_URL"),
		})
		if err != nil {
			panic(fmt.Sprintf("something wrong in mongo config error: %s", err.Error()))
		}
		mongoConfigured = true
	}
	os.Exit(m.Run())
}
