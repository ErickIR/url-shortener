package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:p%40ssw0rd@devcluster.itpxz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return err
	}

	fmt.Println(databases)
	return nil
}
