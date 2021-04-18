# Algotrading-Golang
An automated tool to sell and buy stock

# Golang Crypto Trading Bot

A golang implementation of trading bot for cryptocurrency exchanges. EXchange used - Alpaka

## Start a MongoDB container using the image from docker hub.
```
docker run -p 27017:27017 --name mongo -d mongo
```
Find the running container and start a Bash session with it using:
```
docker exec -it mongo bash
```
Add some records by connecting to MongoDB Shell:
```
:/# mongo
use mydb
db.availableExchanges.insert({ "name": "Alpaca"})
```

## To compile your proto file to Go stubs in exchange service
```protoc -I. proto/exchange.proto --go_out=plugins=grpc:.```

## To run gRPC Client
1. To create an exchange
```
go run 
main.go create -e "Alpaca" -n "Alpaca Exchange" -t "paper_trading" -u "user2" -k "PKOP7ALK9WCI4BH5OX4R" -s "S0NuGhDTNXZ1wp3z9TmuUWhst53ydKAtZ7dtsYhI"
```
2. To read an exchange by Exchange Id
```
go run main.go read -i "607bd212016f5a5de8117285"
```
3. To list all Exchanges in DB
```
go run main.go list
```
4. To list all Exchanges by User ID
```
go run main.go -u "user2"
```
5. To delete an Exchange
```
go run main.go delete -i "607815df7f51e077fd7ac87a"
```
6. To update an Exchange
```
go run main.go update -i "607815df7f51e077fd7ac87a" -e "Alpaca" -n "Alpaca Exchange" -t "paper_trading" -u "user1" -k "1234trewpoiuyhfjdksa" -s "ertyertyuiolkjhgfdhjkl,mnbvcxdsfghjk"
```