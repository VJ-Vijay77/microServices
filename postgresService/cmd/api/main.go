package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"postgres/models"
)



type Config struct{
	Db *sqlx.DB
}

var PORT = ":8082"

func main() {
	fmt.Println(os.Getenv("POSTGRES_URI"))
	app := &Config{}
	
	app = ConnectPostgres()		
	
	app.Db.Exec(models.Schema)
	

	r := gin.Default()
	app.Routes(r)
	log.Println("Starting Postgres Server on ", PORT)
	r.Run(PORT)

}
