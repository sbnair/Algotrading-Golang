package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

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
	strategy.Status = "stopped"
	strategy.Version = 1

	data := StrategyItem{
		// Id:       primitive.NilObjectID,
		StrategyName:              strategy.GetStrategyName(),
		SelectedExchange:          strategy.GetSelectedExchange(),
		StrategyType:              strategy.GetStrategyType(),
		StartOrderType:            strategy.GetStartOrderType(),
		DealStartCondition:        strategy.GetDealStartCondition(),
		BaseOrderSize:             strategy.GetBaseOrderSize(),
		SafetyOrderSize:           strategy.GetSafetyOrderSize(),
		MaxSafetyTradeCount:       strategy.GetMaxSafetyTradeCount(),
		MaxActiveSafetyTradeCount: strategy.GetMaxActiveSafetyTradeCount(),
		PriceDevation:             strategy.GetPriceDevation(),
		SafetyOrderVolumeScale:    strategy.GetSafetyOrderVolumeScale(),
		SafetyOrderStepScale:      strategy.GetSafetyOrderStepScale(),
		TakeProfit:                strategy.GetTakeProfit(),
		TargetProfit:              strategy.GetTargetProfit(),
		AllocateFundsToStrategy:   strategy.GetAllocateFundsToStrategy(),
		UserId:                    strategy.GetUserId(),
		Version:                   strategy.GetVersion(),
		Status:                    strategy.GetStatus(),
		Stock:                     strategy.GetStock(),
	}

	//fmt.Println(data)
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

	data.Id = oid
	eventData := EventHistoryItem{
		OperationType: "insert",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "strategy",
		Name:          data.StrategyName,
		UserId:        data.UserId,
		StrategyId:    data.Id.Hex(),
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
	//strategy.StrategyType = strategypb.Strategy_Type_name[int32(data.StrategyType)]
	//fmt.Println(strategypb.Strategy_Type_name[int32(data.StrategyType)])
	// return the strategy in a CreateStretegyRes type
	createStrategyResponse := &strategypb.CreateStrategyRes{
		Strategy: strategy,
	}
	//fmt.Println(createStrategyResponse)
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
				Id:                        data.Id.Hex(),
				StrategyName:              data.StrategyName,
				SelectedExchange:          data.SelectedExchange,
				StrategyType:              data.StrategyType,
				StartOrderType:            data.StartOrderType,
				DealStartCondition:        data.DealStartCondition,
				BaseOrderSize:             data.BaseOrderSize,
				SafetyOrderSize:           data.SafetyOrderSize,
				MaxSafetyTradeCount:       data.MaxSafetyTradeCount,
				MaxActiveSafetyTradeCount: data.MaxActiveSafetyTradeCount,
				PriceDevation:             data.PriceDevation,
				SafetyOrderVolumeScale:    data.SafetyOrderVolumeScale,
				SafetyOrderStepScale:      data.SafetyOrderStepScale,
				TakeProfit:                data.TakeProfit,
				TargetProfit:              data.TargetProfit,
				AllocateFundsToStrategy:   data.AllocateFundsToStrategy,
				UserId:                    data.UserId,
				Version:                   data.Version,
				Status:                    data.Status,
				Stock:                     data.Stock,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

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
			Id:                        data.Id.Hex(),
			StrategyName:              data.StrategyName,
			SelectedExchange:          data.SelectedExchange,
			StrategyType:              data.StrategyType,
			StartOrderType:            data.StartOrderType,
			DealStartCondition:        data.DealStartCondition,
			BaseOrderSize:             data.BaseOrderSize,
			SafetyOrderSize:           data.SafetyOrderSize,
			MaxSafetyTradeCount:       data.MaxSafetyTradeCount,
			MaxActiveSafetyTradeCount: data.MaxActiveSafetyTradeCount,
			PriceDevation:             data.PriceDevation,
			SafetyOrderVolumeScale:    data.SafetyOrderVolumeScale,
			SafetyOrderStepScale:      data.SafetyOrderStepScale,
			TakeProfit:                data.TakeProfit,
			TargetProfit:              data.TargetProfit,
			AllocateFundsToStrategy:   data.AllocateFundsToStrategy,
			UserId:                    data.UserId,
			Version:                   data.Version,
			Status:                    data.Status,
			Stock:                     data.Stock,
		},
	}
	return response, nil
}

