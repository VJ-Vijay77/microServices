package main

import "github.com/gin-gonic/gin"


func(app *Config) Routes(r *gin.Engine) {
	r.GET("/postgres",app.Postgres)
	r.POST("/insert",app.Insert)
	r.GET("/getone/:name",app.GetOne)
	r.POST("/get",app.GetOnePost)
}