# beers-api

<p align="center">
  <img width="460" height="300" src="https://user-images.githubusercontent.com/22316360/45831785-2694d700-bcf8-11e8-9490-25b7c82e2621.png">
</p>


This is an example of Microservice application structure in Golang.
To run the application in a docker container you have to run these commands:

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

