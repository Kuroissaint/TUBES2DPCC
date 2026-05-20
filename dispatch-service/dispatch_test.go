package main

import (
	"testing"
)

// Unit Test: Menguji logika pencarian driver terdekat
func TestFindNearestDriver_Unit(t *testing.T) {
	// Skenario: Mencari driver di sekitar titik koordinat tertentu
	orderLat, orderLon := -6.200000, 106.816666
	
	// Kita panggil di log agar variabel "terpakai" dan tidak error build
	t.Logf("Mencari driver terdekat untuk lokasi: %f, %f", orderLat, orderLon)

	// Simulasi: Karena algoritmanya belum dibuat, kita set hasil kosong
	foundDriverID := "" 

	if foundDriverID == "" {
		// Dibuat FAILED sesuai permintaan tugas tahap 2
		t.Errorf("Unit Test FAILED: Algoritma pencarian driver (Matching) belum diimplementasi")
	}
}

// Unit Test: Menguji status penugasan driver
func TestAssignDriverStatus_Unit(t *testing.T) {
	statusAwal := "pending"
	
	// Kita gunakan blank identifier (_) agar statusAwal dianggap terpakai
	_ = statusAwal

	// Harusnya status berubah jadi "assigned" setelah dipasangkan
	statusAkhir := "pending" 

	if statusAkhir != "assigned" {
		t.Errorf("Unit Test FAILED: Logika penugasan gagal, status tetap: %s", statusAkhir)
	}
}