package reviews

type Service interface {
	AddReview(review *Review) error

	GetReview(id string) (*Review, error)
	GetAllReviewsByBeerID(beerID string) ([]*Review, error)
}

type Database interface {
	AddReview(review *Review) error

	GetReview(id string) (*Review, error)
	GetAllReviewsByBeerID(beerID string) ([]*Review, error)
}
