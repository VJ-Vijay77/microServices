package main

import (
	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

type Config struct {
}

func main() {
	app := Config{}
	r := gin.Default()
	app.Routes(r)
	r.Run(PORT)
}
