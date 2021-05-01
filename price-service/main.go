package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	pricepb "github.com/vikjdk7/Algotrading-Golang/price-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PriceServiceServer) ListMyPositions(req *pricepb.ListMyPositionReq, stream pricepb.PriceService_ListMyPositionsServer) error {
	exchange_id := req.GetExchangeId()
	fmt.Println(exchange_id)

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(exchange_id)
	if err != nil {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied strategy id to a MongoDB ObjectId: %v", err),
		)
	}
	resultReadExchange := exchangedb.FindOne(mongoCtx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataRead := ExchangeItem{}
	// decode and write to dataRead
	if err := resultReadExchange.Decode(&dataRead); err != nil {
		return status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", oid, err))
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
		positions, err := alpacaClient.ListPositions()
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
		//fmt.Println(positions)

		for _, position := range positions {
			avg_entry_price, _ := position.EntryPrice.Float64()
			qty, _ := position.Qty.Float64()
			market_value, _ := position.MarketValue.Float64()
			cost_basis, _ := position.CostBasis.Float64()
			unrealized_pl, _ := position.UnrealizedPL.Float64()
			unrealized_plpc, _ := position.UnrealizedPLPC.Float64()
			current_price, _ := position.CurrentPrice.Float64()
			lastday_price, _ := position.LastdayPrice.Float64()
			change_today, _ := position.ChangeToday.Float64()
			stream.Send(&pricepb.ListMyPositionRes{
				Position: &pricepb.Position{
					AssetId:        position.AssetID,
					Symbol:         position.Symbol,
					Exchange:       position.Exchange,
					AssetClass:     position.Class,
					AccountId:      position.AccountID,
					AvgEntryPrice:  avg_entry_price,
					Qty:            qty,
					Side:           position.Side,
					MarketValue:    market_value,
					CostBasis:      cost_basis,
					UnrealizedPl:   unrealized_pl,
					UnrealizedPlpc: unrealized_plpc,
					CurrentPrice:   current_price,
					LastdayPrice:   lastday_price,
					ChangeToday:    change_today,
				},
			})
		}
	}

	return nil
}

type PriceServiceServer struct{}

type ExchangeItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SelectedExchange string             `bson:"selected_exchange"`
	ExchangeName     string             `bson:"exchange_name"`
	ExchangeType     string             `bson:"exchange_type"`
	UserId           string             `bson:"user_id"`
	ApiKey           string             `bson:"api_key"`
	ApiSecret        string             `bson:"api_secret"`
}

/*
type PositionItem struct {
	AssetId        string  `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Symbol         string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange       string  `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetClass     string  `protobuf:"bytes,4,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	AccountId      string  `protobuf:"bytes,5,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AvgEntryPrice  float64 `protobuf:"fixed64,6,opt,name=avg_entry_price,json=avgEntryPrice,proto3" json:"avg_entry_price,omitempty"`
	Qty            float64 `protobuf:"fixed64,7,opt,name=qty,proto3" json:"qty,omitempty"`
	Side           string  `protobuf:"bytes,8,opt,name=side,proto3" json:"side,omitempty"`
	MarketValue    float64 `protobuf:"fixed64,9,opt,name=market_value,json=marketValue,proto3" json:"market_value,omitempty"`
	CostBasis      float64 `protobuf:"fixed64,10,opt,name=cost_basis,json=costBasis,proto3" json:"cost_basis,omitempty"`
	UnrealizedPl   float64 `protobuf:"fixed64,11,opt,name=unrealized_pl,json=unrealizedPl,proto3" json:"unrealized_pl,omitempty"`
	UnrealizedPlpc float64 `protobuf:"fixed64,12,opt,name=unrealized_plpc,json=unrealizedPlpc,proto3" json:"unrealized_plpc,omitempty"`
	CurrentPrice   float64 `protobuf:"fixed64,13,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	LastdayPrice   float64 `protobuf:"fixed64,14,opt,name=lastday_price,json=lastdayPrice,proto3" json:"lastday_price,omitempty"`
	ChangeToday    float64 `protobuf:"fixed64,15,opt,name=change_today,json=changeToday,proto3" json:"change_today,omitempty"`
}
*/

var db *mongo.Client
var pricedb *mongo.Collection
var exchangedb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50053...")

	// Start our listener, 50053 is the default gRPC port
	listener, err := net.Listen("tcp", ":50053")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50053: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)

	// Create ExchangeService type
	srv := &PriceServiceServer{}
	// Register the service with the server
	pricepb.RegisterPriceServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoCtx = context.Background()
	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
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
	pricedb = mongoDB.Collection("price")
	exchangedb = mongoDB.Collection("exchange")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50053")

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
