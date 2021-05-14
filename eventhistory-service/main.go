package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	eventhistorypb "github.com/vikjdk7/Algotrading-Golang/eventhistory-service/proto"
)

func (s *EventHistoryServiceServer) ListEventHistoryExchange(req *eventhistorypb.ListEventHistoryExchangeReq, stream eventhistorypb.EventHistoryService_ListEventHistoryExchangeServer) error {

	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}

	data := &EventHistoryExchangeItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := eventhistory_exchangedb.Find(context.Background(), bson.M{"user_id": userIdQuery})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send exchange over stream
		stream.Send(&eventhistorypb.ListEventHistoryExchangeRes{
			EventHistoryExchange: &eventhistorypb.EventHistoryExchange{
				OperationType: data.OperationType,
				Timestamp:     data.Timestamp,
				Db:            data.Db,
				Collection:    data.Collection,
				Name:          data.Name,
				UserId:        data.UserId,
				ExchangeId:    data.ExchangeId,
				OldValue: &eventhistorypb.Exchange{
					SelectedExchange: data.OldValue.SelectedExchange,
					ExchangeName:     data.OldValue.ExchangeName,
					ExchangeType:     data.OldValue.ExchangeType,
					ApiKey:           data.OldValue.ApiKey,
					ApiSecret:        data.OldValue.ApiSecret,
				},
				NewValue: &eventhistorypb.Exchange{
					SelectedExchange: data.NewValue.SelectedExchange,
					ExchangeName:     data.NewValue.ExchangeName,
					ExchangeType:     data.NewValue.ExchangeType,
					ApiKey:           data.NewValue.ApiKey,
					ApiSecret:        data.NewValue.ApiSecret,
				},
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *EventHistoryServiceServer) ListEventHistoryStrategy(req *eventhistorypb.ListEventHistoryStrategyReq, stream eventhistorypb.EventHistoryService_ListEventHistoryStrategyServer) error {

	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}

	data := &EventHistoryStrategyItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := eventhistory_strategydb.Find(context.Background(), bson.M{"user_id": userIdQuery})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send exchange over stream
		stream.Send(&eventhistorypb.ListEventHistoryStrategyRes{
			EventHistoryStrategy: &eventhistorypb.EventHistoryStrategy{
				OperationType: data.OperationType,
				Timestamp:     data.Timestamp,
				Db:            data.Db,
				Collection:    data.Collection,
				Name:          data.Name,
				UserId:        data.UserId,
				StrategyId:    data.StrategyId,
				OldValue: &eventhistorypb.Strategy{
					StrategyName:              data.OldValue.StrategyName,
					SelectedExchange:          data.OldValue.SelectedExchange,
					StrategyType:              data.OldValue.StrategyType,
					StartOrderType:            data.OldValue.StartOrderType,
					DealStartCondition:        data.OldValue.DealStartCondition,
					BaseOrderSize:             data.OldValue.BaseOrderSize,
					SafetyOrderSize:           data.OldValue.SafetyOrderSize,
					MaxSafetyTradeCount:       data.OldValue.MaxSafetyTradeCount,
					MaxActiveSafetyTradeCount: data.OldValue.MaxActiveSafetyTradeCount,
					PriceDevation:             data.OldValue.PriceDevation,
					SafetyOrderVolumeScale:    data.OldValue.SafetyOrderVolumeScale,
					SafetyOrderStepScale:      data.OldValue.SafetyOrderStepScale,
					TakeProfit:                data.OldValue.TakeProfit,
					TargetProfit:              data.OldValue.TargetProfit,
					AllocateFundsToStrategy:   data.OldValue.AllocateFundsToStrategy,
					Version:                   data.OldValue.Version,
					Status:                    data.OldValue.Status,
					Stock:                     data.OldValue.Stock,
				},
				NewValue: &eventhistorypb.Strategy{
					StrategyName:              data.NewValue.StrategyName,
					SelectedExchange:          data.NewValue.SelectedExchange,
					StrategyType:              data.NewValue.StrategyType,
					StartOrderType:            data.NewValue.StartOrderType,
					DealStartCondition:        data.NewValue.DealStartCondition,
					BaseOrderSize:             data.NewValue.BaseOrderSize,
					SafetyOrderSize:           data.NewValue.SafetyOrderSize,
					MaxSafetyTradeCount:       data.NewValue.MaxSafetyTradeCount,
					MaxActiveSafetyTradeCount: data.NewValue.MaxActiveSafetyTradeCount,
					PriceDevation:             data.NewValue.PriceDevation,
					SafetyOrderVolumeScale:    data.NewValue.SafetyOrderVolumeScale,
					SafetyOrderStepScale:      data.NewValue.SafetyOrderStepScale,
					TakeProfit:                data.NewValue.TakeProfit,
					TargetProfit:              data.NewValue.TargetProfit,
					AllocateFundsToStrategy:   data.NewValue.AllocateFundsToStrategy,
					Version:                   data.NewValue.Version,
					Status:                    data.NewValue.Status,
					Stock:                     data.NewValue.Stock,
				},
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type EventHistoryServiceServer struct{}

type EventHistoryExchangeItem struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	OperationType string             `bson:"operation_type"`
	Timestamp     string             `bson:"timestamp"`
	Db            string             `bson:"db"`
	Collection    string             `bson:"collection"`
	Name          string             `bson:"name"`
	UserId        string             `bson:"user_id"`
	ExchangeId    string             `bson:"exchange_id"`
	OldValue      ExchangeItem       `bson:"old_value"`
	NewValue      ExchangeItem       `bson:"new_value"`
}

type ExchangeItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SelectedExchange string             `bson:"selected_exchange"`
	ExchangeName     string             `bson:"exchange_name"`
	ExchangeType     string             `bson:"exchange_type"`
	UserId           string             `bson:"user_id"`
	ApiKey           string             `bson:"api_key"`
	ApiSecret        string             `bson:"api_secret"`
}

type EventHistoryStrategyItem struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	OperationType string             `bson:"operation_type"`
	Timestamp     string             `bson:"timestamp"`
	Db            string             `bson:"db"`
	Collection    string             `bson:"collection"`
	Name          string             `bson:"name"`
	UserId        string             `bson:"user_id"`
	StrategyId    string             `bson:"strategy_id"`
	OldValue      StrategyItem       `bson:"old_value"`
	NewValue      StrategyItem       `bson:"new_value"`
}

type StrategyItem struct {
	Id                        primitive.ObjectID      `bson:"_id,omitempty"`
	StrategyName              string                  `bson:"strategy_name"`
	SelectedExchange          string                  `bson:"selected_exchange"`
	StrategyType              string                  `bson:"strategy_type"`
	StartOrderType            string                  `bson:"start_order_type"`
	DealStartCondition        string                  `bson:"deal_start_condition"`
	BaseOrderSize             float64                 `bson:"base_order_size"`
	SafetyOrderSize           float64                 `bson:"safety_order_size"`
	MaxSafetyTradeCount       string                  `bson:"max_safety_trade_count"`
	MaxActiveSafetyTradeCount string                  `bson:"max_active_safety_trade_count"`
	PriceDevation             string                  `bson:"price_devation"`
	SafetyOrderVolumeScale    string                  `bson:"safety_order_volume_scale"`
	SafetyOrderStepScale      string                  `bson:"safety_order_step_scale"`
	TakeProfit                string                  `bson:"take_profit"`
	TargetProfit              string                  `bson:"target_profit"`
	AllocateFundsToStrategy   string                  `bson:"allocate_funds_to_strategy"`
	UserId                    string                  `bson:"user_id"`
	Version                   int64                   `bson:"version"`
	Status                    string                  `bson:"status"`
	Stock                     []*eventhistorypb.Stock `bson:"stock"`
}

type StockItem struct {
	StockName string `bson:"stock_name"`
}

var db *mongo.Client
var eventhistory_exchangedb *mongo.Collection
var eventhistory_strategydb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50055...")

	// Start our listener, 50055 is the default gRPC port
	listener, err := net.Listen("tcp", ":50055")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50055: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)

	// Create ExchangeService type
	srv := &EventHistoryServiceServer{}
	// Register the service with the server
	eventhistorypb.RegisterEventHistoryServiceServer(s, srv)

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
	mongoDb := db.Database("hedgina_algobot")
	eventhistory_exchangedb = mongoDb.Collection("eventhistory_exchange")
	eventhistory_strategydb = mongoDb.Collection("eventhistory_strategy")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50055")

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
