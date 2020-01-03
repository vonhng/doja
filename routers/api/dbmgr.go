// username: vonhng
// create_time: 2020/1/3 - 08:22
// mail: vonhng.feng@gmail.com
package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"vonhng/doja/models"
	"vonhng/doja/pkg/app"
	"vonhng/doja/pkg/e"
)

type validation struct {
	Username string
}

func DbMgr(c *gin.Context) {
	g := app.Gin{GinContext: c}
	valid := validator.New()
	username := c.Query("username")

	err := valid.Struct(&validation{Username: username})
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	value := models.Query(username)
	g.Response(http.StatusOK, e.SUCCESS, value)
}
