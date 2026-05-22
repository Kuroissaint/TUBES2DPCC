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

type ShoppingCart struct {
	OrderID    string   `json:"order_id"`
	UserID     string   `json:"user_id"`
	MerchantID string   `json:"merchant_id"`
	Items      []string `json:"items"`
	Status     string   `json:"status"`
}

func (c *ShoppingCart) AddToCart(item string) {
	c.Items = append(c.Items, item)
}

func createShoppingHandler(svc ShopOrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cart, err := svc.CreateShoppingOrder()
		if err != nil {
			http.Error(w, "Gagal memproses pesanan", http.StatusInternalServerError)
			return
		}

		response := StandardResponse{
			Status:  "success",
			Data:    cart,
			Message: "Pesanan belanja berhasil dikonfirmasi",
		}

		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	repo := NewShopOrderRepository()
	svc := NewShopOrderService(repo)

	http.HandleFunc("/api/orders/shopping", createShoppingHandler(svc))
	fmt.Println("Shopping Order Service berjalan di port 8009...")
	log.Fatal(http.ListenAndServe(":8009", nil))
}