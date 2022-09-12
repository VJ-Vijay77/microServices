package main

import (
	"github.com/gin-gonic/gin"
)

func(app *Config) HomePage(c *gin.Context) {
	 c.JSON(200,"Hits on HomePage")
}