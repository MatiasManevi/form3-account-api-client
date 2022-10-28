FROM golang:1.19.1-alpine

RUN apk add build-base

RUN export CGO_ENABLED=0
RUN export GO111MODULE=off

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./
# copy dummy data used in test
COPY testdata/ ./testdata

ENTRYPOINT ["go", "test", "-v", "./..."]