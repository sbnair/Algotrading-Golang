package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getAssets() error {

	userdbResult := userdb.FindOne(mongoCtx, bson.M{"first_name": "admin", "last_name": "user"})

	userData := UserItem{}

	if err := userdbResult.Decode(&userData); err != nil {
		return errors.New(fmt.Sprintf("Could not find User with 'first_name' 'admin' and 'last_name' 'user' %v", err))
	}

	exchangeDbResult := exchangedb.FindOne(mongoCtx, bson.M{"user_id": userData.UserId, "selected_exchange": "Alpaca"})
	// Create an empty ExchangeItem to write our decode result to
	exchangeData := ExchangeItem{}
	// decode and write to data
	if err := exchangeDbResult.Decode(&exchangeData); err != nil {
		return errors.New(fmt.Sprintf("Could not find Exchange for Admin User with User Id %s %v", userData.UserId, err))
	}

	if exchangeData.SelectedExchange == "Alpaca" {
		os.Setenv(common.EnvApiKeyID, exchangeData.ApiKey)
		os.Setenv(common.EnvApiSecretKey, exchangeData.ApiSecret)
		if exchangeData.ExchangeType == "paper_trading" {
			alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
		} else if exchangeData.ExchangeType == "live_trading" {
			alpaca.SetBaseUrl("https://api.alpaca.markets")
		} else {
			return errors.New("Exchange Type for Admin User is neither paper_trading nor live_trading")
		}
		alpacaClient := alpaca.NewClient(common.Credentials())
		var status string = "active"
		assetstatus := &status
		assets, err := alpacaClient.ListAssets(assetstatus)
		if err != nil {
			return errors.New("Could not fetch Assets List from Alpaca")
		}

		var insert []interface{}

		for _, v := range assets {
			asset := bson.M{
				"_id":            v.ID,
				"name":           v.Name,
				"exchange":       v.Exchange,
				"asset_class":    v.Class,
				"symbol":         v.Symbol,
				"status":         v.Status,
				"tradable":       v.Tradable,
				"marginable":     v.Marginable,
				"shortable":      v.Shortable,
				"easy_to_borrow": v.EasyToBorrow,
			}
			insert = append(insert, asset)
		}

		ordered := false
		insertManyOptions := &options.InsertManyOptions{
			Ordered: &ordered,
		}
		insertManyResult, err := assetsdb.InsertMany(mongoCtx, insert, insertManyOptions)
		//fmt.Println(len(insertManyResult.InsertedIDs))
		if len(insertManyResult.InsertedIDs) > 0 {
			fmt.Print("Inserted Asset ID's: ")
			fmt.Println(insertManyResult.InsertedIDs)
		} else if err != nil {
			isDuplicate := IsDup(err)
			//fmt.Println(isDuplicate)
			if isDuplicate == false {
				return errors.New("Mongo Insert Failed")
			} else {
				fmt.Println("No new asset to insert in Mongo DB")
			}
		}

	} else {
		return errors.New("Admin User's selected_exchange is not Alpaca")
	}
	return nil
}

func IsDup(err error) bool {
	var e mongo.BulkWriteException
	e = err.(mongo.BulkWriteException)
	for _, we := range e.WriteErrors {
		if we.Code != 11000 {
			return false
		}
	}

	return true
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

type UserItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `bson:"first_name"`
	Last_name     *string            `bson:"last_name"`
	Password      *string            `bson:"password"`
	Email         *string            `bson:"email"`
	Phone         *string            `bson:"phone"`
	Token         *string            `bson:"token"`
	Refresh_token *string            `bson:"refresh_token"`
	Created_at    time.Time          `bson:"created_at"`
	Updated_at    time.Time          `bson:"updated_at"`
	UserId        string             `bson:"user_id"`
}

var db *mongo.Client
var userdb *mongo.Collection
var exchangedb *mongo.Collection
var assetsdb *mongo.Collection
var mongoCtx context.Context

func main() {

	//record starttime of request
	startTime := time.Now()

	//Uncomment to run locally
	//os.Setenv("MONGODB_URL", "mongodb://127.0.0.1:27017")

	MONGODB_URL := os.Getenv("MONGODB_URL")

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoCtx = context.Background()
	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(MONGODB_URL))
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
	userdb = mongoDb.Collection("user")
	exchangedb = mongoDb.Collection("exchange")
	assetsdb = mongoDb.Collection("assets")

	err = getAssets()
	if err != nil {
		db.Disconnect(mongoCtx)
		log.Fatalf("%v\n", err)
		//fmt.Println(err)
	}

	db.Disconnect(mongoCtx)
	fmt.Println("Disconnected from Mongo")
	diff := time.Since(startTime)
	fmt.Print("Time taken for the operation: ")
	fmt.Println(diff)

}
