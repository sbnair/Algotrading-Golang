package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	exchangepb "github.com/vikjdk7/Algotrading-Golang/exchange-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ExchangeServiceServer) CreateExchange(ctx context.Context, req *exchangepb.CreateExchangeReq) (*exchangepb.CreateExchangeRes, error) {
	// Get the protobuf exchange type from the protobuf request type
	// Essentially doing req.Exchange to access the struct with a nil check
	exchange := req.GetExchange()
	fmt.Println(exchange)
	// Now we have to convert this into a ExchangeItem type to convert into BSON
	data := ExchangeItem{
		// ID:       primitive.NilObjectID,
		SelectedExchange: exchange.GetSelectedExchange(),
		ExchangeName:     exchange.GetExchangeName(),
		ExchangeType:     exchange.GetExchangeType(),
		UserId:           exchange.GetUserId(),
		ApiKey:           exchange.GetApiKey(),
		ApiSecret:        exchange.GetApiSecret(),
	}

	exchangeAccount := &exchangepb.ExchangeAccountInfo{}

	//println(exchangeAccount)
	if exchange.SelectedExchange == "Alpaca" {
		os.Setenv(common.EnvApiKeyID, exchange.ApiKey)
		os.Setenv(common.EnvApiSecretKey, exchange.ApiSecret)
		if exchange.ExchangeType == "paper_trading" {
			alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
		} else if exchange.ExchangeType == "live_trading" {
			alpaca.SetBaseUrl("https://api.alpaca.markets")
		}

		alpacaClient := alpaca.NewClient(common.Credentials())
		acct, err := alpacaClient.GetAccount()
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
		println(acct)
		exchangeAccount.Id = acct.ID
		exchangeAccount.AccountNumber = acct.AccountNumber
		exchangeAccount.CreatedAt = acct.CreatedAt.Format(time.RFC3339)
		exchangeAccount.Status = acct.Status
		exchangeAccount.Currency = acct.Currency
		exchangeAccount.Cash, err = strconv.ParseFloat(acct.Cash.String(), 64)
		exchangeAccount.CashWithdrawable, err = strconv.ParseFloat(acct.CashWithdrawable.String(), 64)
		exchangeAccount.TradingBlocked = acct.TradingBlocked
		exchangeAccount.TransfersBlocked = acct.TransfersBlocked
		exchangeAccount.AccountBlocked = acct.AccountBlocked
		exchangeAccount.BuyingPower, err = strconv.ParseFloat(acct.BuyingPower.String(), 64)
		exchangeAccount.PatternDayTrader = acct.PatternDayTrader
		exchangeAccount.DaytradeCount = acct.DaytradeCount
		exchangeAccount.DaytradingBuyingPower, err = strconv.ParseFloat(acct.DaytradingBuyingPower.String(), 64)
		exchangeAccount.RegtBuyingPower, err = strconv.ParseFloat(acct.RegTBuyingPower.String(), 64)
		exchangeAccount.Equity, err = strconv.ParseFloat(acct.Equity.String(), 64)
		exchangeAccount.LastEquity, err = strconv.ParseFloat(acct.LastEquity.String(), 64)
		exchangeAccount.InitialMargin, err = strconv.ParseFloat(acct.InitialMargin.String(), 64)
		exchangeAccount.LongMarketValue, err = strconv.ParseFloat(acct.LongMarketValue.String(), 64)
		exchangeAccount.ShortMarketValue, err = strconv.ParseFloat(acct.ShortMarketValue.String(), 64)

		//jsonpb.Unmarshal(acct, &exchangeAccount)
	}

	// Insert the data into the database
	// *InsertOneResult contains the oid
	result, err := exchangedb.InsertOne(mongoCtx, data)
	// check error
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to blog
	oid := result.InsertedID.(primitive.ObjectID)
	exchange.Id = oid.Hex()

	data.ID = oid
	eventData := EventHistoryItem{
		OperationType: "insert",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "exchange",
		Name:          data.ExchangeName,
		UserId:        data.UserId,
		ExchangeId:    data.ID.Hex(),
		NewValue:      data,
	}
	_, errEventHistory := eventhistorydb.InsertOne(mongoCtx, eventData)
	if errEventHistory != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", errEventHistory),
		)
	}
	// return the blog in a CreateBlogRes type
	createExchangeResponse := &exchangepb.CreateExchangeRes{
		Exchange:            exchange,
		ExchangeAccountInfo: exchangeAccount,
	}
	fmt.Println(createExchangeResponse)
	return createExchangeResponse, nil
}

