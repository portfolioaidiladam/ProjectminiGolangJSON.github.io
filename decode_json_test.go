package belajar_golang_json // Mendefinisikan nama package

import (
	"encoding/json" // Import package encoding/json untuk proses decoding JSON
	"fmt"           // Import package fmt untuk output ke console
	"testing"       // Import package testing untuk unit test
)

// TestDecodeJSON adalah unit test untuk decode JSON menjadi struct Customer
func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Aidil","MiddleName":"Adam","LastName":"BaikHati","Age":36,"Married":true}` // JSON dalam bentuk string
	jsonBytes := []byte(jsonString) // Konversi string JSON ke []byte

	customer := &Customer{} // Membuat pointer ke struct Customer kosong

	err := json.Unmarshal(jsonBytes, customer) // Decode JSON ke struct Customer
	if err != nil { // Jika terjadi error saat decoding
		panic(err) // Hentikan eksekusi dan tampilkan error
	}

	fmt.Println(customer)            // Cetak struct customer hasil decode
	fmt.Println(customer.FirstName)  // Cetak field FirstName
	fmt.Println(customer.MiddleName) // Cetak field MiddleName
	fmt.Println(customer.LastName)   // Cetak field LastName
}

