package beers

import (
	"os"
	"testing"

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

func TestService_AddBeer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDatabase(ctrl)
	db.EXPECT().AddBeer(gomock.Any()).Return(nil)

	svc := NewService(le, db)
	err := svc.AddBeer(&Beer{
		Name:  "moretti",
		Brand: "moretti group",
	})
	assert.Nil(t, err)
}

func TestService_AddBeers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := NewMockDatabase(ctrl)
	db.EXPECT().AddBeers(gomock.Any()).Return(nil)

	svc := NewService(le, db)
	err := svc.AddBeers([]*Beer{
		{
			Name:  "moretti",
			Brand: "moretti group",
		},
		{
			Name:  "peroni",
			Brand: "peroni group",
		},
	})
	assert.Nil(t, err)
}

func TestService_GetBeer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beerID := uuid.New().String()

	db := NewMockDatabase(ctrl)
	db.EXPECT().GetBeer(beerID).Return(&Beer{
		ID:    beerID,
		Name:  "moretti",
		Brand: "moretti group",
	}, nil)

	svc := NewService(le, db)
	beer, err := svc.GetBeer(beerID)
	assert.Nil(t, err)
	assert.Equal(t, beerID, beer.ID)
}

func TestService_GetBeers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beerID := uuid.New().String()
	beerID2 := uuid.New().String()

	db := NewMockDatabase(ctrl)
	db.EXPECT().GetBeers(beerID, beerID2).Return([]*Beer{
		{
			ID:    beerID,
			Name:  "moretti",
			Brand: "moretti group",
		},
		{
			ID:    beerID2,
			Name:  "peroni",
			Brand: "peroni group",
		},
	}, nil)

	svc := NewService(le, db)
	beers, err := svc.GetBeers(beerID, beerID2)
	assert.Nil(t, err)
	assert.Equal(t, beerID, beers[0].ID)
	assert.Equal(t, beerID2, beers[1].ID)
}

func TestService_GetAllBeers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beerID := uuid.New().String()
	beerID2 := uuid.New().String()

	db := NewMockDatabase(ctrl)
	db.EXPECT().GetAllBeers().Return([]*Beer{
		{
			ID:    beerID,
			Name:  "moretti",
			Brand: "moretti group",
		},
		{
			ID:    beerID2,
			Name:  "peroni",
			Brand: "peroni group",
		},
	}, nil)

	svc := NewService(le, db)
	beers, err := svc.GetAllBeers()
	assert.Nil(t, err)
	assert.Equal(t, beerID, beers[0].ID)
	assert.Equal(t, beerID2, beers[1].ID)
}
