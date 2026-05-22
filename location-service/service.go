package main

import (
	"math"
	"github.com/google/uuid"
)

type LocationData struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Interface mirip ShopOrderService
type LocationService interface {
	TrackLocation() (*LocationData, error)
	CalculateDistance(lat1, lon1, lat2, lon2 float64) float64
	ValidateCoordinates(lat, lon float64) bool
}

// Struct implementasi mirip ShopOrderServiceImpl
type LocationServiceImpl struct {
	Repo LocationRepository
}

// Constructor mirip NewShopOrderService
func NewLocationService(repo LocationRepository) LocationService {
	return &LocationServiceImpl{Repo: repo}
}

func (s *LocationServiceImpl) TrackLocation() (*LocationData, error) {
	newID := uuid.New().String()
	data := &LocationData{
		UserID:    newID,
		Latitude:  -6.200000,
		Longitude: 106.816666,
	}

	// Memanggil repository seperti di shop-order
	// err := s.Repo.SaveLocation(data.UserID, data.Latitude, data.Longitude)
	// if err != nil { return nil, err }

	return data, nil
}

func (s *LocationServiceImpl) CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return math.Round(R*c*10) / 10 
}

func (s *LocationServiceImpl) ValidateCoordinates(lat, lon float64) bool {
	return lat >= -90 && lat <= 90 && lon >= -180 && lon <= 180
}