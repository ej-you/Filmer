# -----
# BUILD
# -----

FROM golang:1.24.2-alpine AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY ./internal ./internal
COPY ./config ./config
COPY ./cmd ./cmd

# compile app
RUN go build -o ./app_client ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.21 AS run

# make dir for logs
RUN mkdir /logs

WORKDIR /app
# copy compiled file, templates and static files
COPY --from=build /go/src/app_client .
COPY ./web ./web

# run app
CMD ["/bin/sh", "-c", "/app/app_client 2>> /logs/error-log.log"]
