package database

import (
	"github.com/avasapollo/beers-api/beers"
	"github.com/avasapollo/beers-api/reviews"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

const (
	ReviewsMongoCollectionName = "reviews"
	BeersMongoCollectionName   = "beers"

	MongoDatabaseName = "beers-service"
)

type MongoConfig struct {
	MongoUrl string `envconfig:"MONGO_URL"`
}

type MongoDB struct {
	logger  *logrus.Entry
	config  *MongoConfig
	session *mgo.Session
}

func NewMongoDB(le *logrus.Entry, conf *MongoConfig) (Database, error) {
	session, err := mgo.Dial(conf.MongoUrl)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		logger:  le,
		config:  conf,
		session: session,
	}, nil
}

func (m MongoDB) AddBeer(beer *beers.Beer) error {
	session := m.session.Copy()
	defer session.Close()

	return session.DB(MongoDatabaseName).C(BeersMongoCollectionName).Insert(NewBeerDto(
		beer.ID,
		beer.Name,
		beer.Brand,
	))
}
func (m MongoDB) AddBeers(beers []*beers.Beer) error {
	session := m.session.Copy()
	defer session.Close()
	for _, beer := range beers {
		if err := session.DB(MongoDatabaseName).C(BeersMongoCollectionName).Insert(NewBeerDto(
			beer.ID,
			beer.Name,
			beer.Brand,
		)); err != nil {
			m.logger.WithFields(
				logrus.Fields{"error": err.Error(), "beer": beer}).
				Error("couldn't possible add this beer to mongo")
		}

	}
	return nil
}

func (m MongoDB) GetBeer(beerID string) (*beers.Beer, error) {
	session := m.session.Copy()
	defer session.Close()

	beer := new(BeerDto)

	if err := session.DB(MongoDatabaseName).C(BeersMongoCollectionName).Find(bson.M{"_id": beerID}).One(beer); err != nil {
		return nil, err
	}
	return &beers.Beer{
		ID:    beer.ID,
		Name:  beer.Name,
		Brand: beer.Brand,
	}, nil
}
func (m MongoDB) GetBeers(ids ...string) ([]*beers.Beer, error) {
	session := m.session.Copy()
	defer session.Close()

	var slice []*BeerDto

	if err := session.DB(MongoDatabaseName).C(BeersMongoCollectionName).
		Find(bson.M{"_id": bson.M{"$in": ids}}).All(&slice); err != nil {
		return nil, err
	}
	var result []*beers.Beer
	for _, r := range slice {
		result = append(result, &beers.Beer{
			ID:    r.ID,
			Name:  r.Name,
			Brand: r.Brand,
		})
	}
	return result, nil
}

func (m MongoDB) GetAllBeers() ([]*beers.Beer, error) {
	session := m.session.Copy()
	defer session.Close()

	var slice []*BeerDto

	if err := session.DB(MongoDatabaseName).C(BeersMongoCollectionName).
		Find(nil).All(&slice); err != nil {
		return nil, err
	}

	var result []*beers.Beer
	for _, r := range slice {
		result = append(result, &beers.Beer{
			ID:    r.ID,
			Name:  r.Name,
			Brand: r.Brand,
		})
	}
	return result, nil
}

func (m MongoDB) AddReview(review *reviews.Review) error {
	session := m.session.Copy()
	defer session.Close()

	return session.DB(MongoDatabaseName).C(ReviewsMongoCollectionName).Insert(NewReviewDto(
		review.ID,
		review.BeerID,
		review.Author,
		review.Description,
		review.CreatedAt,
	))
}

func (m MongoDB) GetReview(id string) (*reviews.Review, error) {
	session := m.session.Copy()
	defer session.Close()

	review := new(ReviewDto)

	if err := session.DB(MongoDatabaseName).C(ReviewsMongoCollectionName).Find(bson.M{"_id": id}).One(review); err != nil {
		return nil, err
	}
	return &reviews.Review{
		ID:          review.ID,
		BeerID:      review.BeerID,
		Author:      review.Author,
		Description: review.Description,
		CreatedAt:   review.CreatedAt,
	}, nil
}

func (m MongoDB) GetAllReviewsByBeerID(beerID string) ([]*reviews.Review, error) {
	session := m.session.Copy()
	defer session.Close()

	var slice []*ReviewDto

	if err := session.DB(MongoDatabaseName).C(ReviewsMongoCollectionName).
		Find(bson.M{"beer_id": beerID}).All(&slice); err != nil {
		return nil, err
	}

	var result []*reviews.Review
	for _, r := range slice {
		result = append(result, &reviews.Review{
			ID:          r.ID,
			BeerID:      r.BeerID,
			Author:      r.Author,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
		})
	}
	return result, nil
}
