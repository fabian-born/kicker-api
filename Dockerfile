# FROM alpine:3.13
FROM golang:latest
LABEL maintainers="Fabian Born" \
      app="Kicker API" \
      description="provide a kicker API"
RUN mkdir /app
COPY .  /app
WORKDIR /app
# RUN apk add --no-cache bash go git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/go-yaml/yaml 
CMD ["/app/api"]

USER root
