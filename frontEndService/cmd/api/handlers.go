package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Name string `json:"name"`
	Db   string `json:"db"`
}

func (app *Config) HomePage(c *gin.Context) {
	c.JSON(200, "Hits on HomePage")
}

func (app *Config) Manage(c *gin.Context) {
	var request Message

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(500, "Could not bind")
		return
	}

	switch request.Db {
	case "postgres":
		app.ToPostgres(c,request)
	case "mongo":
		c.JSON(200, "Under working")
	default:
		c.JSON(404, "No such database")
	}
}

func (app *Config) ToPostgres(c *gin.Context ,data Message) {
	 PostgresURL := "http://postgreserver/get/"+data.Name

	 request,err := http.NewRequest("GET",PostgresURL)
	 if err != nil {
		log.Println(err)
		
		
	 }
}
