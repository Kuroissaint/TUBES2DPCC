package main // atau package location jika kamu pakai module terpisah

// LocationRepository adalah interface untuk menyimpan/mengambil data lokasi
type LocationRepository interface {
	SaveLocation(userID string, lat float64, lon float64) error
}