// username: vonhng
// create_time: 2020/1/3 - 08:22
// mail: vonhng.feng@gmail.com
package api

import (
	"doja/models"
	"doja/pkg/app"
	"doja/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type QueryValidation struct {
	Username string
}

type UpdateValidation struct {
	Age string
}

type InsertValidation struct {
	Username string
	Password string
}

type DeleteValidation struct {
	Username string
	Password string
}

func DBQuery(c *gin.Context) {
	g := app.Gin{GinContext: c}
	valid := validator.New()
	username := c.Query("username")

	err := valid.Struct(&QueryValidation{Username: username})
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	value := models.QueryOne(username)
	g.Response(http.StatusOK, e.SUCCESS, value)
}

func DBUpdate(c *gin.Context) {
	g := app.Gin{GinContext: c}

	valid := validator.New()
	age := c.PostForm("age")
	fmt.Println(age)
	err := valid.Struct(&UpdateValidation{Age: age})
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = models.UpdateOne(age)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

func DBInsert(c *gin.Context) {
	g := app.Gin{GinContext: c}
	valid := validator.New()

	username := c.PostForm("username")
	password := c.PostForm("password")
	err := valid.Struct(&InsertValidation{Username: username, Password: password})
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = models.InsertOne(username, password)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

func DBDelete(c *gin.Context) {
	g := app.Gin{GinContext: c}
	valid := validator.New()

	username := c.PostForm("username")
	password := c.PostForm("password")
	err := valid.Struct(&DeleteValidation{Username: username, Password: password})
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = models.DeleteOne(username, password)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}
