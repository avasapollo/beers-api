package eventhub

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	le  *logrus.Entry
	svc Service
)

func TestMain(m *testing.M) {
	le := logrus.New().WithField("app", "testing")
	svc = NewService(le)

	os.Exit(m.Run())
}

func TestService_PublishBeerEventCreatedV1(t *testing.T) {
	err := svc.PublishBeerEventCreatedV1(NewBeerCreatedV1("beer-id-1", "name", "brand"))
	assert.Nil(t, err)
}

func TestService_PublishReviewCreatedV1(t *testing.T) {
	err := svc.PublishReviewCreatedV1(
		NewReviewCreatedV1("review-id-1", "beer-id-1", "author", "description"))
	assert.Nil(t, err)
}
