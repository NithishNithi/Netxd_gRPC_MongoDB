package netxdconfig

import (
	netxdconstants "Netxd_gRPC_MongoDb/netxd_grpc_mongo_server/netxd_constants"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoConnection := options.Client().ApplyURI(netxdconstants.ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return mongoClient, nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