func (s *StrategyServiceServer) DeleteStrategy(ctx context.Context, req *strategypb.DeleteStrategyReq) (*strategypb.DeleteStrategyRes, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	dealsCount, err := dealsdb.CountDocuments(ctx, bson.M{"strategy_id": oid.Hex(), "status": "running"})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Could not find deals with strategy id %s: %v", req.GetId(), err))
	}
	fmt.Print("Deals Count with status running: ")
	fmt.Println(dealsCount)
	if dealsCount < 1 {

		resultReadStrategy := strategydb.FindOne(ctx, bson.M{"_id": oid})
		// Create an empty ExchangeItem to write our decode result to
		dataResultReadStrategy := StrategyItem{}
		// decode and write to data
		if err := resultReadStrategy.Decode(&dataResultReadStrategy); err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", req.GetId(), err))
		}
		// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
		// So we return a boolean instead
		_, err = strategydb.DeleteOne(ctx, bson.M{"_id": oid})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete strategy with id %s: %v", req.GetId(), err))
		}

		_, err = strategy_revisionsdb.DeleteMany(ctx, bson.M{"strategy_id": req.GetId()})
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not delete from strategy_revisionsdb with id %s: %v", req.GetId(), err))
		}

		eventData := EventHistoryItem{
			OperationType: "delete",
			Timestamp:     time.Now().Format(time.RFC3339),
			Db:            "hedgina_algobot",
			Collection:    "strategy",
			Name:          dataResultReadStrategy.StrategyName,
			UserId:        dataResultReadStrategy.UserId,
			StrategyId:    oid.Hex(),
			OldValue:      dataResultReadStrategy,
		}
		_, errEventHistory := eventhistorydb.InsertOne(mongoCtx, eventData)
		if errEventHistory != nil {
			// return internal gRPC error to be handled later
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", errEventHistory),
			)
		}
		return &strategypb.DeleteStrategyRes{
			Success: true,
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot delete strategy with %d running deal(s)", dealsCount))
	}

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

	resultReadStrategy := strategydb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty ExchangeItem to write our decode result to
	dataRead := StrategyRevisionItem{}
	// decode and write to dataRead
	if err := resultReadStrategy.Decode(&dataRead); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", oid, err))
	}
	dataRead.StrategyId = oid.Hex()
	dataRead.Id = primitive.NewObjectID()
	resultInsert, err := strategy_revisionsdb.InsertOne(mongoCtx, dataRead)
	// check error
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	fmt.Sprintf("Inserted in strategy_revisionsdb: %v", resultInsert.InsertedID)

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{}
	oldValues := StrategyItem{}
	newValues := StrategyItem{}

	if strategy.GetStrategyName() != "" {
		update["strategy_name"] = strategy.GetStrategyName()
		oldValues.StrategyName = dataRead.StrategyName
		newValues.StrategyName = strategy.GetStrategyName()
	}
	if strategy.GetSelectedExchange() != "" {
		update["selected_exchange"] = strategy.GetSelectedExchange()
		oldValues.SelectedExchange = dataRead.SelectedExchange
		newValues.SelectedExchange = strategy.GetSelectedExchange()
	}
	if strategy.GetBaseOrderSize() != 0 {
		update["base_order_size"] = strategy.GetBaseOrderSize()
		oldValues.BaseOrderSize = dataRead.BaseOrderSize
		newValues.BaseOrderSize = strategy.GetBaseOrderSize()
	}
	if strategy.GetSafetyOrderSize() != 0 {
		update["safety_order_size"] = strategy.GetSafetyOrderSize()
		oldValues.SafetyOrderSize = dataRead.SafetyOrderSize
		newValues.SafetyOrderSize = strategy.GetSafetyOrderSize()
	}
	if strategy.GetMaxSafetyTradeCount() != "" {
		update["max_safety_trade_count"] = strategy.GetMaxSafetyTradeCount()
		oldValues.MaxSafetyTradeCount = dataRead.MaxSafetyTradeCount
		newValues.MaxSafetyTradeCount = strategy.GetMaxSafetyTradeCount()
	}
	if strategy.GetMaxActiveSafetyTradeCount() != "" {
		update["max_active_safety_trade_count"] = strategy.GetMaxActiveSafetyTradeCount()
		oldValues.MaxActiveSafetyTradeCount = dataRead.MaxActiveSafetyTradeCount
		newValues.MaxActiveSafetyTradeCount = strategy.GetMaxActiveSafetyTradeCount()
	}
	if strategy.GetPriceDevation() != "" {
		update["price_devation"] = strategy.GetPriceDevation()
		oldValues.PriceDevation = dataRead.PriceDevation
		newValues.PriceDevation = strategy.GetPriceDevation()
	}
	if strategy.GetSafetyOrderVolumeScale() != "" {
		update["safety_order_volume_scale"] = strategy.GetSafetyOrderVolumeScale()
		oldValues.SafetyOrderVolumeScale = dataRead.SafetyOrderVolumeScale
		newValues.SafetyOrderVolumeScale = strategy.GetSafetyOrderVolumeScale()
	}
	if strategy.GetSafetyOrderStepScale() != "" {
		update["safety_order_step_scale"] = strategy.GetSafetyOrderStepScale()
		oldValues.SafetyOrderStepScale = dataRead.SafetyOrderStepScale
		newValues.SafetyOrderStepScale = strategy.GetSafetyOrderStepScale()
	}
	if strategy.GetTakeProfit() != "" {
		update["take_profit"] = strategy.GetTakeProfit()
		oldValues.TakeProfit = dataRead.TakeProfit
		newValues.TakeProfit = strategy.GetTakeProfit()
	}
	if strategy.GetTargetProfit() != "" {
		update["target_profit"] = strategy.GetTargetProfit()
		oldValues.TargetProfit = dataRead.TargetProfit
		newValues.TargetProfit = strategy.GetTargetProfit()
	}
	if strategy.GetAllocateFundsToStrategy() != "" {
		update["allocate_funds_to_strategy"] = strategy.GetAllocateFundsToStrategy()
		oldValues.AllocateFundsToStrategy = dataRead.AllocateFundsToStrategy
		newValues.AllocateFundsToStrategy = strategy.GetAllocateFundsToStrategy()
	}
	update["version"] = dataRead.Version + 1
	oldValues.Version = dataRead.Version
	newValues.Version = dataRead.Version + 1

	if strategy.GetStock() != nil {
		update["stock"] = strategy.GetStock()
		oldValues.Stock = dataRead.Stock
		newValues.Stock = strategy.GetStock()
	}

	//fmt.Println(update)

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
	eventData := EventHistoryItem{
		OperationType: "update",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "strategy",
		Name:          dataRead.StrategyName,
		UserId:        dataRead.UserId,
		StrategyId:    oid.Hex(),
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

	return &strategypb.UpdateStrategyRes{
		Strategy: &strategypb.Strategy{
			Id:                        decoded.Id.Hex(),
			StrategyName:              decoded.StrategyName,
			SelectedExchange:          decoded.SelectedExchange,
			StrategyType:              decoded.StrategyType,
			StartOrderType:            decoded.StartOrderType,
			DealStartCondition:        decoded.DealStartCondition,
			BaseOrderSize:             decoded.BaseOrderSize,
			SafetyOrderSize:           decoded.SafetyOrderSize,
			MaxSafetyTradeCount:       decoded.MaxSafetyTradeCount,
			MaxActiveSafetyTradeCount: decoded.MaxActiveSafetyTradeCount,
			PriceDevation:             decoded.PriceDevation,
			SafetyOrderVolumeScale:    decoded.SafetyOrderVolumeScale,
			SafetyOrderStepScale:      decoded.SafetyOrderStepScale,
			TakeProfit:                decoded.TakeProfit,
			TargetProfit:              decoded.TargetProfit,
			AllocateFundsToStrategy:   decoded.AllocateFundsToStrategy,
			UserId:                    decoded.UserId,
			Version:                   decoded.Version,
			Status:                    decoded.Status,
			Stock:                     decoded.Stock,
		},
	}, nil
}

