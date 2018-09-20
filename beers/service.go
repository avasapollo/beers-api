package beers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger   *logrus.Entry
	database Database
}

func NewService(le *logrus.Entry, database Database) Service {
	return &service{
		logger:   le,
		database: database,
	}
}

func (svc service) AddBeer(beer *Beer) error {
	if beer == nil {
		return fmt.Errorf("nil pointer")
	}
	beer.ID = uuid.New().String()
	return svc.database.AddBeer(beer)
}

func (svc service) AddBeers(beers []*Beer) error {
	for _, b := range beers {
		if b == nil {
			return fmt.Errorf("there is a nil pointer")
		}
	}
	return svc.database.AddBeers(beers)
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
