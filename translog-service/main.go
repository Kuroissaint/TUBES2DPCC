package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StandardResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type TransportOrder struct {
	OrderID       string  `json:"order_id"`
	UserID        string  `json:"user_id"`
	Status        string  `json:"status"`
	ServiceType   string  `json:"service_type"`
	ItemDimension float64 `json:"item_dimension,omitempty"`
}

func createOrderHandler(svc TranslogService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		order, err := svc.CreateTransportOrder()
		if err != nil {
			http.Error(w, "Gagal memproses pesanan", http.StatusInternalServerError)
			return
		}

		response := StandardResponse{
			Status:  "success",
			Data:    order,
			Message: "Order transportasi berhasil dibuat",
		}

		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	repo := NewTranslogRepository()
	svc := NewTranslogService(repo)

	http.HandleFunc("/api/orders/transport", createOrderHandler(svc))
	fmt.Println("Transport Order Service berjalan di port 8005...")
	log.Fatal(http.ListenAndServe(":8005", nil))
}