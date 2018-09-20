FROM alpine:latest

ADD beers-api /usr/local/bin/beers-api

CMD ["/usr/local/bin/beers-api"]