# builder image
FROM golang:1.21.0-alpine as gobuilder

RUN apk update && apk add curl make \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN make

# executable image
FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=gobuilder /app/build/linux_amd64/taraxa-snapshotter /taraxa-snapshotter

# Run
CMD [ "/taraxa-snapshotter" ]