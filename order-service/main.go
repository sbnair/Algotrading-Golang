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
	"github.com/shopspring/decimal"
	orderpb "github.com/vikjdk7/Algotrading-Golang/order-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderServiceServer) PlaceOrder(ctx context.Context, req *orderpb.PlaceOrderReq) (*orderpb.PlaceOrderRes, error) {
	fmt.Print("Request Data: ")
	fmt.Println(req)

	exchange_id := req.GetExchangeId()

	//fmt.Println(exchange_id)

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(exchange_id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied Exchange id to a MongoDB Object Id: %v", err),
		)
	}

	resultReadExchange := exchangedb.FindOne(mongoCtx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataRead := ExchangeItem{}
	// decode and write to dataRead
	if err := resultReadExchange.Decode(&dataRead); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Exchange with Object Id %s: %v", oid, err))
	}

	fmt.Print("Exchange Data Read: ")
	fmt.Println(dataRead)

	placeOrderResponse := &orderpb.PlaceOrderRes{}

	if dataRead.SelectedExchange == "Alpaca" {
		os.Setenv(common.EnvApiKeyID, dataRead.ApiKey)
		os.Setenv(common.EnvApiSecretKey, dataRead.ApiSecret)
		if dataRead.ExchangeType == "paper_trading" {
			alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
		} else if dataRead.ExchangeType == "live_trading" {
			alpaca.SetBaseUrl("https://api.alpaca.markets")
		}

		alpacaClient := alpaca.NewClient(common.Credentials())

		placeOrderRequest := alpaca.PlaceOrderRequest{}

		reqSymbolOrderRequest := req.GetSymbol()
		if reqSymbolOrderRequest != "" {
			placeOrderRequest.AssetKey = &reqSymbolOrderRequest
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Symbol cannot be empty or null"))
		}

		qtyOrderRequest := decimal.NewFromFloat(req.GetQty())
		if req.GetQty() > 0.0 {
			placeOrderRequest.Qty = qtyOrderRequest
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Quantity should be greate than 0"))
		}

		sideOrderRequest := req.GetSide()
		if sideOrderRequest == orderpb.Side_buy {
			fmt.Println("side type buy")
			placeOrderRequest.Side = alpaca.Buy
		} else if sideOrderRequest == orderpb.Side_sell {
			placeOrderRequest.Side = alpaca.Sell
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid value for Side. Use either buy or sell"))
		}

		limitPriceOrderRequest := req.GetLimitPrice()
		typeOrderRequest := req.GetOrderType()
		if typeOrderRequest == orderpb.OrderType_market {
			placeOrderRequest.Type = alpaca.Market
		} else if typeOrderRequest == orderpb.OrderType_limit {
			if limitPriceOrderRequest <= 0.0 {
				return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid value for Limit Price. Limit Price should be > 0.0"))
			}
			placeOrderRequest.Type = alpaca.Limit
		} else if typeOrderRequest == orderpb.OrderType_stop {
			placeOrderRequest.Type = alpaca.Stop
		} else if typeOrderRequest == orderpb.OrderType_stop_limit {
			if limitPriceOrderRequest <= 0.0 {
				return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid value for Limit Price. Limit Price should be > 0.0"))
			}
			placeOrderRequest.Type = alpaca.StopLimit
		} else if typeOrderRequest == orderpb.OrderType_trailing_stop {
			placeOrderRequest.Type = alpaca.TrailingStop
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid value for Order Type. Use one of market,limit,stop,stop_limit,trailing_stop"))
		}

		timeInForceOrderRequest := req.GetTimeInForce()
		if timeInForceOrderRequest == orderpb.TimeInForce_day {
			placeOrderRequest.TimeInForce = alpaca.Day
		} else if timeInForceOrderRequest == orderpb.TimeInForce_gtc {
			placeOrderRequest.TimeInForce = alpaca.GTC
		} else if timeInForceOrderRequest == orderpb.TimeInForce_opg {
			placeOrderRequest.TimeInForce = alpaca.OPG
		} else if timeInForceOrderRequest == orderpb.TimeInForce_ioc {
			placeOrderRequest.TimeInForce = alpaca.IOC
		} else if timeInForceOrderRequest == orderpb.TimeInForce_fok {
			placeOrderRequest.TimeInForce = alpaca.FOK
		} else if timeInForceOrderRequest == orderpb.TimeInForce_gtx {
			placeOrderRequest.TimeInForce = alpaca.GTX
		} else if timeInForceOrderRequest == orderpb.TimeInForce_gtd {
			placeOrderRequest.TimeInForce = alpaca.GTD
		} else if timeInForceOrderRequest == orderpb.TimeInForce_cls {
			placeOrderRequest.TimeInForce = alpaca.CLS
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid value for Time in Force. Use one of day,gtc,opg,ioc,fok,gtx,gtd,cls"))
		}

		if limitPriceOrderRequest > 0.0 {
			if typeOrderRequest == orderpb.OrderType_limit || typeOrderRequest == orderpb.OrderType_stop_limit {
				reqLimitPrice := decimal.NewFromFloat(req.GetLimitPrice())
				placeOrderRequest.LimitPrice = &reqLimitPrice
			} else {
				return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Limit Price cannot be used with order type %v", typeOrderRequest))
			}
		}

		fmt.Print("PlaceOrderRequest Data: ")
		fmt.Println(placeOrderRequest)
		orderPlaced, err := alpacaClient.PlaceOrder(placeOrderRequest)
		fmt.Print("OrderPlaced Result: ")
		fmt.Println(orderPlaced)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not place the order %v", err))
		}

		qtyOrderPlacedRes, _ := orderPlaced.Qty.Float64()
		notionalOrderPlacedRes, _ := orderPlaced.Notional.Float64()
		filledQtyOrderPlacedRes, _ := orderPlaced.FilledQty.Float64()
		//limitPriceOrderPlacedRes, _ := orderPlaced.LimitPrice.Float64()
		//filledAvgPriceOrderPlacedRes, _ := orderPlaced.FilledAvgPrice.Float64()
		//stopPriceOrderPlacedRes, _ := orderPlaced.StopPrice.Float64()
		//trailPriceOrderPlacedRes, _ := orderPlaced.TrailPrice.Float64()
		//trailPercentOrderPlacedRes, _ := orderPlaced.TrailPercent.Float64()
		//hwmOrderPlacedRes, _ := orderPlaced.Hwm.Float64()

		insert := OrderItem{
			Id:            orderPlaced.ID,
			ClientOrderId: orderPlaced.ClientOrderID,
			CreatedAt:     orderPlaced.CreatedAt.String(),
			UpdatedAt:     orderPlaced.UpdatedAt.String(),
			SubmittedAt:   orderPlaced.SubmittedAt.String(),
			//FilledAt:       orderPlaced.FilledAt.String(),
			//ExpiredAt:      orderPlaced.ExpiredAt.String(),
			//CanceledAt:     orderPlaced.CanceledAt.String(),
			//FailedAt:       orderPlaced.FailedAt.String(),
			//ReplacedAt:     orderPlaced.ReplacedAt.String(),
			AssetId:     orderPlaced.AssetID,
			Symbol:      orderPlaced.Symbol,
			Exchange:    orderPlaced.Exchange,
			AssetClass:  orderPlaced.Class,
			Qty:         qtyOrderPlacedRes,
			Notional:    notionalOrderPlacedRes,
			FilledQty:   filledQtyOrderPlacedRes,
			OrderType:   typeOrderRequest.String(),
			Side:        sideOrderRequest.String(),
			TimeInForce: timeInForceOrderRequest.String(),
			LimitPrice:  limitPriceOrderRequest,
			//FilledAvgPrice: filledAvgPriceOrderPlacedRes,
			//StopPrice:      stopPriceOrderPlacedRes,
			//TrailPrice:     trailPriceOrderPlacedRes,
			//TrailPercent:   trailPercentOrderPlacedRes,
			//Hwm:            hwmOrderPlacedRes,
			Status:        orderPlaced.Status,
			ExtendedHours: orderPlaced.ExtendedHours,
			//Legs:           orderPlaced.Legs,
			UserId: dataRead.UserId,
		}

		fmt.Print("insert: ")
		fmt.Println(insert)

		_, err = orderdb.InsertOne(mongoCtx, insert)
		if err != nil {
			// return internal gRPC error to be handled later
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}

		order := &orderpb.Order{
			Id:            orderPlaced.ID,
			ClientOrderId: orderPlaced.ClientOrderID,
			CreatedAt:     orderPlaced.CreatedAt.String(),
			UpdatedAt:     orderPlaced.UpdatedAt.String(),
			SubmittedAt:   orderPlaced.SubmittedAt.String(),
			//FilledAt:       orderPlaced.FilledAt.String(),
			//ExpiredAt:      orderPlaced.ExpiredAt.String(),
			//CanceledAt:     orderPlaced.CanceledAt.String(),
			//FailedAt:       orderPlaced.FailedAt.String(),
			AssetId:     orderPlaced.AssetID,
			Symbol:      orderPlaced.Symbol,
			Exchange:    orderPlaced.Exchange,
			AssetClass:  orderPlaced.Class,
			Qty:         qtyOrderPlacedRes,
			Notional:    notionalOrderPlacedRes,
			FilledQty:   filledQtyOrderPlacedRes,
			OrderType:   typeOrderRequest.String(),
			Side:        sideOrderRequest.String(),
			TimeInForce: timeInForceOrderRequest.String(),
			LimitPrice:  limitPriceOrderRequest,
			//FilledAvgPrice: filledAvgPriceOrderPlacedRes,
			//StopPrice:      stopPriceOrderPlacedRes,
			//TrailPrice:     trailPriceOrderPlacedRes,
			//TrailPercent:   trailPercentOrderPlacedRes,
			//Hwm:            hwmOrderPlacedRes,
			Status:        orderPlaced.Status,
			ExtendedHours: orderPlaced.ExtendedHours,
			//Legs:           orderPlaced.Legs,
			UserId: dataRead.UserId,
		}

		placeOrderResponse.Order = order
	} else {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot use exchange other than Alpaca"))
	}
	fmt.Println(placeOrderResponse)
	return placeOrderResponse, nil
}

