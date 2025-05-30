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

# compile migrator
COPY ./cmd/migrator .
COPY ./config .
COPY ./internal/pkg/migrate .
RUN go build -o ./migrator ./cmd/migrator/main.go

# compile app
COPY . .
RUN go build -o ./server ./cmd/server/main.go

# ---
# RUN
# ---

FROM alpine:3.21 AS run

WORKDIR /app
# make dir for logs
RUN mkdir /app/logs
ENV LOG_DIR=/app/logs

# copy compiled app and migrator files, files for swagger and migrations
COPY --from=build /go/src/server .
COPY --from=build /go/src/migrator .
COPY --from=build /go/src/migrations ./migrations
COPY --from=build /go/src/docs ./docs

# run app
CMD ["/bin/sh", "-c", "/app/server"]
