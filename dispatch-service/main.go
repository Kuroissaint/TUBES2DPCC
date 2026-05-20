package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// WebResponse adalah standar respon API sesuai kesepakatan tim
type WebResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// DispatchData menggunakan snake_case untuk JSON sesuai standar tim
type DispatchData struct {
	DispatchID string `json:"dispatch_id"`
	OrderID    string `json:"order_id"`
	DriverID   string `json:"driver_id"`
	Status     string `json:"status"`
}

func main() {
	router := gin.Default()

	// Endpoint Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, WebResponse{
			Status:  "success",
			Message: "Dispatch & Matching Service is running",
		})
	})

	// Endpoint untuk simulasi memasangkan driver (Matching)
	router.POST("/dispatch/assign", func(c *gin.Context) {
		// Menggunakan UUID untuk ID sesuai kesepakatan tim
		newDispatchID := uuid.New().String()
		
		response := WebResponse{
			Status: "success",
			Data: DispatchData{
				DispatchID: newDispatchID,
				OrderID:    uuid.New().String(), // Simulasi ID Order
				DriverID:   uuid.New().String(), // Simulasi ID Driver yang didapat
				Status:     "matching_completed",
			},
			Message: "Driver has been successfully matched to the order",
		}
		c.JSON(http.StatusOK, response)
	})

	// Mengambil port dari environment atau default ke 8003
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8003"
	}

	fmt.Println("Dispatch Service running on port:", port)
	router.Run(":" + port)
}