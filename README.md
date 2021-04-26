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
db.users.insert({"name":"Vikash"})
```

## To compile your proto file to Go stubs
1. For Exchange Service:
```protoc -I. proto/exchange.proto --go_out=plugins=grpc:.```
2. For Strategy Service:
```protoc -I. proto/strategy.proto --go_out=plugins=grpc:.```

## Run the gRPC Client

### For Exchange Service
1. To create an exchange
```
go run main.go create -e "Alpaca" -n "Alpaca Exchange" -t "paper_trading" -u "607c612ba353cad51f8103fb" -k "PKOP7ALK9WCI4BH5OX4R" -s "S0NuGhDTNXZ1wp3z9TmuUWhst53ydKAtZ7dtsYhI"
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
go run main.go create -n "Strategy 1" -e "Alpaca" -b 10.0 -s 20.0 -t "5" -p "2%" -v "1%" -c "1%" -m "5%" -z "3%" -f "10000" -u "user1" -d "G1,G2"
```
2. To list all the Strategies by User Id
```
go run main.go listbyuser -u "user1"
```
3. To read a Strategy
```
go run main.go read -i "608597be6854e8575cdb5c3d"
```
4. To Delete a Strategy
```
go run main.go delete -i "608597be6854e8575cdb5c3d"
```
5. To update a Strategy
```
go run main.go update -i "608597be6854e8575cdb5c3d" -n "Strategy 2 Updated" -e "Alpaca" -b 10.0 -s 20.0 -t "5" -p "2%" -v "1%" -c "1%" -m "5%" -z "7%" -f "10000" -d "G1,G2,G3"
```
### Strategy & Deals Bot
1. To Start a Strategy Bot
```
go run main.go startstrategybot -i "60827aeb8babc1d81794a4f7" -d "G1,G2"
```
2. To List all Deals for a User
```
go run main.go listdealsbyuser -u "user1"
```