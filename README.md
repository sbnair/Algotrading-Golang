# Algotrading-Golang
An automated tool to sell and buy stock

# Golang Crypto Trading Bot

A golang implementation of trading bot for cryptocurrency exchanges. EXchange used - Alpaka

## Locally Start a MongoDB container using the image from docker hub.
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
4. For Order Service:
```protoc -I. proto/order.proto --go_out=plugins=grpc:.```
## Run the gRPC Client

### For Exchange Service
1. To create an exchange
```
go run main.go create -e "Alpaca" -n "Alpaca Exchange 1" -t "paper_trading" -u "57bf425a34ce5ee85891b914" -k "PKI0XITWRU9E47IEOLV3" -s "d2EdBHfpzZmMABVuJey2weVu35mImiMTGO6pPXO2"
```
2. To read an exchange by Exchange Id
```
go run main.go read -i "60955018a020171d0214f053"
```
3. To list all Exchanges in DB
```
go run main.go list
```
4. To list all Exchanges by User ID
```
go run main.go -u "57bf425a34ce5ee85891b914"
```
5. To delete an Exchange
```
go run main.go delete -i "607815df7f51e077fd7ac87a"
```
6. To update an Exchange
```
go run main.go update -i "607815df7f51e077fd7ac87a" -e "Alpaca" -n "Alpaca Exchange" -t "paper_trading" -u "57bf425a34ce5ee85891b914" -k "PKI0XITWRU9E47IEOLV3" -s "d2EdBHfpzZmMABVuJey2weVu35mImiMTGO6pPXO2"
```

### For Strategy Service
1. To create a Strategy
```
go run main.go create -n "Strategy 1" -e "609e8aa2f5d820ae105be543" -b 10.0 -s 20.0 -t "5" -w "2" -p "2%" -v "1%" -c "1%" -m "5%" -z "3%" -f "10000" -u "57bf425a34ce5ee85891b914" -d "SAP,GOOGL"
```
2. To list all the Strategies by User Id
```
go run main.go listbyuser -u "57bf425a34ce5ee85891b914"
```
3. To read a Strategy by Strategy Id
```
go run main.go read -i "6086c3b8e33d324e030c78ed"
```
4. To Delete a Strategy by Strategy Id
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
go run main.go startstrategybot -i "608d6d04c3271e63ee719eab" -d "SAP"
```
2. To List all Deals for a User
```
go run main.go listdealsbyuser -u "57bf425a34ce5ee85891b914"
```

### For Price Service
1. List My Positions for an exchange
```
go run main.go listmypositions -e "60955018a020171d0214f053"
```
2. List All Assets
```
go run main.go listassets
```
3. List Asset by Symbol
```
go run main.go listassetbysymbol -s "GOOGL"
```
4. List Asset by Name
```
go run main.go listassetbyname -n "SAP SE"
```

### For Order Service
1. Create an order of type "market"
```
go run main.go placeorder -e "609566b5a020171d0214f057" -a "SAP" -q 0.1 -s "buy" -t "market" -f "day"
```
2. Create an order of type "limit"
```
go run main.go placeorder -e "6094d1ee9a18b28d455491fb" -a "SAP" -q 1 -s "buy" -t "limit" -f "day" -l 200.0
```
3. List all orders by User
```
go run main.go listordersbyuser -u "57bf425a34ce5ee85891b914"
```
4. Cancel an order
```
go run main.go cancelorder -o "4db7a9b2-0e0d-4674-b442-2ae73be3a786" -e "60955018a020171d0214f053"
```

### For Event History Service
1. List all Events History for Exchange
```
go run main.go listeventhistoryexchange -u "57bf425a34ce5ee85891b914"
```
2. List all Events History for Strategy
```
go run main.go listeventhistorystrategy -u "57bf425a34ce5ee85891b914"
```

### For CEP Rule Service
1. Calculate Funds allocation for Strategy
```
go run main.go calculatefundallocation -b 100.0 -s 3.0 -v 2.0 -a 3.0 -d 2
```
## Run the User Authentication Service
1. Git Clone
2. go build -o new -v
3. ```go run main.go```. The app is then available on http://localhost:8000
4. For User Signup, POST to http://localhost:8000/users/signup with body
```
{
    "first_name":"Vikash",
    "last_name":"Bharti",
    "email":"vikjdk7@gmail.com",
    "password":"algobot1",
    "phone":"7779920620"
}
```
5. For User Login, POST to http://localhost:8000/users/login with body
```
{
    "email":"vikjdk7@gmail.com",
    "password":"algobot1"
}
```
6. To access API's: GET http://localhost:8000/api-1 with header "token" and value token retrieved from login api call.

## Run assets-job to fill asset information from Alpaca into Mongo

### Pre-Requisites
1. Create a admin user in DB with first_name=admin and last_name=user
2. Create an exchange for this user with valid API-KEY & SECRET
### Run the App/Job
1. To run locally, uncomment line number 146 in main.go
2. Run the app using ```go run main.go```
### Build its Docker Image
1. Build the docker image using
```
docker build -t assets-job .
```
## Docker Operations
1. Build a docker image:
```
docker build -t price-service .
```
2. Run the docker container
```
docker run -p 50053:50053 price-service
```
3. Tag the docker Image
```
docker tag price-service:latest vikash99/price-service
```
4. Push the docker image to dockerhub
```
docker push vikash99/price-service
```
## Kubernetes Deployments

### Setup EFK Stack on Kubernetes Cluster for Logging
1. Create a namespace
```
kubectl apply -f kubernetes-deployments/logging/namespace.yaml
```
If the namespace name is changed, change it in other elastisearch.yaml,kibana.yaml & fluentd.yaml as well.

2. Setup Elastisearch <br />
Enter your cluster name at line 55 & storage class name at line 92
```
kubectl apply -f kubernetes-deployments/logging/elastisearch.yaml
```
3. Setup Kibana
```
kubectl apply -f kubernetes-deployments/logging/kibana.yaml
```
4. Setup fluentd DeamonSet
```
kubectl apply -f kubernetes-deployments/logging/fluentd.yaml
```
5. Open Kibana dashboard 
```
kubectl port-forward $kibanaPodName 5601:5601 -n logging
```

### Create a standalone mongodb statefulset
1. To create a standalone mongodb
```
kubectl apply -f kubernetes-deployments/mongodb/mongodb.yaml
```
2. To exec into mongodb from CLI
```
kubectl -n hedgina exec -it mongodb-0 -- mongo mongodb://mongoadmin:mongopassword@mongodb-0.database:27017/?authSource=admin
```
or
```
kubectl -n hedgina exec -it mongodb-0 -- mongo mongodb://mongodb-0.database:27017 --username mongoadmin --password mongopassword
```
3. To connect to mongodb from inside the cluster use the connection String: ```mongodb://mongoadmin:mongopassword@mongodb-0.database:27017/?authSource=admin``` using standard connection string format: ```mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]```

4. Example of mongodb connection string for multiple replicaset ```mongodb://mongoadmin:mongopassword@mongodb-0.database:27017,mongodb-1.database:27017,mongodb-2.database:27017/?authSource=admin```

5. To connect to mongodb running on k8s cluster from your local, port-forward the mongodb pod to localhost:27017 using ```kubectl -n hedgina port-forward mongodb-0 27017:27017```. To stop, hit Ctrl+C