package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	ceprulepb "github.com/vikjdk7/Algotrading-Golang/cep-rule-service/proto"
	"google.golang.org/grpc"
)

func (s *CepRuleServiceServer) CalculateStrategyFundAllocation(ctx context.Context, req *ceprulepb.CalculateStrategyFundAllocationReq) (*ceprulepb.CalculateStrategyFundAllocationRes, error) {

	baseOrderSize := req.GetBaseOrderSize()
	safetyOrderSize := req.GetSafetyOrderSize()
	safetyOrderVolumeScale := req.GetSafetyOrderVolumeScale()
	//safetyOrderStepScale := req.GetSafetyOrderStepScale()
	totalNoDeals := req.GetTotalNoDeals()
	maxActiveSafetyTradeCount := req.GetMaxActiveSafetyTradeCount()

	totalFundForAllocation := safetyOrderSize
	fundForSafetyOrder := safetyOrderSize

	for i := 0; i < int(maxActiveSafetyTradeCount-1.0); i++ {
		fundForSafetyOrder *= safetyOrderVolumeScale
		totalFundForAllocation += fundForSafetyOrder
	}
	totalFundForAllocation = (baseOrderSize + totalFundForAllocation + 1.0) * float64(totalNoDeals)

	return &ceprulepb.CalculateStrategyFundAllocationRes{
		TotalFundForAllocation: totalFundForAllocation,
	}, nil
}

type CepRuleServiceServer struct{}

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50056...")

	// Start our listener, 50056 is the default gRPC port
	listener, err := net.Listen("tcp", ":50056")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50056: %v", err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)

	// Create ExchangeService type
	srv := &CepRuleServiceServer{}
	// Register the service with the server
	ceprulepb.RegisterCepRuleServiceServer(s, srv)

	/*
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

	*/
	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50056")

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
	//fmt.Println("Closing MongoDB connection")
	//db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
