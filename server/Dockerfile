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
RUN go build -o ./app_server ./cmd/app/main.go

# ---
# RUN
# ---

FROM alpine:3.20 AS run

WORKDIR /app
# make dir for logs
RUN mkdir /app/logs
ENV LOG_DIR /app/logs

# copy compiled file, files for swagger and wallet data
COPY --from=build /go/src/app_server .
COPY --from=build /go/src/migrations ./migrations
COPY --from=build /go/src/docs ./docs

# run app
CMD ["/bin/sh", "-c", "/app/app_server"]
