package reviews

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger   *logrus.Entry
	database Database
}

func NewService(le *logrus.Entry, database Database) Service {
	return service{
		logger:   le,
		database: database,
	}
}

func (svc service) AddReview(review *Review) error {
	if review == nil {
		return fmt.Errorf("nil pointer")
	}

	review.ID = uuid.New().String()
	review.CreatedAt = time.Now()

	return svc.database.AddReview(review)
}

func (svc service) GetReview(id string) (*Review, error) {
	return svc.database.GetReview(id)
}

func (svc service) GetAllReviewsByBeerID(beerID string) ([]*Review, error) {
	return svc.database.GetAllReviewsByBeerID(beerID)
}
