package main

import (
	"log"
	"postgres/models"

	"github.com/gin-gonic/gin"
)

func (app *Config) Postgres(c *gin.Context) {
	c.JSON(200, "This is postgres..")
}

func (app *Config) Insert(c *gin.Context) {
	var users models.Users

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, "Could not get data frpm struct")
		return
	}

	_, err := app.Db.Exec("INSERT INTO users(name,email,password,age)VALUES($1,$2,$3,$4)", users.Name, users.Email, users.Password, users.Age)
	if err != nil {
		c.JSON(400, "Could not insert to postgres!")
		return
	}

	c.JSON(200, gin.H{
		"status":   "Inserted Successfully",
		"inserted": users,
	})
}

func (app *Config) GetOne(c *gin.Context) {

	Name := c.Param("name")

	users := models.Users{}
	w := c.Writer

	err := app.Db.Get(&users, "SELECT * FROM users WHERE name=$1", Name)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(404)
		c.JSON(404, "Couldnt find data")
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	_,err = w.Write([]byte(users.Email))
	if err != nil{
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"name":  users.Name,
		"email": users.Email,
		"age":   users.Age,
	})
}
