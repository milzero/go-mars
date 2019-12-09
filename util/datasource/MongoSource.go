package datasource

import (
	"StreamChannelSwitch/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoSource struct {
	url         string
	clinet      *mongo.Client
	collections map[string]*mongo.Collection
	datebase    string
}

func NewMongoSource(config *config.Config) *MongoSource {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		panic("can not connect to mongo")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return &MongoSource{clinet: client,
		url:         config.MongoUrl,
		datebase:    config.MongoDBName,
		collections: make(map[string]*mongo.Collection)}
}

func (r *MongoSource) Ping() bool {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := r.clinet.Ping(ctx, readpref.Primary())
	if err != nil {
		return false
	}
	return true
}

func (r *MongoSource) InserOneDoc(collectionName string, doc interface{}) error {

	collection, ok := r.collections[collectionName]
	if ok == false {
		collection = r.clinet.Database(r.datebase).Collection(collectionName)
		r.collections[collectionName] = collection
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}
