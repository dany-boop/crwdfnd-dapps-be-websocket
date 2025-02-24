package main

import (
	"backend/config"
	"backend/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()
	// router.Use(middleware.RateLimit())

	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)

	// router.GET("/ws", handlers.HandleConnections)
	// go handlers.HandleMessages()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
