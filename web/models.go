package web

import "fmt"

type BeerRequest struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func (b BeerRequest) Validate() error {
	if b.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if b.Brand == "" {
		return fmt.Errorf("brand is empty")
	}
	return nil
}

type ReviewRequest struct {
	BeerID      string `json:"-"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (b ReviewRequest) Validate() error {
	if b.BeerID == "" {
		return fmt.Errorf("beer_id is empty")
	}

	if b.Author == "" {
		return fmt.Errorf("author is empty")
	}

	if b.Description == "" {
		return fmt.Errorf("description is empty")
	}
	return nil
}

type OkResponseMultiple struct {
	Results interface{} `json:"results"`
}

func NewOkMultipleResponse(payload interface{}) *OkResponseMultiple {
	return &OkResponseMultiple{
		Results: payload,
	}
}
