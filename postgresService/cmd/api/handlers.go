package main

import "github.com/gin-gonic/gin"


func(app *Config) Postgres(c *gin.Context) {
	c.JSON(200,"This is postgres..")
}