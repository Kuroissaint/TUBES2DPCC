package main

import "errors"

// Driver mewakili struktur data driver
type Driver struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Status string  `json:"status"`
}

// DispatchService adalah interface untuk logika bisnis
type DispatchService interface {
	FindNearestDriver(lat, lng float64) (*Driver, error)
	AssignDriver(driverID string, status string) (*Driver, error)
}

type DispatchServiceImpl struct {
	Repo DispatchRepository
}

func NewDispatchService(repo DispatchRepository) DispatchService {
	return &DispatchServiceImpl{Repo: repo}
}

func (s *DispatchServiceImpl) FindNearestDriver(lat, lng float64) (*Driver, error) {
	if lat == 0 && lng == 0 {
		return nil, errors.New("lokasi tidak valid")
	}
	
	return &Driver{
		ID:     "DRV-001",
		Name:   "Alex Marquez",
		Lat:    lat,
		Lng:    lng,
		Status: "available",
	}, nil
}

func (s *DispatchServiceImpl) AssignDriver(driverID string, status string) (*Driver, error) {
	finalStatus := status
	if status == "pending" || status == "" {
		finalStatus = "assigned"
	}

	driver := &Driver{
		ID:     driverID,
		Name:   "Marc Marquez",
		Status: finalStatus,
	}

	// Nanti di sini panggil repo jika DB sudah ada:
	// _ = s.Repo.UpdateDriverStatus(driverID, finalStatus)

	return driver, nil
}