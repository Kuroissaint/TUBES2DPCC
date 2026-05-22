package main

import (
	"testing"
	"shop-order-service/mocks"
	"github.com/golang/mock/gomock"
)

func TestAddToCartLogic(t *testing.T) {
	cart := ShoppingCart{Items: []string{}}
	cart.AddToCart("Mie Goreng Spesial")

	if len(cart.Items) != 1 || cart.Items[0] != "Mie Goreng Spesial" {
		t.Errorf("Add To Cart logic failed")
	}
}

func TestCreateShoppingOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockShopOrderRepository(ctrl)
	// Ekspektasi 5 argumen sesuai dengan fungsi SaveCart di repository
	mockRepo.EXPECT().SaveCart(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	svc := NewShopOrderService(mockRepo)
	cart, err := svc.CreateShoppingOrder()

	if err != nil {
		t.Errorf("Tidak ekspektasi error, tapi dapat: %v", err)
	}
	if len(cart.Items) != 2 {
		t.Errorf("Ekspektasi 2 item, tapi dapat: %d", len(cart.Items))
	}
}