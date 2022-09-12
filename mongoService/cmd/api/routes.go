package main

import (
	"github.com/gin-gonic/gin"
)

func(app *Config)  Routess(r *gin.Engine) {

	r.GET("/mongo", app.Mongo)

}