func (s *OrderServiceServer) ListOrders(req *orderpb.ListOrdersReq, stream orderpb.OrderService_ListOrdersServer) error {
	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}
	data := &OrderItem{}

	cursor, err := orderdb.Find(context.Background(), bson.M{"user_id": userIdQuery})
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
		stream.Send(&orderpb.ListOrdersRes{
			Order: &orderpb.Order{
				Id:            data.Id,
				ClientOrderId: data.ClientOrderId,
				CreatedAt:     data.CreatedAt,
				UpdatedAt:     data.UpdatedAt,
				SubmittedAt:   data.SubmittedAt,
				//FilledAt:       data.FilledAt.String(),
				//ExpiredAt:      data.ExpiredAt.String(),
				//CanceledAt:     data.CanceledAt.String(),
				//FailedAt:       data.FailedAt.String(),
				AssetId:     data.AssetId,
				Symbol:      data.Symbol,
				Exchange:    data.Exchange,
				AssetClass:  data.AssetClass,
				Qty:         data.Qty,
				Notional:    data.Notional,
				FilledQty:   data.FilledQty,
				OrderType:   data.OrderType,
				Side:        data.Side,
				TimeInForce: data.TimeInForce,
				LimitPrice:  data.LimitPrice,
				//FilledAvgPrice: filledAvgPricedataRes,
				//StopPrice:      stopPricedataRes,
				//TrailPrice:     trailPricedataRes,
				//TrailPercent:   trailPercentdataRes,
				//Hwm:            hwmdataRes,
				Status:        data.Status,
				ExtendedHours: data.ExtendedHours,
				//Legs:           data.Legs,
				UserId: data.UserId,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
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

type OrderItem struct {
	Id             string       `bson:"_id,omitempty"`
	ClientOrderId  string       `bson:"client_order_id"`
	CreatedAt      string       `bson:"created_at"`
	UpdatedAt      string       `bson:"updated_at"`
	SubmittedAt    string       `bson:"submitted_at"`
	FilledAt       string       `bson:"filled_at"`
	ExpiredAt      string       `bson:"expired_at"`
	CanceledAt     string       `bson:"canceled_at"`
	FailedAt       string       `bson:"failed_at"`
	ReplacedAt     string       `bson:"replaced_at"`
	Replaces       string       `bson:"replaces"`
	ReplacedBy     string       `bson:"replaced_by"`
	AssetId        string       `bson:"asset_id"`
	Symbol         string       `bson:"symbol"`
	Exchange       string       `bson:"exchange"`
	AssetClass     string       `bson:"asset_class"`
	Qty            float64      `bson:"qty"`
	Notional       float64      `bson:"notional"`
	FilledQty      float64      `bson:"filled_qty"`
	OrderType      string       `bson:"order_type"`
	Side           string       `bson:"side"`
	TimeInForce    string       `bson:"time_in_force"`
	LimitPrice     float64      `bson:"limit_price"`
	FilledAvgPrice float64      `bson:"filled_avg_price"`
	StopPrice      float64      `bson:"stop_price"`
	TrailPrice     float64      `bson:"trail_price"`
	TrailPercent   float64      `bson:"trail_percent"`
	Hwm            float64      `bson:"hwm"`
	Status         string       `bson:"status"`
	ExtendedHours  bool         `bson:"extended_hours"`
	Legs           []*OrderItem `bson:"legs"`
	UserId         string       `bson:"user_id"`
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
