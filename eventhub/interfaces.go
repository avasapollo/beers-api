package eventhub

type Service interface {
	PublishBeerEventCreatedV1(event *BeerCreatedV1) error
	PublishReviewCreatedV1(event *ReviewCreatedV1) error
}
