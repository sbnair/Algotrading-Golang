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

## To compile your proto file to Go stubs
1. For Exchange Service:
```protoc -I. proto/exchange.proto --go_out=plugins=grpc:.```
2. For Strategy Service:
```protoc -I. proto/strategy.proto --go_out=plugins=grpc:.```
3. For Price Service:
```protoc -I. proto/price.proto --go_out=plugins=grpc:.```

## Run the gRPC Client

### For Exchange Service
1. To create an exchange
```
go run main.go create -e "Alpaca" -n "Alpaca Exchange" -t "paper_trading" -u "60865a63a6116f40ab12d863" -k "PKOP7ALK9WCI4BH5OX4R" -s "S0NuGhDTNXZ1wp3z9TmuUWhst53ydKAtZ7dtsYhI"
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
go run main.go read -i "6086c3b8e33d324e030c78ed"
```
4. To Delete a Strategy
```
go run main.go delete -i "60867adb76c5d62a565126c6"
```
5. To update a Strategy
```
go run main.go update -i "6086c3b8e33d324e030c78ed" -n "Strategy 1 Updated" -e "Alpaca" -b 10.0 -s 20.0 -t "5" -p "2%" -v "1%" -c "1%" -m "5%" -z "7%" -f "10000" -d "G1,G2,G3"
```
### Strategy & Deals Bot
1. To Start a Strategy Bot
```
go run main.go startstrategybot -i "6086c3b8e33d324e030c78ed" -d "G1,G2"
```
2. To List all Deals for a User
```
go run main.go listdealsbyuser -u "user1"
```

### For Price Service
1. List My Positions for an exchange
```
go run main.go listmypositions -e "6086b79f8ec6cae85d53584a"
```
## Run the User Authentication Service
1. Git Clone
2. go build -o new -v
3. ```go run main.go```. The app is then available on http://localhost:8000
4. For User Signup, POST to http://localhost:8000/users/signup with body
```
{
    "first_name":"Neha",
    "last_name":"Kumari",
    "email":"neha190495@gmail.com",
    "password":"algobot1",
    "phone":"7416516791"
}
```
5. For User Login, POST to http://localhost:8000/users/login with body
```
{
    "email":"neha190495@gmail.com",
    "password":"algobot1"
}
```
6. To access API's: GET http://localhost:8000/api-1 with header "token" and value token retrieved from login api call.