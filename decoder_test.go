package belajar_golang_json

import (
	"encoding/json" // Import package encoding/json untuk proses encoding dan decoding JSON
	"fmt"           // Import package fmt untuk output ke console
	"os"            // Import package os untuk operasi file
	"testing"       // Import package testing untuk unit test
)

// TestStreamDecoder menguji proses decoding JSON dari file secara streaming
func TestStreamDecoder(t *testing.T) {

	// Membuka file Customer.json untuk dibaca
	reader, _ := os.Open("Customer.json")
	// Membuat decoder JSON dari file reader
	decoder := json.NewDecoder(reader)

	customer := &Customer{} // Membuat pointer ke struct Customer
	// Melakukan decoding JSON dari file ke struct Customer
	decoder.Decode(customer)

	// Menampilkan hasil decoding ke console
	fmt.Println(customer)

}

