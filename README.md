# beers-api

This is an example of Microservice application structure in Golang.
To run the application in docker container

```
make all
docker build -t beers-api .
docker run -d -p 8000:8000 --name beers-api beers-api
```

This application is created to show a possible structure of a Microservice with Golang, it doesn't contains all the CRUD operations.

The possible endpoint are:

### POST /v1/beers
Add a beer in the memory storage.
Body of the request

```
{
    "name" : "poretti",
    "brand": "poretti-group"
}

```

### POST /v1/beers/{beer_id}/reviews
Add a review on a beer in the memory storage.
Body of the request

```
{
    "name" : "Andrea",
    "description": "This beer is amazing"
}

```

### GET /v1/beers
Get all the beers from the memory storage

### GET /v1/beers/{beer_id}
Get a specific beer from the memory storage

### GET /v1/beers/{beer_id}/reviews
Get the reviews about a specific beer from the memory storage


