# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN go build -o ./app_client ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.20 AS run

# make dir for logs
RUN mkdir /logs

WORKDIR /app
# copy compiled file, files for swagger and wallet data
COPY --from=build /go/src/app_client .
COPY --from=build /go/src/web ./web

# run app
CMD ["/bin/sh", "-c", "/app/app_client 2>> /logs/error-log.log"]
