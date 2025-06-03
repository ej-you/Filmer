# -----
# BUILD
# -----

FROM golang:1.24.3-alpine AS build

# set up workdir
WORKDIR /go/src

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# compile app
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./internal ./internal
RUN go build -o ./admin_panel ./cmd/admin_panel/main.go

# ---
# RUN
# ---

FROM alpine:3.22 AS run

WORKDIR /app
# make dir for logs
RUN mkdir /app/logs
ENV LOG_DIR=/app/logs

# copy compiled app web files
COPY --from=build /go/src/admin_panel .
COPY ./web ./web

# run app
CMD ["/bin/sh", "-c", "/app/admin_panel"]
