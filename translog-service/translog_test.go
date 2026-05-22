package main

import (
	"testing"
	"translog-service/mocks"
	"github.com/golang/mock/gomock"
)

func TestValidateStatusTransition(t *testing.T) {
	svc := NewTranslogService(nil)

	err := svc.ValidateStatusTransition("SEARCHING", "COMPLETED")
	if err == nil {
		t.Errorf("Ekspektasi error untuk transisi SEARCHING -> COMPLETED, tapi sukses")
	}

	err = svc.ValidateStatusTransition("SEARCHING", "IN_PROGRESS")
	if err != nil {
		t.Errorf("Ekspektasi sukses untuk transisi SEARCHING -> IN_PROGRESS, tapi dapat error: %v", err)
	}
}

func TestCreateTransportOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTranslogRepository(ctrl)
	// Ekspektasi 5 argumen sesuai dengan fungsi SaveOrder di repository
	mockRepo.EXPECT().SaveOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	svc := NewTranslogService(mockRepo)
	order, err := svc.CreateTransportOrder()

	if err != nil {
		t.Errorf("Tidak ekspektasi error, tapi dapat: %v", err)
	}
	if order.Status != "SEARCHING" {
		t.Errorf("Ekspektasi status SEARCHING, dapat: %s", order.Status)
	}
}