func (s *ExchangeServiceServer) ReadExchange(ctx context.Context, req *exchangepb.ReadExchangeReq) (*exchangepb.ReadExchangeRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := exchangedb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	data := ExchangeItem{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Exchange with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadExchangeRes type
	response := &exchangepb.ReadExchangeRes{
		Exchange: &exchangepb.Exchange{
			Id:               oid.Hex(),
			SelectedExchange: data.SelectedExchange,
			ExchangeName:     data.ExchangeName,
			ExchangeType:     data.ExchangeType,
			UserId:           data.UserId,
			ApiKey:           data.ApiKey,
			ApiSecret:        data.ApiSecret,
		},
	}
	return response, nil
}

func (s *ExchangeServiceServer) DeleteExchange(ctx context.Context, req *exchangepb.DeleteExchangeReq) (*exchangepb.DeleteExchangeRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	resultReadExchange := exchangedb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataResultReadExchange := ExchangeItem{}
	// decode and write to data
	if err := resultReadExchange.Decode(&dataResultReadExchange); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Exchange with Object Id %s: %v", req.GetId(), err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = exchangedb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete exchange with id %s: %v", req.GetId(), err))
	}

	eventData := EventHistoryItem{
		OperationType: "delete",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "exchange",
		Name:          dataResultReadExchange.ExchangeName,
		UserId:        dataResultReadExchange.UserId,
		ExchangeId:    oid.Hex(),
		OldValue:      dataResultReadExchange,
	}
	_, errEventHistory := eventhistorydb.InsertOne(mongoCtx, eventData)
	if errEventHistory != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", errEventHistory),
		)
	}

	return &exchangepb.DeleteExchangeRes{
		Success: true,
	}, nil
}

func (s *ExchangeServiceServer) UpdateExchange(ctx context.Context, req *exchangepb.UpdateExchangeReq) (*exchangepb.UpdateExchangeRes, error) {
	// Get the blog data from the request
	exchange := req.GetExchange()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(exchange.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied exchange id to a MongoDB ObjectId: %v", err),
		)
	}

	resultReadExchange := exchangedb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataResultReadExchange := ExchangeItem{}
	// decode and write to data
	if err := resultReadExchange.Decode(&dataResultReadExchange); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Exchange with Object Id %s: %v", exchange.GetId(), err))
	}

	oldValues := ExchangeItem{}
	newValues := ExchangeItem{}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{}

	if exchange.GetSelectedExchange() != "" {
		update["selected_exchange"] = exchange.GetSelectedExchange()
		oldValues.SelectedExchange = dataResultReadExchange.SelectedExchange
		newValues.SelectedExchange = exchange.GetApiSecret()
	}
	if exchange.GetExchangeName() != "" {
		update["exchange_name"] = exchange.GetExchangeName()
		oldValues.ExchangeName = dataResultReadExchange.ExchangeName
		newValues.ExchangeName = exchange.GetExchangeName()
	}
	if exchange.GetExchangeType() != "" {
		update["exchange_type"] = exchange.GetExchangeType()
		oldValues.ExchangeType = dataResultReadExchange.ExchangeType
		newValues.ExchangeType = exchange.GetExchangeType()
	}
	if exchange.GetUserId() != "" {
		update["user_id"] = exchange.GetUserId()
		oldValues.UserId = dataResultReadExchange.UserId
		newValues.UserId = exchange.GetUserId()
	}
	if exchange.GetApiKey() != "" {
		update["api_key"] = exchange.GetApiKey()
		oldValues.ApiKey = dataResultReadExchange.ApiKey
		newValues.ApiKey = exchange.GetApiKey()
	}
	if exchange.GetApiSecret() != "" {
		update["api_secret"] = exchange.GetApiSecret()
		oldValues.ApiSecret = dataResultReadExchange.ApiSecret
		newValues.ApiSecret = exchange.GetApiSecret()
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := exchangedb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := ExchangeItem{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find exchange with supplied ID: %v", err),
		)
	}

	eventData := EventHistoryItem{
		OperationType: "update",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "exchange",
		Name:          dataResultReadExchange.ExchangeName,
		UserId:        dataResultReadExchange.UserId,
		ExchangeId:    oid.Hex(),
		OldValue:      oldValues,
		NewValue:      newValues,
	}
	_, errEventHistory := eventhistorydb.InsertOne(mongoCtx, eventData)
	if errEventHistory != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", errEventHistory),
		)
	}

	return &exchangepb.UpdateExchangeRes{
		Exchange: &exchangepb.Exchange{
			Id:               decoded.ID.Hex(),
			SelectedExchange: decoded.SelectedExchange,
			ExchangeName:     decoded.ExchangeName,
			ExchangeType:     decoded.ExchangeType,
			UserId:           decoded.UserId,
			ApiKey:           decoded.ApiKey,
			ApiSecret:        decoded.ApiSecret,
		},
	}, nil
}

