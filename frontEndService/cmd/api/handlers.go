package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
		app.errorJSON(c.Writer,err)
	}

	client := &http.Client{}
	response,err := client.Do(request)
	if err != nil {
		app.errorJSON(c.Writer,err)
	}

	defer response.Body.Close()


	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(c.Writer,err)
		return
	}else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(c.Writer,errors.New("error calling postgres server"))
		return
	}


	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil{
	app.errorJSON(c.Writer,err)
		return
	}

	if jsonFromService.Error{
		app.errorJSON(c.Writer,err)
		return
	}

	var payload jsonResponse
	payload.Error =false
	payload.Message = "Success"
	payload.Data = jsonFromService.Data

	



}
