//go:build functional
// +build functional

package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq" // Driver postgres
)

// Functional Test: Menguji koneksi ke database asli
func TestDatabaseConnection_Functional(t *testing.T) {
	// Mengambil info DB dari env (sesuai yang kita set di k8s/deployment.yaml)
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "db_location" // Default sesuai strategi tim
	}

	dbUser := "postgres"
	dbPass := "123"
	dbHost := "localhost" // Untuk test lokal

	// Simulasi string koneksi
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
		dbHost, dbUser, dbPass, dbName)

	// Mencoba membuka koneksi
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal inisialisasi driver: %v", err)
	}
	defer db.Close()

	// Melakukan ping ke database
	err = db.Ping()
	
	// Sengaja dibuat FAILED karena database mungkin belum up atau tabel belum dibuat
	if err != nil {
		t.Errorf("Functional Test FAILED: Tidak bisa terhubung ke database %s. Error: %v", dbName, err)
	}
}