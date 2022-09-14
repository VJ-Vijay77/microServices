package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Msg string	`json:"msg"`
}
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
	 PostgresURL := "http://172.24.0.4:8080/get/"+data.Name

	 request,err := http.NewRequest("GET",PostgresURL,nil)
	 if err != nil {
		log.Println(err)
		c.JSON(500,"error in new request")
		return
	}
	request.Header.Set("Content-Type","application/json")

	 client := &http.Client{}
	 response,err := client.Do(request)
	 if err != nil {
		log.Println(err)
		c.JSON(500,"error in Do request")
		return
	 }

	 if response.StatusCode == 404 {
		log.Println("Couldnt Find data")
		return
	 }
	 if response.StatusCode != 200 {
		log.Println("Some mistake")
		return
	 }

	 var res Response
	 err = json.NewDecoder(response.Body).Decode(&res)
	 if err != nil{
		log.Println(err)
		c.JSON(400,"error in decode")
		return
	}

	c.JSON(200,gin.H{
		"msg":res,
	})
	 

}
