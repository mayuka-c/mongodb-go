package db

import (
	"context"
	"fmt"

	"github.com/mayuka-c/mongodb-go/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dataBase   = "test"
	collection = "content-data"
)

type InterfaceDB interface {
	InsertOne(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error)
	InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error)
}

type mongoClient struct {
	client *mongo.Client
}

func NewDBConnection(ctx context.Context, dbConfig config.DBConfiguration) *mongoClient {

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbConfig.Endpoint))
	if err != nil {
		panic(err)
	}
	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("connected successfully")

	return &mongoClient{
		client: client,
	}
}

// InsertOne is a user defined method, used to insert
// documents into collection returns result of InsertOne
// and error if any.
func (m *mongoClient) InsertOne(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := m.client.Database(dataBase).Collection(collection)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

// InsertMany is a user defined method, used to insert
// documents into collection returns result of
// InsertMany and error if any.
func (m *mongoClient) InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	// select database and collection ith Client.Database
	// method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertMany accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertMany(ctx, docs)
	return result, err
}