func (s *StrategyServiceServer) StartBot(ctx context.Context, req *strategypb.StartBotReq) (*strategypb.StartBotRes, error) {
	strategyId := req.GetStrategyId()
	stocks := req.GetStocks()
	//fmt.Println(strategyId)
	//fmt.Println(stocks)

	oid, err := primitive.ObjectIDFromHex(strategyId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	resultStrategyRead := strategydb.FindOne(ctx, bson.M{"_id": oid})

	// Create an empty ExchangeItem to write our decode result to
	strategyData := StrategyItem{}
	// decode and write to strategyData
	if err := resultStrategyRead.Decode(&strategyData); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", strategyId, err))
	}

	//fmt.Println(len(stocks))
	var insert []interface{}
	for _, v := range stocks {
		//fmt.Println(stocks[i].StockName)
		//fmt.Println(v.StockName)
		deal := bson.M{
			"strategy_id": strategyId,
			"version":     strategyData.Version,
			"user_id":     strategyData.UserId,
			"stock":       v.StockName,
			"status":      "running",
		}
		insert = append(insert, deal)
	}

	//fmt.Println(insert)

	insertManyResult, err := dealsdb.InsertMany(mongoCtx, insert)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	fmt.Print("Inserted Deal ID's: ")
	fmt.Println(insertManyResult.InsertedIDs)

	decoded := StrategyItem{}
	filter := bson.M{"_id": oid}
	resultUpdateStrategy := strategydb.FindOneAndUpdate(ctx, filter, bson.M{"$set": bson.M{"status": "running"}}, options.FindOneAndUpdate().SetReturnDocument(1))
	err = resultUpdateStrategy.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find strategy with supplied ID: %v", err),
		)
	}
	eventData := EventHistoryItem{
		OperationType: "startbot",
		Timestamp:     time.Now().Format(time.RFC3339),
		Db:            "hedgina_algobot",
		Collection:    "strategy",
		Name:          strategyData.StrategyName,
		UserId:        strategyData.UserId,
		StrategyId:    oid.Hex(),
		OldValue: StrategyItem{
			Status: strategyData.Status,
		},
		NewValue: StrategyItem{
			Stock:  stocks,
			Status: "running",
		},
	}
	_, errEventHistory := eventhistorydb.InsertOne(mongoCtx, eventData)
	if errEventHistory != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", errEventHistory),
		)
	}

	return &strategypb.StartBotRes{
		Success: true,
	}, nil
}

