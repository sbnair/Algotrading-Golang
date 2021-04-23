package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	strategypb "github.com/vikjdk7/Algotrading-Golang/strategy-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *StrategyServiceServer) CreateStrategy(ctx context.Context, req *strategypb.CreateStrategyReq) (*strategypb.CreateStrategyRes, error) {
	strategy := req.GetStrategy()
	//fmt.Println(strategy)
	strategy.StrategyType = "Long"
	strategy.StartOrderType = "Limit"
	strategy.DealStartCondition = "Open new trade asap"
	strategy.Status = "Stopped"
	strategy.Version = 1

	data := StrategyItem{
		// Id:       primitive.NilObjectID,
		StrategyName:            strategy.GetStrategyName(),
		SelectedExchange:        strategy.GetSelectedExchange(),
		StrategyType:            strategy.GetStrategyType(),
		StartOrderType:          strategy.GetStartOrderType(),
		DealStartCondition:      strategy.GetDealStartCondition(),
		BaseOrderSize:           strategy.GetBaseOrderSize(),
		SafetyOrderSize:         strategy.GetSafetyOrderSize(),
		MaxSafetyTradeAcc:       strategy.GetMaxSafetyTradeAcc(),
		PriceDevation:           strategy.GetPriceDevation(),
		SafetyOrderVolumeScale:  strategy.GetSafetyOrderVolumeScale(),
		SafetyOrderStepScale:    strategy.GetSafetyOrderStepScale(),
		TakeProfit:              strategy.GetTakeProfit(),
		TargetProfit:            strategy.GetTargetProfit(),
		AllocateFundsToStrategy: strategy.GetAllocateFundsToStrategy(),
		UserId:                  strategy.GetUserId(),
		Version:                 strategy.GetVersion(),
		Status:                  strategy.GetStatus(),
		Stock:                   strategy.GetStock(),
	}

	fmt.Println(data)
	// Insert the data into the database
	// *InsertOneResult contains the oid
	result, err := strategydb.InsertOne(mongoCtx, data)
	// check error
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to strategy
	oid := result.InsertedID.(primitive.ObjectID)
	strategy.Id = oid.Hex()
	//strategy.StrategyType = strategypb.Strategy_Type_name[int32(data.StrategyType)]
	//fmt.Println(strategypb.Strategy_Type_name[int32(data.StrategyType)])
	// return the strategy in a CreateStretegyRes type
	createStrategyResponse := &strategypb.CreateStrategyRes{
		Strategy: strategy,
	}
	fmt.Println(createStrategyResponse)
	return createStrategyResponse, nil
}

