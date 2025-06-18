# -----
# BUILD
# -----

FROM golang:1.24.4-alpine3.22 AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# compile app
COPY . .
RUN go build -o ./movie_sync ./cmd/movie_sync/main.go

# ---
# RUN
# ---

FROM alpine:3.22 AS run

WORKDIR /app

# copy compiled app
COPY --from=build /go/src/movie_sync .

# run app
CMD ["/bin/sh", "-c", "/app/movie_sync"]
