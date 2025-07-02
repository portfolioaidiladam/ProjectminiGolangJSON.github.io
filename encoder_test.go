package belajar_golang_json

import (
	"encoding/json" // Import package encoding/json untuk proses encoding dan decoding JSON
	"os"            // Import package os untuk operasi file
	"testing"       // Import package testing untuk unit test
)

// TestEncoder menguji proses encoding struct Customer ke file JSON menggunakan stream encoder
func TestEncoder(t *testing.T) {
	// Membuat file baru CustomerOut.json untuk output JSON
	writer, _ := os.Create("CustomerOut.json")
	// Membuat encoder JSON yang menulis ke file writer
	encoder := json.NewEncoder(writer)

	// Membuat data customer
	customer := Customer{
		FirstName:  "Aidil",      // Nama depan
		MiddleName: "Adam",       // Nama tengah
		LastName:   "BaikHati",   // Nama belakang
	}

	// Melakukan encoding struct customer ke file JSON
	encoder.Encode(customer)
}

