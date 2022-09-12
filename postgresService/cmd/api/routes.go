package main

import "github.com/gin-gonic/gin"


func(app *Config) Routes(r *gin.Engine) {
	r.GET("/postgres")
}