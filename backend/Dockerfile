## Build
FROM golang:1.16-alpine AS build

# Copy our go.mod and go.sum files into /app
COPY go.mod go.sum /go/src/github.com/sammsaski/wishlist-golang/backend/

# Change working directory
WORKDIR /go/src/github.com/sammsaski/wishlist-golang/backend

# Install dependencies
RUN go mod download

# Copy source code into image
COPY . /go/src/github.com/sammsaski/wishlist-golang/backend

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/wishlist github.com/sammsaski/wishlist-golang/backend


## Deploy
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=build /go/src/github.com/sammsaski/wishlist-golang/backend/build/wishlist /usr/bin/wishlist/backend_executable
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/wishlist/backend_executable"]
