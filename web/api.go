package web

import (
	"log"
	"net/http"

	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/avasapollo/beers-api/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiRest struct {
	logger  *logrus.Entry
	router  *mux.Router
	beers   beers.Service
	reviews reviews.Service
}

func NewRestApi(le *logrus.Entry, beersSvc beers.Service, reviewsSvc reviews.Service) *ApiRest {
	api := &ApiRest{
		logger:  le,
		router:  mux.NewRouter(),
		beers:   beersSvc,
		reviews: reviewsSvc,
	}
	api.AddHandles()
	return api
}

func (api *ApiRest) AddHandles() {
	api.router.HandleFunc("/v1/beers", api.GetAllBeers).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers/{id}", api.GetBeer).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers/{id}/reviews", api.GetBeerReviews).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers", api.AddBeer).Methods(http.MethodPost)
	api.router.HandleFunc("/v1/beers/{id}/reviews", api.AddBeerReview).Methods(http.MethodPost)
}

func (api ApiRest) GetMuxRouter() *mux.Router {
	return api.router
}

func (api ApiRest) ListenServe() {
	log.Fatal(http.ListenAndServe(":8000", api.router))
}

func (api ApiRest) AddBeer(w http.ResponseWriter, r *http.Request) {
	request := new(BeerRequest)

	// parse body of the request
	if err := utils.ParseJsonBodyRequest(r, request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "the body of the request is not valid")
		return
	}

	// request validation
	if err := request.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	beer := &beers.Beer{
		Name:  request.Name,
		Brand: request.Brand,
	}
	if err := api.beers.AddBeer(beer); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something wrong")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, NewOkResponse(beer))
}

func (api ApiRest) AddBeerReview(w http.ResponseWriter, r *http.Request) {
	request := new(ReviewRequest)

	// parse body of the request
	if err := utils.ParseJsonBodyRequest(r, request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "the body of the request is not valid")
		return
	}
	request.BeerID = mux.Vars(r)["id"]
	// request validation
	if err := request.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	review := &reviews.Review{
		BeerID:      request.BeerID,
		Author:      request.Author,
		Description: request.Description,
	}

	if err := api.reviews.AddReview(review); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something wrong")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, NewOkResponse(review))
}

func (api ApiRest) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	result, err := api.beers.GetAllBeers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, NewOkResponse(result))
}

func (api ApiRest) GetBeer(w http.ResponseWriter, r *http.Request) {
	beerID := mux.Vars(r)["id"]

	beer, err := api.beers.GetBeer(beerID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, NewOkResponse(beer))
}

func (api ApiRest) GetBeerReviews(w http.ResponseWriter, r *http.Request) {
	beerID := mux.Vars(r)["id"]

	reviews, err := api.reviews.GetAllReviewsByBeerID(beerID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, NewOkResponse(reviews))
}
