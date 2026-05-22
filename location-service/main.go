package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	// Inisialisasi Repository dan Service (seperti di service lain)
	// var repo LocationRepository = ... (bisa pakai mock dulu atau DB asli)
	// service := NewLocationService(repo)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, WebResponse{
			Status:  "success",
			Message: "Location Service is running",
		})
	})

	router.POST("/track", func(c *gin.Context) {
		// Nanti di sini kamu tinggal memanggil service.TrackLocation()
		// locationData, err := service.TrackLocation()
		// Lalu masukkan hasilnya ke WebResponse
        
		c.JSON(http.StatusOK, gin.H{
            "message": "Handler ini akan segera dihubungkan dengan service",
        })
	})

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8002"
	}
	router.Run(":" + port)
}