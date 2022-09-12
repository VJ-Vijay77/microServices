package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"postgres/models"
)

type Config struct{}
var PORT = ":8082"
func main() {
	app := Config{}
	db,err := sqlx.Connect("postgres",os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("Connection to Postgres success....")
	}

	db.Exec(models.Schema)

	r := gin.Default()
	app.Routes(r)
	r.Run(PORT)

}