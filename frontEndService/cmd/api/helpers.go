package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any	`json:"data,omitempty"`
}




func (app *Config) errorJSON(c gin.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var jsRes jsonResponse
	jsRes.Error = true
	jsRes.Message = err.Error()

	return app.writeJSON(c,statusCode,jsRes)

}

func(app *Config) writeJSON(c gin.ResponseWriter,status int,data any,headers ...http.Header) error {
	out,err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key,value := range headers[0] {
			c.Header()[key] = value
		}
	}

	c.Header().Set("Content-Type","application/json")
	c.WriteHeader(status)
	_,err = c.Write(out)
	if err != nil {
		return err
	}

	return nil
}