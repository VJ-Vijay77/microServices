package main

import "github.com/gin-gonic/gin"


func(app *Config) Mongo(c *gin.Context) {
	c.JSON(200,"This is Mongo Db")
}