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

func Connect() *mongo.Client {
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
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func Query(client *mongo.Client, key string) interface{} {
	collection := client.Database(config.Setting.DB.DBName).Collection("login")
	defer Disconnect(client)
	OneAuth := &DBAuth{}
	err := collection.FindOne(context.TODO(), bson.D{{"username", key}}).Decode(&OneAuth)
	if err != nil {
		logging.Fatal(err)
	}
	logging.Info(*OneAuth)
	return OneAuth.Password
}

func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
