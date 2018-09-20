package eventhub

import "github.com/sirupsen/logrus"

type service struct {
	logger *logrus.Entry
}

func NewService(le *logrus.Entry) Service {
	return &service{
		logger: le,
	}
}

func (svc service) PublishBeerEventCreatedV1(event *BeerCreatedV1) error {
	svc.logger.Infof("message: %v", event)
	return nil
}

func (svc service) PublishReviewCreatedV1(event *ReviewCreatedV1) error {
	svc.logger.Infof("message: %v", event)
	return nil
}
