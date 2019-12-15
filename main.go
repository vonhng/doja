// username: vonhng
// create_time: 2019/12/3 - 09:56
// mail: qianyong.feng@woqutech.com
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("auth").C("test")

	result := Login{}
	err = c.Find(bson.M{"user": "fqy"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Password)
}