func (s *StrategyServiceServer) ListDeals(req *strategypb.ListDealReq, stream strategypb.StrategyService_ListDealsServer) error {
	userIdQuery := req.GetUserId()
	if len(userIdQuery) == 0 {
		return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not find UserId in Req"))
	}

	// Initiate a StrategyItem type to write decoded data to
	data := &DealItem{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := dealsdb.Find(context.Background(), bson.M{"user_id": userIdQuery})
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

		strategyData := &StrategyItem{}
		strategyId, err := primitive.ObjectIDFromHex(data.StrategyId)

		resultReadStrategy := strategydb.FindOne(mongoCtx, bson.M{"_id": strategyId, "version": data.Version})

		if err := resultReadStrategy.Decode(&strategyData); err != nil {

			fmt.Println("Inside error not nil")
			fmt.Println(err)
			if err == mongo.ErrNoDocuments {
				fmt.Println(fmt.Sprintf("No Strategy in strategy collection. Check Strategy revision for StrategyId: %s , Version: %d", data.StrategyId, data.Version))
				resultReadStrategyRevisions := strategy_revisionsdb.FindOne(mongoCtx, bson.M{"strategy_id": data.StrategyId, "version": data.Version})
				if errReadRevisions := resultReadStrategyRevisions.Decode(&strategyData); errReadRevisions != nil {
					if errReadRevisions == mongo.ErrNoDocuments {
						fmt.Println(fmt.Sprintf("No Strategy in strategy revisions collection for StrategyId: %s , Version: %d", data.StrategyId, data.Version))
					} else {
						//fmt.Println("not in error no documents")
						return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
					}
				}
			} else {
				//fmt.Println("not in error no documents")
				return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
			}
			//return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Strategy with Object Id %s: %v", oid, err))
		}
		// If no error is found send exchange over stream
		stream.Send(&strategypb.ListDealRes{
			Deal: &strategypb.Deal{
				Id:         data.Id.Hex(),
				StrategyId: data.StrategyId,
				Version:    data.Version,
				Stock:      data.Stock,
				UserId:     data.UserId,
				Status:     data.Status,
			},
			Strategy: &strategypb.Strategy{
				StrategyName:              strategyData.StrategyName,
				SelectedExchange:          strategyData.SelectedExchange,
				StrategyType:              strategyData.StrategyType,
				StartOrderType:            strategyData.StartOrderType,
				DealStartCondition:        strategyData.DealStartCondition,
				BaseOrderSize:             strategyData.BaseOrderSize,
				SafetyOrderSize:           strategyData.SafetyOrderSize,
				MaxSafetyTradeCount:       strategyData.MaxSafetyTradeCount,
				MaxActiveSafetyTradeCount: strategyData.MaxActiveSafetyTradeCount,
				PriceDevation:             strategyData.PriceDevation,
				SafetyOrderVolumeScale:    strategyData.SafetyOrderVolumeScale,
				SafetyOrderStepScale:      strategyData.SafetyOrderStepScale,
				TakeProfit:                strategyData.TakeProfit,
				TargetProfit:              strategyData.TargetProfit,
				AllocateFundsToStrategy:   strategyData.AllocateFundsToStrategy,
				UserId:                    strategyData.UserId,
				Version:                   strategyData.Version,
				Status:                    strategyData.Status,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

type StrategyServiceServer struct{}
type StrategyItem struct {
	Id                        primitive.ObjectID  `bson:"_id,omitempty"`
	StrategyName              string              `bson:"strategy_name"`
	SelectedExchange          string              `bson:"selected_exchange"`
	StrategyType              string              `bson:"strategy_type"`
	StartOrderType            string              `bson:"start_order_type"`
	DealStartCondition        string              `bson:"deal_start_condition"`
	BaseOrderSize             float64             `bson:"base_order_size"`
	SafetyOrderSize           float64             `bson:"safety_order_size"`
	MaxSafetyTradeCount       string              `bson:"max_safety_trade_count"`
	MaxActiveSafetyTradeCount string              `bson:"max_active_safety_trade_count"`
	PriceDevation             string              `bson:"price_devation"`
	SafetyOrderVolumeScale    string              `bson:"safety_order_volume_scale"`
	SafetyOrderStepScale      string              `bson:"safety_order_step_scale"`
	TakeProfit                string              `bson:"take_profit"`
	TargetProfit              string              `bson:"target_profit"`
	AllocateFundsToStrategy   string              `bson:"allocate_funds_to_strategy"`
	UserId                    string              `bson:"user_id"`
	Version                   int64               `bson:"version"`
	Status                    string              `bson:"status"`
	Stock                     []*strategypb.Stock `bson:"stock"`
}

type StrategyRevisionItem struct {
	Id                        primitive.ObjectID  `bson:"_id,omitempty"`
	StrategyName              string              `bson:"strategy_name"`
	SelectedExchange          string              `bson:"selected_exchange"`
	StrategyType              string              `bson:"strategy_type"`
	StartOrderType            string              `bson:"start_order_type"`
	DealStartCondition        string              `bson:"deal_start_condition"`
	BaseOrderSize             float64             `bson:"base_order_size"`
	SafetyOrderSize           float64             `bson:"safety_order_size"`
	MaxSafetyTradeCount       string              `bson:"max_safety_trade_count"`
	MaxActiveSafetyTradeCount string              `bson:"max_active_safety_trade_count"`
	PriceDevation             string              `bson:"price_devation"`
	SafetyOrderVolumeScale    string              `bson:"safety_order_volume_scale"`
	SafetyOrderStepScale      string              `bson:"safety_order_step_scale"`
	TakeProfit                string              `bson:"take_profit"`
	TargetProfit              string              `bson:"target_profit"`
	AllocateFundsToStrategy   string              `bson:"allocate_funds_to_strategy"`
	UserId                    string              `bson:"user_id"`
	Version                   int64               `bson:"version"`
	Status                    string              `bson:"status"`
	Stock                     []*strategypb.Stock `bson:"stock"`
	StrategyId                string              `bson:"strategy_id"`
}

type Stock struct {
	StockName string `bson:"stock_name"`
}

type DealItem struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	StrategyId string             `bson:"strategy_id"`
	Version    int64              `bson:"version"`
	Stock      string             `bson:"stock"`
	UserId     string             `bson:"user_id"`
	Status     string             `bson:"status"`
}

type EventHistoryItem struct {
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

var db *mongo.Client
var strategydb *mongo.Collection
var strategy_revisionsdb *mongo.Collection
var dealsdb *mongo.Collection
var eventhistorydb *mongo.Collection
var mongoCtx context.Context

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50052...")

	// Start our listener, 50052 is the default gRPC port
	listener, err := net.Listen("tcp", ":50052")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :50052: %v", err)
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
	mongoDB := db.Database("hedgina_algobot")
	strategydb = mongoDB.Collection("strategy")
	strategy_revisionsdb = mongoDB.Collection("strategy_revisions")
	dealsdb = mongoDB.Collection("deals")
	eventhistorydb = mongoDB.Collection("eventhistory_strategy")
	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50052")

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
