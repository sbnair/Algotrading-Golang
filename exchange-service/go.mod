module github.com/vikjdk7/Algotrading-Golang/exchange-service

go 1.16

//replace github.com/vikjdk7/Algotrading-Golang/exchange-service => ./exchange-service

require (
	github.com/alpacahq/alpaca-trade-api-go v1.8.1
	github.com/golang/protobuf v1.5.2
	go.mongodb.org/mongo-driver v1.5.1
	google.golang.org/grpc v1.37.0
)
