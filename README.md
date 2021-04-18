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

## Run the gRPC Client

### For Exchange Service
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

### For Strategy Service
1. To create a Strategy
```
go run main.go create -a "bot1" -b "alpaca" -c "simple" -d "BTC/USDT" -e "Long" -f "USD" -g 10.00 -i 5.00 -j "Market"
```
2. To list all the Strategies
```
go run main.go list
```
3. To read a Strategy
```
go run main.go read -i "607c2d01ca735a091de58254"
```
4. To Delete a Strategy
```
go run main.go delete -i "607c2d01ca735a091de58254"
```
5. To update a Strategy
```
go run main.go update -1 "607c3a6bca735a091de58255" -a "bot2" -b "alpaca" -c "simple" -d "BTC/SDT" -e "Long" -f "USD" -g 10.00 -i 5.00 -j "Market"
```