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
	return &ApiRest{
		logger:  le,
		router:  mux.NewRouter(),
		beers:   beersSvc,
		reviews: reviewsSvc,
	}
}

func (api ApiRest) ListenServe() {
	api.router.HandleFunc("/v1/beers", api.getAllBeers).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers/{id}", api.getBeer).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers/{id}/reviews", api.getBeerReviews).Methods(http.MethodGet)
	api.router.HandleFunc("/v1/beers", api.addBeer).Methods(http.MethodPost)
	api.router.HandleFunc("/v1/beers/{id}/reviews", api.addBeer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", api.router))
}

func (api ApiRest) addBeer(w http.ResponseWriter, r *http.Request) {
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
		Brand: request.Name,
	}
	if err := api.beers.AddBeer(beer); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something wrong")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, beer)
}

func (api ApiRest) addBeerReview(w http.ResponseWriter, r *http.Request) {
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

	if err := api.reviews.AddReview(&reviews.Review{
		BeerID:      request.BeerID,
		Author:      request.Author,
		Description: request.Description,
	}); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something wrong")
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, review)
}

func (api ApiRest) getAllBeers(w http.ResponseWriter, r *http.Request) {
	result, err := api.beers.GetAllBeers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, result)
}

func (api ApiRest) getBeer(w http.ResponseWriter, r *http.Request) {
	beerID := mux.Vars(r)["id"]

	beer, err := api.beers.GetBeer(beerID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, beer)
}

func (api ApiRest) getBeerReviews(w http.ResponseWriter, r *http.Request) {
	beerID := mux.Vars(r)["id"]

	reviews, err := api.reviews.GetAllReviewsByBeerID(beerID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, reviews)
}
