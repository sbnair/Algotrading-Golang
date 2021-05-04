package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	orderpb "github.com/vikjdk7/Algotrading-Golang/order-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderServiceServer) ListOrders(req *orderpb.ListOrdersReq, stream orderpb.OrderService_ListOrdersServer) error {
	exchange_id := req.GetExchangeId()
	fmt.Println(exchange_id)

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(exchange_id)
	if err != nil {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied Exchange id to a MongoDB Object Id: %v", err),
		)
	}
	resultReadExchange := exchangedb.FindOne(mongoCtx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataRead := ExchangeItem{}
	// decode and write to dataRead
	if err := resultReadExchange.Decode(&dataRead); err != nil {
		return status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Exchange with Object Id %s: %v", oid, err))
	}

	if dataRead.SelectedExchange == "Alpaca" {
		os.Setenv(common.EnvApiKeyID, dataRead.ApiKey)
		os.Setenv(common.EnvApiSecretKey, dataRead.ApiSecret)
		if dataRead.ExchangeType == "paper_trading" {
			alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
		} else if dataRead.ExchangeType == "live_trading" {
			alpaca.SetBaseUrl("https://api.alpaca.markets")
		}
		alpacaClient := alpaca.NewClient(common.Credentials())
		orderstatus := "all"
		until := time.Now()
		limit := 500
		nested := true
		orders, err := alpacaClient.ListOrders(&orderstatus, &until, &limit, &nested)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		for _, order := range orders {

			stream.Send(&orderpb.ListOrdersRes{
				Order: &orderpb.Order{
					Id:            order.ID,
					ClientOrderId: order.ClientOrderID,
					CreatedAt:     order.CreatedAt.String(),
					UpdatedAt:     order.UpdatedAt.String(),
					SubmittedAt:   order.SubmittedAt.String(),
					FilledAt:      order.FilledAt.String(),
					ExpiredAt:     order.ExpiredAt.String(),
					CanceledAt:    order.CanceledAt.String(),
					FailedAt:      order.FailedAt.String(),
					ReplacedAt:    order.ReplacedAt.String(),
					AssetId:       order.AssetID,
					Symbol:        order.Symbol,
					Exchange:      order.Exchange,
					AssetClass:    order.Class,
					Status:        order.Status,
					ExtendedHours: order.ExtendedHours,
				},
			})
		}
	} else {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot use exchange other than Alpaca"))
	}
	return nil
}

type OrderServiceServer struct{}

type ExchangeItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SelectedExchange string             `bson:"selected_exchange"`
	ExchangeName     string             `bson:"exchange_name"`
	ExchangeType     string             `bson:"exchange_type"`
	UserId           string             `bson:"user_id"`
	ApiKey           string             `bson:"api_key"`
	ApiSecret        string             `bson:"api_secret"`
}

var db *mongo.Client
var exchangedb *mongo.Collection
var orderdb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50054...")

	// Start our listener, 50054 is the default gRPC port
	listener, err := net.Listen("tcp", ":50054")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50054: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)

	// Create ExchangeService type
	srv := &OrderServiceServer{}
	// Register the service with the server
	orderpb.RegisterOrderServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	//Uncomment to run locally
	os.Setenv("MONGODB_URL", "mongodb://127.0.0.1:27017")

	MONGODB_URL := os.Getenv("MONGODB_URL")

	// non-nil empty context
	mongoCtx = context.Background()
	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(MONGODB_URL))
	// Handle potential errors
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	// Bind our collection to our global variable for use in other methods
	mongoDB := db.Database("hedgina_algobot")
	exchangedb = mongoDB.Collection("exchange")
	orderdb = mongoDB.Collection("order")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50054")

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
