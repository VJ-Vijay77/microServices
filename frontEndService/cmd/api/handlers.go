package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Msg string	`json:"msg"`
}
type PostgresLoad struct{
	Name string `json:"name"`
}
type Message struct {
	Data PostgresLoad `json:"postgres"`
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
		app.ToPostgres(c,request.Data)
	case "mongo":
		c.JSON(200, "Under working")
	default:
		c.JSON(404, "No such database")
	}
}

func (app *Config) ToPostgres(c *gin.Context ,data PostgresLoad) {
	 
	jsonData,_ := json.MarshalIndent(data,"","\t")
	
	URLstring := "http://postgreserver/get"
	request,err := http.NewRequest("POST",URLstring,bytes.NewBuffer(jsonData))
	if err != nil {
		
	}
}
