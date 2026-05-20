package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Struktur respon standar sesuai kesepakatan tim
type WebResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Contoh data lokasi (Tracking)
type LocationData struct {
	UserID    string  `json:"user_id"`    // Menggunakan snake_case
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	router := gin.Default()

	// 1. Endpoint Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, WebResponse{
			Status:  "success",
			Message: "Location Service is running",
		})
	})

	// 2. Contoh Endpoint Tracking (Simulasi)
	router.POST("/track", func(c *gin.Context) {
		// Contoh penggunaan UUID sesuai kesepakatan
		newID := uuid.New().String()

		response := WebResponse{
			Status: "success",
			Data: LocationData{
				UserID:    newID,
				Latitude:  -6.200000,
				Longitude: 106.816666,
			},
			Message: "Location tracked successfully",
		}
		c.JSON(http.StatusOK, response)
	})

	// Mengambil port dari environment atau default ke 8002
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}

	fmt.Println("Location Service running on port:", port)
	router.Run(":" + port)
}