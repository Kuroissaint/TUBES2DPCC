package main

import (
	"testing"
)

// Unit Test: Hanya menguji logika, tanpa database.
func TestCalculateDistance_Unit(t *testing.T) {
	// Skenario: Menghitung jarak antara dua koordinat
	lat1, lon1 := -6.200000, 106.816666
	lat2, lon2 := -6.210000, 106.826666

	// Supaya tidak error "declared and not used", kita panggil variabelnya di log
	t.Logf("Testing distance between: %f,%f and %f,%f", lat1, lon1, lat2, lon2)

	result := 0.0 // Simulasi hasil karena fungsi belum ada
	expected := 1.5 
	
	if result != expected {
		t.Errorf("Unit Test FAILED: Logika belum diimplementasi. Got %f, want %f", result, expected)
	}
}

func TestValidateCoordinates_Unit(t *testing.T) {
	invalidLat := 100.0
	
	// Gunakan blank identifier (_) supaya variabel invalidLat dianggap "terpakai"
	_ = invalidLat 

	isValid := true 

	if isValid {
		t.Errorf("Unit Test FAILED: Fungsi validasi belum memfilter angka di luar range -90 s/d 90")
	}
}