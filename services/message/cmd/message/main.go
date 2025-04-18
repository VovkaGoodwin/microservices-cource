package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
