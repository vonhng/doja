// username: vonhng
// create_time: 2019/12/30 - 00:05
// mail: vonhng.feng@gmail.com
package models

import (
	"context"
	"doja/pkg/config"
	"doja/pkg/logging"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBAuth struct {
	Username string `json:"username",yaml:"username",valid:"Required; MaxSize(50)"`
	Password string `json:"password",yaml:"password",valid:"Required; MaxSize(50)"`
}

var (
	MgoClient  *mongo.Client
	Collection *mongo.Collection
	err        error
)

func Connect() (err error) {
	ConnectionUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=SCRAM-SHA-1",
		config.Setting.DB.Username,
		config.Setting.DB.Password,
		config.Setting.DB.Host,
		config.Setting.DB.Port,
		config.Setting.DB.DBName)
	logging.Info(ConnectionUrl)

	// Set client options
	clientOptions := options.Client().ApplyURI(ConnectionUrl)

	// Connect to MongoDB
	MgoClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logging.Error(err)
	}
	// Check the connection
	err = MgoClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Error(err)
	}
	Collection = MgoClient.Database(config.Setting.DB.DBName).Collection("login")
	fmt.Println("Connected to MongoDB!")
	return err
}

func Ping() error {
	err = MgoClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Error(err)
		err = Connect()
	}
	return err
}

func InsertOne(key, value string) error {
	err = Ping()
	if err != nil {
		logging.Error(err)
		return err
	}
	res, err := Collection.InsertOne(context.Background(), bson.M{"username": key, "password": value})
	if err != nil {
		logging.Error(err)
	}
	logging.Info(res.InsertedID)
	return err
}

func QueryOne(value string) interface{} {
	err = Ping()
	if err != nil {
		logging.Error(err)
		return nil
	}
	OneAuth := &DBAuth{}
	err := Collection.FindOne(context.TODO(), bson.D{{"username", value}}).Decode(&OneAuth)
	if err != nil {
		logging.Error(err)
		return nil
	}
	logging.Info(*OneAuth)
	return OneAuth.Password
}

func UpdateOne(value string) error {
	err = Ping()
	if err != nil {
		logging.Error(err)
		return err
	}
	filter := bson.D{{"username", "fqy"}}
	update := bson.D{{"$set", bson.D{{"password", value}}}}
	updateResult, err := Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logging.Error(err)
	}
	logging.Info("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return err
}

func DeleteOne(key, value string) error {
	err = Ping()
	if err != nil {
		logging.Error(err)
		return err
	}
	// find and delete the document for which the _id field matches id
	// specify the Projection option to only include the name and age fields in the returned document
	opts := options.FindOneAndDelete().SetProjection(bson.D{{"username", key}})
	var deletedDocument bson.M
	err := Collection.FindOneAndDelete(context.TODO(), bson.D{{"password", value}}, opts).Decode(&deletedDocument)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil
		}
		logging.Error(err)
	}
	fmt.Printf(" documents was deleted")
	return err
}

func Disconnect() {
	err = Ping()
	if err != nil {
		logging.Info("Already  connection disconnected")
	}
	err := MgoClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
