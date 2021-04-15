# Algotrading-Golang
An automated tool to sell and buy stock

# Golang Crypto Trading Bot

A golang implementation of trading bot for cryptocurrency exchanges. EXchange used - Alpaka

## Start a MongoDB container using the image from docker hub.
```docker run -p 27017:27017 --name mongo -d mongo```
Find the running container and start a Bash session with it using:
```docker exec -it mongo bash```
Add some records by connecting to MongoDB Shell:
```
:/# mongo
use mydb
db.availableExchanges.insert({ "name": "Alpaca"})
```

## To compile your proto file to Go stubs in exchange service
protoc -I. proto/exchange.proto --go_out=plugins=grpc:.