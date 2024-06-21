package main

import (
	"github.com/gin-gonic/gin"
	"github.com/infotitanz/dancerapy-api/internal/controller"
	"github.com/infotitanz/dancerapy-api/internal/database"
	"github.com/infotitanz/dancerapy-api/internal/middleware"
	"github.com/infotitanz/dancerapy-api/internal/model"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set logrus to use JSON formatter
	log.SetFormatter(&log.JSONFormatter{})

	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	err := database.Database.AutoMigrate(&model.User{})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error migrating User model")
		return
	}
	err = database.Database.AutoMigrate(&model.Training{})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error migrating Training model")
		return
	}
}

func serveApplication() {
	router := gin.Default()

	router.GET("/api/health", controller.HealthCheck)

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	// Protected routes
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("trainings", controller.AddTraining)
	protectedRoutes.GET("trainings", controller.GetUserTrainings)
	protectedRoutes.PUT("trainings/:id", controller.UpdateTraining)

	err := router.Run(":8080")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error starting server")
		return
	}
	log.WithFields(log.Fields{
		"port": "8080",
	}).Info("Server running on port 8080")
}
