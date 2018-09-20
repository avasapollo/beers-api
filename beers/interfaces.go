package beers

type Service interface {
	AddBeer(beer *Beer) error
	AddBeers(beers []*Beer) error

	GetBeer(beerID string) (*Beer,error)
	GetBeers(ids ...string) ([]*Beer,error)
	GetAllBeers() ([]*Beer,error)
}

type Database interface {
	AddBeer(beer *Beer) error
	AddBeers(beers []*Beer) error

	GetBeer(beerID string) (*Beer,error)
	GetBeers(ids ...string) ([]*Beer,error)
	GetAllBeers() ([]*Beer,error)
}