package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.InitDB()
	// Connect to PostgreSQL
	dsn := "host=localhost user=postgres password=mysecret dbname=donationdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	// donationRepo := repositories.NewDonationRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	// donationService := services.NewDonationService(donationRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService)
	// donationHandler := handlers.NewDonationHandler(donationService)
	// campaignHandler := handlers.CampaignHandler{}

	// Setup router
	router := gin.Default()

	// Auth routes
	router.POST("/register", authHandler.RegisterUser)
	router.POST("/login", authHandler.LoginUser)

	// Protected routes (require JWT)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user/profile", authHandler.GetUserProfile)
		// protected.POST("/donate", donationHandler.CreateDonation)
		// protected.GET("/donations", donationHandler.GetAllDonations)
	}

	// Public routes
	// router.GET("/campaigns", campaignHandler.GetAllCampaigns)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
