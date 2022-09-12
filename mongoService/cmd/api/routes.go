package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func(app *Config) Routes() http.Handler {
	r := gin.Default()

	r.GET("/mongo",app.Mongo)
	return r
}