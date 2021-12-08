# FROM alpine:3.13
FROM golang:latest
LABEL maintainers="Fabian Born" \
      app="Kicker API" \
      description="provide a kicker API"
RUN mkdir /app
WORKDIR /src
# RUN apk add --no-cache bash go git
RUN ls /src
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-yaml/yaml
RUN go build .

COPY ./src  /app
WORKDIR /app
CMD ["/app/api"]

USER root
