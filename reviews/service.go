package reviews

import (
	"fmt"
	"time"

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
	return service{
		logger:   le,
		database: database,
		broker:   eventhubSvc,
	}
}

func (svc service) AddReview(review *Review) error {
	if review == nil {
		return fmt.Errorf("nil pointer")
	}

	review.ID = uuid.New().String()
	review.CreatedAt = time.Now()

	if err := svc.database.AddReview(review); err != nil {
		return err
	}

	go svc.broker.PublishReviewCreatedV1(
		eventhub.NewReviewCreatedV1(review.ID, review.BeerID, review.Author, review.Description))

	return nil
}

func (svc service) GetReview(id string) (*Review, error) {
	return svc.database.GetReview(id)
}

func (svc service) GetAllReviewsByBeerID(beerID string) ([]*Review, error) {
	return svc.database.GetAllReviewsByBeerID(beerID)
}
