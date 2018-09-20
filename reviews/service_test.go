package reviews

import (
	"os"
	"testing"
	"time"

	"github.com/avasapollo/beers-api/eventhub"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	le *logrus.Entry
)

func TestMain(m *testing.M) {
	le = logrus.New().WithField("app", "testing")
	os.Exit(m.Run())
}

func TestService_AddReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDatabase(ctrl)
	db.EXPECT().AddReview(gomock.Any()).Return(nil)
	broker := eventhub.NewMockService(ctrl)
	broker.EXPECT().PublishReviewCreatedV1(gomock.Any()).Return(nil)

	beerID := uuid.New().String()

	svc := NewService(le, db, broker)
	err := svc.AddReview(&Review{
		BeerID:      beerID,
		Author:      "andrea",
		Description: "nice beer!",
	})
	assert.Nil(t, err)
}

func TestService_GetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beerID := uuid.New().String()
	reviewID := uuid.New().String()

	db := NewMockDatabase(ctrl)
	db.EXPECT().GetReview(gomock.Any()).Return(&Review{
		ID:          reviewID,
		BeerID:      beerID,
		Author:      "andrea",
		Description: "nice beer!",
		CreatedAt:   time.Now(),
	}, nil)

	broker := eventhub.NewMockService(ctrl)

	svc := NewService(le, db, broker)
	review, err := svc.GetReview(reviewID)

	assert.Nil(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, beerID, review.BeerID)
	assert.Equal(t, reviewID, review.ID)
}

func TestService_GetAllReviewsByBeerID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beerID := uuid.New().String()
	reviewID := uuid.New().String()
	reviewID2 := uuid.New().String()

	db := NewMockDatabase(ctrl)
	db.EXPECT().GetAllReviewsByBeerID(beerID).Return([]*Review{
		{
			ID:          reviewID,
			BeerID:      beerID,
			Author:      "andrea",
			Description: "nice beer!",
			CreatedAt:   time.Now(),
		},
		{
			ID:          reviewID2,
			BeerID:      beerID,
			Author:      "andrea",
			Description: "nice beer!",
			CreatedAt:   time.Now(),
		},
	}, nil)

	svc := NewService(le, db)
	reviews, err := svc.GetAllReviewsByBeerID(beerID)

	assert.Nil(t, err)
	assert.Len(t, reviews, 2)
	assert.Equal(t, beerID, reviews[0].BeerID)
	assert.Equal(t, beerID, reviews[1].BeerID)

	assert.Equal(t, reviewID, reviews[0].ID)
	assert.Equal(t, reviewID2, reviews[1].ID)
}
