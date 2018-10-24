# beers-api

<p align="center">
  <img width="460" height="300" src="https://user-images.githubusercontent.com/22316360/45831785-2694d700-bcf8-11e8-9490-25b7c82e2621.png">
</p>

This is an example of Microservice application structure in Golang.

The dependencies will be installed from `dep`.
You need `dep` to use the Makefile. You can find more info about dep here https://github.com/golang/dep .

To run the application in a docker container you have to run these commands:
```
make all
docker build -t beers-api .
docker run -d -p 8000:8000 --name beers-api beers-api
```
## Environment
The are 2 database implementation: `memory-db` and `mongo-db` you can find both in the database package
To use mongo database you need just to set the environment variable `MONGO_URL` with a mongo url.
For example usually I run a mongo docker container and I expose the mongo port from the container,
if you want to follow this way you can add the url in the docker file:
`ENV MONGO_URL mongodb://localhost:27017`


## REST Api
This application is created just to show a possible structure of a Microservice with Golang.
It doesn't contain all the CRUD operations, this is an example to show just a possible structure about Microservice in Golang.
The server is in listening on port 8000 so `http://localhost:8000`

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
    "author" : "Andrea",
    "description": "This beer is amazing"
}

```

### GET /v1/beers
Get all the beers from the memory storage

### GET /v1/beers/{beer_id}
Get a specific beer from the memory storage

### GET /v1/beers/{beer_id}/reviews
Get the reviews about a specific beer from the memory storage

## Broker
When you will create a beer or a review, the service will publish a message on the broker.

