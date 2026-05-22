package main

import (
	"testing"
	"github.com/golang/mock/gomock"
	// Sesuaikan "dispatch-service/mocks" dengan nama di go.mod kamu
	"dispatch-service/mocks" 
)

func TestFindNearestDriver_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDispatchRepository(ctrl)
	service := NewDispatchService(mockRepo)

	lat := -6.200000
	lng := 106.816666
	t.Logf("Mencari driver terdekat untuk lokasi: %f, %f", lat, lng)

	// Memanggil lewat service
	driver, err := service.FindNearestDriver(lat, lng)
	if err != nil {
		t.Fatalf("Fungsi error: %v", err)
	}

	if driver.ID == "" {
		t.Errorf("Unit Test FAILED: ID Driver kosong")
	}
	if driver.Name != "Alex Marquez" {
		t.Errorf("Unit Test FAILED: Nama driver tidak sesuai, dapat: %s", driver.Name)
	}
}

func TestAssignDriverStatus_Unit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDispatchRepository(ctrl)
	service := NewDispatchService(mockRepo)

	driverID := "DRV-001"
	status := "pending"

	// Memanggil lewat service
	driver, err := service.AssignDriver(driverID, status)
	if err != nil {
		t.Fatalf("Fungsi error: %v", err)
	}

	if driver.Status == "pending" {
		t.Fatalf("Unit Test FAILED: Logika penugasan gagal, status tetap: pending")
	}

	if driver.Status != "assigned" {
		t.Errorf("Expect status 'assigned', got '%s'", driver.Status)
	}
}