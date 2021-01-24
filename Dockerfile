# build stage
FROM golang:1.11.2-alpine3.8 AS build-env
ADD . /go/src/giova333/rest-golang-postgres-docker/
RUN cd /go/src/giova333/rest-golang-postgres-docker/ && go build -o user-service

# final stage
FROM alpine:3.8
COPY --from=build-env /go/src/giova333/rest-golang-postgres-docker/ /app/
EXPOSE 8080
ENTRYPOINT ./app/user-service