package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any	`json:"data,omitempty"`
}




func (app *Config) errorJSON(c gin.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var jsRes jsonResponse
	jsRes.Error = true
	jsRes.Message = err.Error()

	return 

}
