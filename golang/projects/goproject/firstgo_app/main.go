package main

import (
	"firstgo_app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// foo := handlers.Bar()
	// fmt.Println("foo -> %v", foo)
	handlers.Bar()
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
