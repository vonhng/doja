// username: vonhng
// create_time: 2019/12/30 - 00:05
// mail: vonhng.feng@gmail.com
package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"vonhng/doja/pkg/config"
	"vonhng/doja/pkg/logging"
)

type DBAuth struct {
	Username string `json:"username",yaml:"username",valid:"Required; MaxSize(50)"`
	Password string `json:"password",yaml:"password",valid:"Required; MaxSize(50)"`
}

var (
	MgoClient *mongo.Client
	err       error
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

func Query(key string) interface{} {
	err = Ping()
	if err != nil {
		logging.Fatal(err)
	}
	collection := MgoClient.Database(config.Setting.DB.DBName).Collection("login")
	OneAuth := &DBAuth{}
	err := collection.FindOne(context.TODO(), bson.D{{"username", key}}).Decode(&OneAuth)
	if err != nil {
		logging.Fatal(err)
		return nil
	}
	logging.Info(*OneAuth)
	return OneAuth.Password
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