func (s *ExchangeServiceServer) ListExchanges(req *exchangepb.ListExchangeReq, stream exchangepb.ExchangeService_ListExchangesServer) error {
	// Initiate a BlogItem type to write decoded data to
	data := &ExchangeItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := exchangedb.Find(context.Background(), bson.M{})
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
		stream.Send(&exchangepb.ListExchangeRes{
			Exchange: &exchangepb.Exchange{
				Id:               data.ID.Hex(),
				SelectedExchange: data.SelectedExchange,
				ExchangeName:     data.ExchangeName,
				ExchangeType:     data.ExchangeType,
				UserId:           data.UserId,
				ApiKey:           data.ApiKey,
				ApiSecret:        data.ApiSecret,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *ExchangeServiceServer) ListExchangesByUser(req *exchangepb.ListExchangesByUserReq, stream exchangepb.ExchangeService_ListExchangesByUserServer) error {
	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}
	data := &ExchangeItem{}

	cursor, err := exchangedb.Find(context.Background(), bson.M{"user_id": userIdQuery})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send exchange over stream
		stream.Send(&exchangepb.ListExchangesByUserRes{
			Exchange: &exchangepb.Exchange{
				Id:               data.ID.Hex(),
				SelectedExchange: data.SelectedExchange,
				ExchangeName:     data.ExchangeName,
				ExchangeType:     data.ExchangeType,
				UserId:           data.UserId,
				ApiKey:           data.ApiKey,
				ApiSecret:        data.ApiSecret,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type ExchangeServiceServer struct{}

type ExchangeItem struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SelectedExchange string             `bson:"selected_exchange"`
	ExchangeName     string             `bson:"exchange_name"`
	ExchangeType     string             `bson:"exchange_type"`
	UserId           string             `bson:"user_id"`
	ApiKey           string             `bson:"api_key"`
	ApiSecret        string             `bson:"api_secret"`
}

type EventHistoryItem struct {
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

var db *mongo.Client
var exchangedb *mongo.Collection
var eventhistorydb *mongo.Collection
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

	// Create ExchangeService type
	srv := &ExchangeServiceServer{}
	// Register the service with the server
	exchangepb.RegisterExchangeServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	//Uncomment to run locally
	//os.Setenv("MONGODB_URL", "mongodb://127.0.0.1:27017")

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
	exchangedb = mongoDb.Collection("exchange")
	eventhistorydb = mongoDb.Collection("eventhistory_exchange")

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
