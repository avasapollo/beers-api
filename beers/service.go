package beers

import (
	"fmt"

	"github.com/avasapollo/beers-api/eventhub"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger   *logrus.Entry
	database Database
	broker   eventhub.Service
}

func NewService(le *logrus.Entry, database Database, eventhubSvc eventhub.Service) Service {
	return &service{
		logger:   le,
		database: database,
		broker:   eventhubSvc,
	}
}

func (svc service) AddBeer(beer *Beer) error {
	if beer == nil {
		return fmt.Errorf("nil pointer")
	}
	beer.ID = uuid.New().String()
	if err := svc.database.AddBeer(beer); err != nil {
		return err
	}

	go svc.broker.PublishBeerEventCreatedV1(eventhub.NewBeerCreatedV1(beer.ID, beer.Name, beer.Brand))

	return nil
}

func (svc service) AddBeers(beers []*Beer) error {
	for _, b := range beers {
		if b == nil {
			return fmt.Errorf("there is a nil pointer")
		}
	}

	if err := svc.database.AddBeers(beers); err != nil {
		return err
	}

	for _, b := range beers {
		go svc.broker.PublishBeerEventCreatedV1(eventhub.NewBeerCreatedV1(b.ID, b.Name, b.Brand))
	}

	return nil
}

func (svc service) GetBeer(beerID string) (*Beer, error) {
	return svc.database.GetBeer(beerID)
}

func (svc service) GetBeers(ids ...string) ([]*Beer, error) {
	return svc.database.GetBeers(ids...)
}

func (svc service) GetAllBeers() ([]*Beer, error) {
	return svc.database.GetAllBeers()
}
