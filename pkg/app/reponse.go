// username: vonhng
// create_time: 2019/12/28 - 20:46
// mail: vonhng.feng@gmail.com
package app

import (
	"doja/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	GinContext *gin.Context
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.GinContext.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.MsgMap[errCode],
		Data: data,
	})
}
