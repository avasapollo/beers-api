package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	le *logrus.Entry
)

func TestMain(m *testing.M) {
	le = logrus.New().WithField("service", "testing")

	os.Exit(m.Run())
}

func TestApiRest_AddBeer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beersSvc := beers.NewMockService(ctrl)
	beersSvc.EXPECT().AddBeer(&beers.Beer{
		Name:  "poretti",
		Brand: "poretti-group",
	}).Return(nil)

	reviewsSvc := reviews.NewMockService(ctrl)

	apiSvc := NewRestApi(le, beersSvc, reviewsSvc)

	req := httptest.NewRequest("POST", "/v1/beers", bytes.NewBufferString(
		`
			{
				"name" : "poretti",
				"brand": "poretti-group"
			}

	`,
	))
	resp := executeRequest(apiSvc.GetMuxRouter(), req)
	assert.Equal(t, http.StatusCreated, resp.Result().StatusCode)
}

func TestApiRest_AddBeerReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beersSvc := beers.NewMockService(ctrl)
	reviewsSvc := reviews.NewMockService(ctrl)
	reviewsSvc.EXPECT().AddReview(gomock.Any()).Return(nil)
	apiSvc := NewRestApi(le, beersSvc, reviewsSvc)

	req := httptest.NewRequest("POST", "/v1/beers/beer-id-1/reviews", bytes.NewBufferString(
		`
			{
    			"author" : "Andrea",
    			"description": "This beer is amazing"
			}`,
	))
	resp := executeRequest(apiSvc.GetMuxRouter(), req)
	assert.Equal(t, http.StatusCreated, resp.Result().StatusCode)
}

func TestApiRest_GetBeer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beersSvc := beers.NewMockService(ctrl)
	beersSvc.EXPECT().GetBeer("beer-id-1").Return(&beers.Beer{
		ID:    "beer-id-1",
		Name:  "poretti",
		Brand: "poretti-group",
	}, nil)
	reviewsSvc := reviews.NewMockService(ctrl)

	apiSvc := NewRestApi(le, beersSvc, reviewsSvc)

	req := httptest.NewRequest("GET", "/v1/beers/beer-id-1", nil)

	resp := executeRequest(apiSvc.GetMuxRouter(), req)
	assert.Equal(t, http.StatusOK, resp.Result().StatusCode)

	okResp := new(beers.Beer)
	json.NewDecoder(resp.Body).Decode(okResp)

	assert.Equal(t, "beer-id-1", okResp.ID)
	assert.Equal(t, "poretti", okResp.Name)
	assert.Equal(t, "poretti-group", okResp.Brand)
}

func TestApiRest_GetBeerReviews(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	beersSvc := beers.NewMockService(ctrl)
	reviewsSvc := reviews.NewMockService(ctrl)
	reviewsSvc.EXPECT().GetAllReviewsByBeerID("beer-id-1").Return([]*reviews.Review{
		{
			ID:          "review-id-1",
			BeerID:      "beer-id-1",
			Author:      "author",
			Description: "description",
			CreatedAt:   time.Now(),
		}}, nil)

	apiSvc := NewRestApi(le, beersSvc, reviewsSvc)

	req := httptest.NewRequest("GET", "/v1/beers/beer-id-1/reviews", nil)

	resp := executeRequest(apiSvc.GetMuxRouter(), req)
	assert.Equal(t, http.StatusOK, resp.Result().StatusCode)

	okResp := new(OkResponseMultiple)
	json.NewDecoder(resp.Body).Decode(okResp)
	assert.Equal(t, "review-id-1", okResp.Results.([]interface{})[0].(map[string]interface{})["id"])
	assert.Equal(t, "beer-id-1", okResp.Results.([]interface{})[0].(map[string]interface{})["beer_id"])
	assert.Equal(t, "author", okResp.Results.([]interface{})[0].(map[string]interface{})["author"])
	assert.Equal(t, "description", okResp.Results.([]interface{})[0].(map[string]interface{})["description"])
}

func executeRequest(muxRouter *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	muxRouter.ServeHTTP(rr, req)
	return rr
}
