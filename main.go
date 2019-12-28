// username: vonhng
// create_time: 2019/12/3 - 09:56
// mail: qianyong.feng@woqutech.com
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"vonhng/doja/routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", 8000)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	//routersInit.Run(":8000")
	server.ListenAndServe()
}