func (s *StrategyServiceServer) ListStrategies(req *strategypb.ListStrategyReq, stream strategypb.StrategyService_ListStrategiesServer) error {
	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}
	// Initiate a StrategyItem type to write decoded data to
	data := &StrategyItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := strategydb.Find(context.Background(), bson.M{"user_id": userIdQuery})
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
		stream.Send(&strategypb.ListStrategyRes{
			Strategy: &strategypb.Strategy{
				Id:                      data.Id.Hex(),
				StrategyName:            data.StrategyName,
				SelectedExchange:        data.SelectedExchange,
				StrategyType:            data.StrategyType,
				StartOrderType:          data.StartOrderType,
				DealStartCondition:      data.DealStartCondition,
				BaseOrderSize:           data.BaseOrderSize,
				SafetyOrderSize:         data.SafetyOrderSize,
				MaxSafetyTradeAcc:       data.MaxSafetyTradeAcc,
				PriceDevation:           data.PriceDevation,
				SafetyOrderVolumeScale:  data.SafetyOrderVolumeScale,
				SafetyOrderStepScale:    data.SafetyOrderStepScale,
				TakeProfit:              data.TakeProfit,
				TargetProfit:            data.TargetProfit,
				AllocateFundsToStrategy: data.AllocateFundsToStrategy,
				UserId:                  data.UserId,
				Version:                 data.Version,
				Status:                  data.Status,
				Stock:                   data.Stock,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

/* ---------------------------------------

func (s *StrategyServiceServer) ReadStrategy(ctx context.Context, req *strategypb.ReadStrategyReq) (*strategypb.ReadStrategyRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := strategydb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	data := StrategyItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadStrategyRes type
	response := &strategypb.ReadStrategyRes{
		Strategy: &strategypb.Strategy{
			Id:                            data.Id.Hex(),
			BotName:                       data.BotName,
			SelectedStrategy:              data.SelectedStrategy,
			BotType:                       data.BotType,
			Pairs:                         data.Pairs,
			StrategyType:                  data.StrategyType,
			ProfitCurrency:                data.ProfitCurrency,
			BaseOrderSize:                 data.BaseOrderSize,
			SafetyOrderSize:               data.SafetyOrderSize,
			OrderType:                     data.OrderType,
			TargetProfit:                  data.TargetProfit,
			ProfitType:                    data.ProfitType,
			TrailingDeviation:             data.TrailingDeviation,
			StopLoss:                      data.StopLoss,
			StopLossAction:                data.StopLossAction,
			MaxSafetyTradeAcc:             data.MaxSafetyTradeAcc,
			MaxActiveSafetyTradeAcc:       data.MaxActiveSafetyTradeAcc,
			PriceDevation:                 data.PriceDevation,
			SafetyOrderVolumeScale:        data.SafetyOrderVolumeScale,
			SafetyOrderStepScale:          data.SafetyOrderStepScale,
			Active:                        data.Active,
			NoDealIfDailyVolumeIsLessThan: data.NoDealIfDailyVolumeIsLessThan,
			MinPriceToOpenDeal:            data.MinPriceToOpenDeal,
			MaxPriceToOpenDeal:            data.MaxPriceToOpenDeal,
			CooldownBewtweenDeals:         data.CooldownBewtweenDeals,
			OpenDealStop:                  data.OpenDealStop,
			UserId:                        data.UserId,
		},
	}
	return response, nil
}

func (s *StrategyServiceServer) DeleteStrategy(ctx context.Context, req *strategypb.DeleteStrategyReq) (*strategypb.DeleteStrategyRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = strategydb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete strategy with id %s: %v", req.GetId(), err))
	}
	return &strategypb.DeleteStrategyRes{
		Success: true,
	}, nil
}
func (s *StrategyServiceServer) UpdateStrategy(ctx context.Context, req *strategypb.UpdateStrategyReq) (*strategypb.UpdateStrategyRes, error) {
	// Get the blog data from the request
	strategy := req.GetStrategy()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(strategy.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied strategy id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"bot_name":                             strategy.GetBotName(),
		"selected_strategy":                    strategy.GetSelectedStrategy(),
		"bot_type":                             strategy.GetBotType(),
		"pairs":                                strategy.GetPairs(),
		"strategy_type":                        strategy.GetStrategyType(),
		"profit_currency":                      strategy.GetProfitCurrency(),
		"base_order_size":                      strategy.GetBaseOrderSize(),
		"safety_order_size":                    strategy.GetSafetyOrderSize(),
		"order_type":                           strategy.GetOrderType(),
		"target_profit":                        strategy.GetTargetProfit(),
		"profit_type":                          strategy.GetProfitType(),
		"trailing_deviation":                   strategy.GetTrailingDeviation(),
		"stop_loss":                            strategy.GetStopLoss(),
		"stop_loss_action":                     strategy.GetStopLossAction(),
		"max_safety_trade_acc":                 strategy.GetMaxSafetyTradeAcc(),
		"max_active_safety_trade_acc":          strategy.GetMaxActiveSafetyTradeAcc(),
		"price_devation":                       strategy.GetPriceDevation(),
		"safety_order_volume_scale":            strategy.GetSafetyOrderVolumeScale(),
		"safety_order_step_scale":              strategy.GetSafetyOrderStepScale(),
		"active":                               strategy.GetActive(),
		"no_deal_if_daily_volume_is_less_than": strategy.GetNoDealIfDailyVolumeIsLessThan(),
		"min_price_to_open_deal":               strategy.GetMinPriceToOpenDeal(),
		"max_price_to_open_deal":               strategy.GetMaxPriceToOpenDeal(),
		"cooldown_bewtween_deals":              strategy.GetCooldownBewtweenDeals(),
		"open_deal_stop":                       strategy.GetOpenDealStop(),
		"user_id":                              strategy.GetUserId(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := strategydb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := StrategyItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find strategy with supplied ID: %v", err),
		)
	}
	return &strategypb.UpdateStrategyRes{
		Strategy: &strategypb.Strategy{
			Id:                            decoded.Id.Hex(),
			BotName:                       decoded.BotName,
			SelectedStrategy:              decoded.SelectedStrategy,
			BotType:                       decoded.BotType,
			Pairs:                         decoded.Pairs,
			StrategyType:                  decoded.StrategyType,
			ProfitCurrency:                decoded.ProfitCurrency,
			BaseOrderSize:                 decoded.BaseOrderSize,
			SafetyOrderSize:               decoded.SafetyOrderSize,
			OrderType:                     decoded.OrderType,
			TargetProfit:                  decoded.TargetProfit,
			ProfitType:                    decoded.ProfitType,
			TrailingDeviation:             decoded.TrailingDeviation,
			StopLoss:                      decoded.StopLoss,
			StopLossAction:                decoded.StopLossAction,
			MaxSafetyTradeAcc:             decoded.MaxSafetyTradeAcc,
			MaxActiveSafetyTradeAcc:       decoded.MaxActiveSafetyTradeAcc,
			PriceDevation:                 decoded.PriceDevation,
			SafetyOrderVolumeScale:        decoded.SafetyOrderVolumeScale,
			SafetyOrderStepScale:          decoded.SafetyOrderStepScale,
			Active:                        decoded.Active,
			NoDealIfDailyVolumeIsLessThan: decoded.NoDealIfDailyVolumeIsLessThan,
			MinPriceToOpenDeal:            decoded.MinPriceToOpenDeal,
			MaxPriceToOpenDeal:            decoded.MaxPriceToOpenDeal,
			CooldownBewtweenDeals:         decoded.CooldownBewtweenDeals,
			OpenDealStop:                  decoded.OpenDealStop,
			UserId:                        decoded.UserId,
		},
	}, nil
}


/* ---------------------------------------*/

type StrategyServiceServer struct{}
type StrategyItem struct {
	Id                      primitive.ObjectID  `bson:"_id,omitempty"`
	StrategyName            string              `bson:"strategy_name"`
	SelectedExchange        string              `bson:"selected_exchange"`
	StrategyType            string              `bson:"strategy_type"`
	StartOrderType          string              `bson:"start_order_type"`
	DealStartCondition      string              `bson:"deal_start_condition"`
	BaseOrderSize           float64             `bson:"base_order_size"`
	SafetyOrderSize         float64             `bson:"safety_order_size"`
	MaxSafetyTradeAcc       string              `bson:"max_safety_trade_acc"`
	PriceDevation           string              `bson:"price_devation"`
	SafetyOrderVolumeScale  string              `bson:"safety_order_volume_scale"`
	SafetyOrderStepScale    string              `bson:"safety_order_step_scale"`
	TakeProfit              string              `bson:"take_profit"`
	TargetProfit            string              `bson:"target_profit"`
	AllocateFundsToStrategy string              `bson:"allocate_funds_to_strategy"`
	UserId                  string              `bson:"user_id"`
	Version                 int64               `bson:"version"`
	Status                  string              `bson:"status"`
	Stock                   []*strategypb.Stock `bson:"stock"`
}

type Stock struct {
	StockName string `bson:"stock_name"`
}

var db *mongo.Client
var strategydb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50051...")

	// Start our listener, 50051 is the default gRPC port
	listener, err := net.Listen("tcp", ":50051")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)

	// Create StrategyService type
	srv := &StrategyServiceServer{}
	// Register the service with the server
	strategypb.RegisterStrategyServiceServer(s, srv)

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
	strategydb = db.Database("hedgina_algobot").Collection("strategy")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

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
