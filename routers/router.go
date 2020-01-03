// username: vonhng
// create_time: 2019/12/28 - 20:00
// mail: vonhng.feng@gmail.com
package routers

import (
	"github.com/gin-gonic/gin"
	"vonhng/doja/routers/api"
)

func InitRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/login", api.Login)
	r.GET("/query", api.DBQuery)
	r.POST("/update", api.DBUpdate)
	r.POST("/insert", api.DBInsert)
	r.POST("/delete", api.DBDelete)
	return r
}
