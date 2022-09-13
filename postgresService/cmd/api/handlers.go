package main

import (
	"postgres/models"

	"github.com/gin-gonic/gin"
)


func(app *Config) Postgres(c *gin.Context) {
	c.JSON(200,"This is postgres..")
}

func(app *Config) Insert(c *gin.Context) {
	var users models.Users
	
	if err := c.ShouldBindJSON(&users); err != nil{
		c.JSON(400,"Could not get data frpm struct")
		return
	}
	

	_,err := app.Db.Exec("INSERT INTO users(name,email,password,age)VALUES($1,$2,$3,$4)",users.Name,users.Email,users.Password,users.Age)
	if err != nil {
		c.JSON(400,"Could not insert to postgres!")
		return
	}

	c.JSON(200,gin.H{
		"status":"Inserted Successfully",
		"inserted":users,
})
}