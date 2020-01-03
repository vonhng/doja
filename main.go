// username: vonhng
// create_time: 2019/12/3 - 09:56
// mail: vonhng.feng@gmail.com
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vonhng/doja/models"
	"vonhng/doja/pkg/config"
	"vonhng/doja/pkg/logging"
	"vonhng/doja/routers"
)

func init() {
	config.ParseConfig()
	logging.SetUpLog()
	err := models.Connect()
	if err != nil {
		logging.Fatal("connect mongodb failed")
	}
}

func main() {
	defer models.Disconnect()
	gin.SetMode(gin.DebugMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf("%s:%s", config.Setting.Web.Address, config.Setting.Web.Port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	logging.Debug(fmt.Sprintf("start http server listening %s", endPoint))
	//routersInit.Run(":8000")
	_ = server.ListenAndServe()
}
