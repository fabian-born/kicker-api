# FROM alpine:3.13
FROM golang:latest
LABEL maintainers="Fabian Born" \
      app="Kicker API" \
      description="provide a kicker API"
RUN mkdir /app
WORKDIR /app
COPY . /app
# RUN apk add --no-cache bash go git
RUN ls /app
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-yaml/yaml
RUN go build -o /app/api ./src

EXPOSE 8084

CMD ["/app/api"]

USER root

