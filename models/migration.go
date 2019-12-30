// username: vonhng
// create_time: 2019/12/30 - 00:05
// mail: vonhng.feng@gmail.com
package models

import (
	"fmt"
	"github.com/grafana/grafana/pkg/infra/log"
	"gopkg.in/mgo.v2"
	"vonhng/doja/pkg/config"
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
		log.Fatal(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(config.Setting.DB.DBName).C("test")
	OneAuth := &DBAuth{}
	err = c.Find(bson.M{"username": "fqy"}).One(OneAuth)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(OneAuth.Password)
}
