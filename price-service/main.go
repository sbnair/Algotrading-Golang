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

func (s *PriceServiceServer) ListAssets(req *pricepb.ListAssetsReq, stream pricepb.PriceService_ListAssetsServer) error {
	data := &AssetItem{}
	cursor, err := assetsdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&pricepb.ListAssetsRes{
			Asset: &pricepb.Asset{
				Id:           data.Id.Hex(),
				Name:         data.Name,
				Exchange:     data.Exchange,
				AssetClass:   data.AssetClass,
				Symbol:       data.Symbol,
				Status:       data.Status,
				Tradable:     data.Tradable,
				Marginable:   data.Marginable,
				Shortable:    data.Shortable,
				EasyToBorrow: data.EasyToBorrow,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *PriceServiceServer) ListAssetBySymbol(req *pricepb.ListAssetBySymbolReq, stream pricepb.PriceService_ListAssetBySymbolServer) error {
	symbolQuery := req.GetSymbol()
	if len(symbolQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find Symbol in Req"))
	}

	data := &AssetItem{}
	cursor, err := assetsdb.Find(context.Background(), bson.M{"symbol": symbolQuery})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&pricepb.ListAssetBySymbolRes{
			Asset: &pricepb.Asset{
				Id:           data.Id.Hex(),
				Name:         data.Name,
				Exchange:     data.Exchange,
				AssetClass:   data.AssetClass,
				Symbol:       data.Symbol,
				Status:       data.Status,
				Tradable:     data.Tradable,
				Marginable:   data.Marginable,
				Shortable:    data.Shortable,
				EasyToBorrow: data.EasyToBorrow,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *PriceServiceServer) ListAssetByName(req *pricepb.ListAssetByNameReq, stream pricepb.PriceService_ListAssetByNameServer) error {
	nameQuery := req.GetName()
	if len(nameQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find Name in Req"))
	}

	data := &AssetItem{}
	cursor, err := assetsdb.Find(context.Background(), bson.M{"name": nameQuery})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&pricepb.ListAssetByNameRes{
			Asset: &pricepb.Asset{
				Id:           data.Id.Hex(),
				Name:         data.Name,
				Exchange:     data.Exchange,
				AssetClass:   data.AssetClass,
				Symbol:       data.Symbol,
				Status:       data.Status,
				Tradable:     data.Tradable,
				Marginable:   data.Marginable,
				Shortable:    data.Shortable,
				EasyToBorrow: data.EasyToBorrow,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
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

type AssetItem struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Exchange     string             `bson:"exchange"`
	AssetClass   string             `bson:"asset_class"`
	Symbol       string             `bson:"symbol"`
	Status       string             `bson:"status"`
	Tradable     bool               `bson:"tradable"`
	Marginable   bool               `bson:"marginable"`
	Shortable    bool               `bson:"shortable"`
	EasyToBorrow bool               `bson:"easy_to_borrow"`
}

var db *mongo.Client
var pricedb *mongo.Collection
var exchangedb *mongo.Collection
var assetsdb *mongo.Collection
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
	assetsdb = mongoDB.Collection("assets")

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
