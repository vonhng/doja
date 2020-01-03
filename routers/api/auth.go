// username: vonhng
// create_time: 2019/12/28 - 20:11
// mail: vonhng.feng@gmail.com
package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"vonhng/doja/models"
	"vonhng/doja/pkg/app"
	"vonhng/doja/pkg/e"
)

type Auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func (auth *Auth) check() bool {
	password := models.Query(auth.Username)
	return password == auth.Password
}

func Login(c *gin.Context) {
	g := app.Gin{GinContext: c}
	valid := validator.New()

	username := c.Query("username")
	password := c.Query("password")
	log.Printf("%s, %s", username, password)

	auth := Auth{Username: username, Password: password}
	err := valid.Struct(&auth)

	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if isExist := auth.check(); !isExist {
		g.Response(http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, map[string]string{
			"username": username,
			"password": password,
		})
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": "小猫咪，嘿嘿",
	})
	return
}
