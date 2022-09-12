package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)
var PORT = ":8080"

type Config struct {
	Db *sqlx.DB
}

func main() {
	var app *Config
	r := gin.Default()
	app.Routes(r)
	r.Run(PORT)
}
