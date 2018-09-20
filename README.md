# beers-api

This is an example of microservice application structure in Golang.
To run the application in docker container

```
make all
docker build -t beers-api .
docker run -d -p 8000:8000 --name beers-api beers-api
```

