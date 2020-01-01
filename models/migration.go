// username: vonhng
// create_time: 2019/12/30 - 00:05
// mail: vonhng.feng@gmail.com
package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"vonhng/doja/pkg/config"
	"vonhng/doja/pkg/logging"
)

type DBAuth struct {
	Username string `json:"username",yaml:"username",valid:"Required; MaxSize(50)"`
	Password string `json:"password",yaml:"password",valid:"Required; MaxSize(50)"`
}

func (c *DBAuth) Connect() {
	ConnectionUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=SCRAM-SHA-1",
		config.Setting.DB.Username,
		config.Setting.DB.Password,
		config.Setting.DB.Host,
		config.Setting.DB.Port,
		config.Setting.DB.DBName)
	session, err := mgo.Dial(ConnectionUrl)
	if err != nil {
		logging.Error(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	query := session.DB(config.Setting.DB.DBName).C("test")
	OneAuth := &DBAuth{}
	err = query.Find(bson.M{"username": "fqy"}).One(OneAuth)
	if err != nil {
		logging.Error(err)
	}
	logging.Info(OneAuth.Password)
}
