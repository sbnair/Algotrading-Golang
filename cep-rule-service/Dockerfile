FROM golang:1.16-alpine3.13 AS build_base
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/cep-rule-service
COPY . .

RUN go mod download

# Build the Go app
RUN go build -o ./out/cep-rule-service .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/cep-rule-service/out/cep-rule-service /app/cep-rule-service

# Set Env Variable
ENV MONGODB_URL=mongodb://mongoadmin:mongopassword@mongodb-0.database:27017/?authSource=admin

# Run the binary program produced by `go install`
CMD ["/app/cep-rule-service"]